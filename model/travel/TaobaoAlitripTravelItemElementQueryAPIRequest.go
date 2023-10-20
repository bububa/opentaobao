package travel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitriptravelitemelementqueryAPIRequest 【API3.0】资源元素查询接口 API请求
// taobao.alitrip.travel.item.element.query
//
// 提供资源元素查询接口，支持商家查询已经发布过的资源元素
type TaobaoalitriptravelitemelementqueryAPIRequest struct {
	model.Params
	// 需要查询的资源元素列表，最大列表长度为50
	_outerIds []string
}

// NewTaobaoalitriptravelitemelementqueryRequest 初始化TaobaoalitriptravelitemelementqueryAPIRequest对象
func NewTaobaoalitriptravelitemelementqueryRequest() *TaobaoalitriptravelitemelementqueryAPIRequest {
	return &TaobaoalitriptravelitemelementqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitriptravelitemelementqueryAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.item.element.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitriptravelitemelementqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitriptravelitemelementqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOuterIds is OuterIds Setter
// 需要查询的资源元素列表，最大列表长度为50
func (r *TaobaoalitriptravelitemelementqueryAPIRequest) SetOuterIds(_outerIds []string) error {
	r._outerIds = _outerIds
	r.Set("outer_ids", _outerIds)
	return nil
}

// GetOuterIds OuterIds Getter
func (r TaobaoalitriptravelitemelementqueryAPIRequest) GetOuterIds() []string {
	return r._outerIds
}
