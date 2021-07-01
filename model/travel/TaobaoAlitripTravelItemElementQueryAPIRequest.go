package travel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripTravelItemElementQueryAPIRequest
【API3.0】资源元素查询接口 API请求
taobao.alitrip.travel.item.element.query

提供资源元素查询接口，支持商家查询已经发布过的资源元素 */
type TaobaoAlitripTravelItemElementQueryAPIRequest struct {
	model.Params
	// 需要查询的资源元素列表，最大列表长度为50
	_outerIds []string
}

// New
