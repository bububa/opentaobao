package travel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/travel"
)

// TaobaoAlitripTravelItemElementQuery 【API3.0】资源元素查询接口
// taobao.alitrip.travel.item.element.query
//
// 提供资源元素查询接口，支持商家查询已经发布过的资源元素
func TaobaoAlitripTravelItemElementQuery(clt *core.SDKClient, req *travel.TaobaoAlitripTravelItemElementQueryAPIRequest, resp *travel.TaobaoAlitripTravelItemElementQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
