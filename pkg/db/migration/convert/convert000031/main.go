package convert000031

import (
	"fmt"

	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
	"gorm.io/gorm"

	"github.com/instill-ai/pipeline-backend/pkg/datamodel"
)

const batchSize = 100

type UpdateRecipeFields struct {
	DB     *gorm.DB
	Logger *zap.Logger
}

func (u *UpdateRecipeFields) Migrate() error {
	if err := u.migratePipeline(); err != nil {
		return err
	}

	return u.migratePipelineRelease()
}

func (u *UpdateRecipeFields) migratePipeline() error {
	pipelines := make([]*datamodel.Pipeline, 0, batchSize)
	return u.DB.Select("uid", "recipe_yaml", "recipe").FindInBatches(&pipelines, batchSize, func(tx *gorm.DB, _ int) error {
		for _, p := range pipelines {
			isRecipeUpdated := false
			l := u.Logger.With(zap.String("pipelineUID", p.UID.String()))

			for id, comp := range p.Recipe.Component {
				isComponentUpdated, err := u.updateFields(comp)
				if err != nil {
					l.With(zap.String("componentID", id), zap.Error(err)).
						Error("Failed to update pipeline.")
					return fmt.Errorf("updating pipeline component: %w", err)
				}

				isRecipeUpdated = isComponentUpdated || isRecipeUpdated
			}

			if isRecipeUpdated {
				recipeYAML, err := yaml.Marshal(p.Recipe)
				if err != nil {
					return fmt.Errorf("marshalling recipe: %w", err)
				}
				result := tx.Model(p).Where("uid = ?", p.UID).Update("recipe_yaml", string(recipeYAML))
				if result.Error != nil {
					l.Error("Failed to update pipeline.")
					return fmt.Errorf("updating pipeline recipe: %w", result.Error)
				}
			}
		}

		return nil
	}).Error
}

func (u *UpdateRecipeFields) migratePipelineRelease() error {
	pipelineReleases := make([]*datamodel.PipelineRelease, 0, batchSize)
	return u.DB.Select("uid", "recipe_yaml", "recipe").FindInBatches(&pipelineReleases, batchSize, func(tx *gorm.DB, _ int) error {
		for _, pr := range pipelineReleases {
			isRecipeUpdated := false
			l := u.Logger.With(zap.String("pipelineReleaseUID", pr.UID.String()))

			for id, comp := range pr.Recipe.Component {
				isComponentUpdated, err := u.updateFields(comp)
				if err != nil {
					l.With(zap.String("componentID", id), zap.Error(err)).
						Error("Failed to update pipeline release.")
					return fmt.Errorf("updating pipeline release component: %w", err)
				}

				isRecipeUpdated = isComponentUpdated || isRecipeUpdated
			}

			if isRecipeUpdated {
				recipeYAML, err := yaml.Marshal(pr.Recipe)
				if err != nil {
					return fmt.Errorf("marshalling recipe: %w", err)
				}

				result := tx.Model(pr).Where("uid = ?", pr.UID).Update("recipe_yaml", string(recipeYAML))
				if result.Error != nil {
					l.Error("Failed to update pipeline release.")
					return fmt.Errorf("updating pipeline release recipe: %w", result.Error)
				}
			}
		}

		return nil
	}).Error
}

func (u *UpdateRecipeFields) updateFields(comp *datamodel.Component) (bool, error) {
	if comp.Type == "iterator" {
		isComponentUpdated := false
		for _, subComp := range comp.Component {
			isSubComponentUpdated, err := u.updateFields(subComp)
			if err != nil {
				return false, fmt.Errorf("updating iterator component: %w", err)
			}

			isComponentUpdated = isSubComponentUpdated || isComponentUpdated
		}

		return isComponentUpdated, nil
	}

	isUpdated := false

	// Update the 'memo' field to 'description' if it exists
	if _, exists := comp.Metadata["memo"]; exists {
		comp.Metadata["description"] = comp.Metadata["memo"]
		delete(comp.Metadata, "memo")
		isUpdated = true
	}

	// Add more field updates here if needed

	return isUpdated, nil
}