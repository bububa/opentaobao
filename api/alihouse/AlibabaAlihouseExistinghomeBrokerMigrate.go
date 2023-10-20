package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihouseexistinghomebrokermigrate 融合店经纪人迁移
// alibaba.alihouse.existinghome.broker.migrate
//
// 融合店经纪人迁移
func Alibabaalihouseexistinghomebrokermigrate(clt *core.SDKClient, req *alihouse.AlibabaalihouseexistinghomebrokermigrateAPIRequest, session string) (*alihouse.AlibabaalihouseexistinghomebrokermigrateAPIResponse, error) {
	var resp alihouse.AlibabaalihouseexistinghomebrokermigrateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
