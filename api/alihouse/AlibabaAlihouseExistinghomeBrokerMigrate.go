package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeBrokerMigrate 融合店经纪人迁移
// alibaba.alihouse.existinghome.broker.migrate
//
// 融合店经纪人迁移
func AlibabaAlihouseExistinghomeBrokerMigrate(clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeBrokerMigrateAPIRequest, session string) (*alihouse.AlibabaAlihouseExistinghomeBrokerMigrateAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseExistinghomeBrokerMigrateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
