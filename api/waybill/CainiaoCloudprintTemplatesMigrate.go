package waybill

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/waybill"
)

// CainiaoCloudprintTemplatesMigrate 云打印模板迁移接口
// cainiao.cloudprint.templates.migrate
//
// 云打印模板迁移接口
func CainiaoCloudprintTemplatesMigrate(ctx context.Context, clt *core.SDKClient, req *waybill.CainiaoCloudprintTemplatesMigrateAPIRequest, resp *waybill.CainiaoCloudprintTemplatesMigrateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
