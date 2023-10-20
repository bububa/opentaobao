package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeBrokerMigrate 融合店经纪人迁移
// alibaba.alihouse.existinghome.broker.migrate
//
// 融合店经纪人迁移
func AlibabaAlihouseExistinghomeBrokerMigrate(clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeBrokerMigrateAPIRequest, resp *alihouse.AlibabaAlihouseExistinghomeBrokerMigrateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
