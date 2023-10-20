package waybill

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/waybill"
)

// Cainiaocloudprinttemplatesmigrate 云打印模板迁移接口
// cainiao.cloudprint.templates.migrate
//
// 云打印模板迁移接口
func Cainiaocloudprinttemplatesmigrate(clt *core.SDKClient, req *waybill.CainiaocloudprinttemplatesmigrateAPIRequest, session string) (*waybill.CainiaocloudprinttemplatesmigrateAPIResponse, error) {
	var resp waybill.CainiaocloudprinttemplatesmigrateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
