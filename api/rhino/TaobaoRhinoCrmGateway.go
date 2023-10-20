package rhino

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/rhino"
)

// Taobaorhinocrmgateway crm实体变更回调接口
// taobao.rhino.crm.gateway
//
// crm实体变更回调接口
func Taobaorhinocrmgateway(clt *core.SDKClient, req *rhino.TaobaorhinocrmgatewayAPIRequest, session string) (*rhino.TaobaorhinocrmgatewayAPIResponse, error) {
	var resp rhino.TaobaorhinocrmgatewayAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
