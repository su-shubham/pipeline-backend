package migration

import (
	"context"

	"github.com/instill-ai/pipeline-backend/pkg/db/migration/convert/convert000013"
	"github.com/instill-ai/pipeline-backend/pkg/db/migration/convert/convert000015"
	"github.com/instill-ai/pipeline-backend/pkg/db/migration/convert/convert000016"
	"github.com/instill-ai/pipeline-backend/pkg/db/migration/convert/convert000019"
	"github.com/instill-ai/pipeline-backend/pkg/db/migration/convert/convert000020"
	"github.com/instill-ai/pipeline-backend/pkg/db/migration/convert/convert000021"
	"github.com/instill-ai/pipeline-backend/pkg/db/migration/convert/convert000022"
	"github.com/instill-ai/pipeline-backend/pkg/db/migration/convert/convert000024"
	"github.com/instill-ai/pipeline-backend/pkg/db/migration/convert/convert000029"
	"github.com/instill-ai/pipeline-backend/pkg/external"
	"github.com/instill-ai/pipeline-backend/pkg/logger"

	database "github.com/instill-ai/pipeline-backend/pkg/db"
)

type migration interface {
	Migrate() error
}

// Migrate executes custom code as part of a database migration. This code is
// intended to be run only once and typically goes along a change
// in the database schemas. Some use cases might be backfilling a table or
// updating some existing records according to the schema changes.
//
// Note that the changes in the database schemas shouldn't be run here, only
// code accompanying them.
func Migrate(version uint) error {
	var m migration
	ctx := context.Background()
	l, _ := logger.GetZapLogger(ctx)

	db := database.GetConnection().WithContext(ctx)
	defer database.Close(db)
	mgmtPrivateServiceClient, mgmtPrivateServiceClientConn := external.InitMgmtPrivateServiceClient(ctx)
	if mgmtPrivateServiceClientConn != nil {
		defer mgmtPrivateServiceClientConn.Close()
	}

	switch version {
	case 13:
		m = new(convert000013.Migration)
	case 15:
		m = new(convert000015.Migration)
	case 16:
		m = new(convert000016.Migration)
	case 19:
		m = &convert000019.JQInputToKebabCaseConverter{
			DB:     db,
			Logger: l,
		}
	case 20:
		m = &convert000020.NamespaceIDMigrator{
			DB:         db,
			Logger:     l,
			MgmtClient: mgmtPrivateServiceClient,
		}
	case 21:
		m = &convert000021.ConvertToTextTaskConverter{
			DB:     db,
			Logger: l,
		}
	case 22:
		m = &convert000022.ConvertWebsiteToWebConverter{
			DB:     db,
			Logger: l,
		}
	case 24:
		m = &convert000024.ConvertToTextTaskConverter{
			DB:     db,
			Logger: l,
		}
	case 29:
		m = &convert000029.ConvertToArtifactType{
			DB:     db,
			Logger: l,
		}
	default:
		return nil
	}

	return m.Migrate()
}
