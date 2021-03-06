package waybill

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/waybill"
)

// CainiaoCloudprintTemplatesMigrate 云打印模板迁移接口
// cainiao.cloudprint.templates.migrate
//
// 云打印模板迁移接口
func CainiaoCloudprintTemplatesMigrate(clt *core.SDKClient, req *waybill.CainiaoCloudprintTemplatesMigrateAPIRequest, session string) (*waybill.CainiaoCloudprintTemplatesMigrateAPIResponse, error) {
	var resp waybill.CainiaoCloudprintTemplatesMigrateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
