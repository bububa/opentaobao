package travel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/travel"
)

/* TaobaoAlitripTravelItemElementQuery
【API3.0】资源元素查询接口
taobao.alitrip.travel.item.element.query

提供资源元素查询接口，支持商家查询已经发布过的资源元素 */
func TaobaoAlitripTravelItemElementQuery(clt *core.SDKClient, req *travel.TaobaoAlitripTravelItemElementQueryAPIRequest, session string) (*travel.TaobaoAlitripTravelItemElementQueryAPIResponse, error) {
	var resp travel.TaobaoAlitripTravelItemElementQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
