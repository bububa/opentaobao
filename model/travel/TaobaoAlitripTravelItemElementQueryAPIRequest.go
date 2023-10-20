package travel

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelItemElementQueryAPIRequest 【API3.0】资源元素查询接口 API请求
// taobao.alitrip.travel.item.element.query
//
// 提供资源元素查询接口，支持商家查询已经发布过的资源元素
type TaobaoAlitripTravelItemElementQueryAPIRequest struct {
	model.Params
	// 需要查询的资源元素列表，最大列表长度为50
	_outerIds []string
}

// NewTaobaoAlitripTravelItemElementQueryRequest 初始化TaobaoAlitripTravelItemElementQueryAPIRequest对象
func NewTaobaoAlitripTravelItemElementQueryRequest() *TaobaoAlitripTravelItemElementQueryAPIRequest {
	return &TaobaoAlitripTravelItemElementQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAlitripTravelItemElementQueryAPIRequest) Reset() {
	r._outerIds = r._outerIds[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelItemElementQueryAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.item.element.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelItemElementQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlitripTravelItemElementQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOuterIds is OuterIds Setter
// 需要查询的资源元素列表，最大列表长度为50
func (r *TaobaoAlitripTravelItemElementQueryAPIRequest) SetOuterIds(_outerIds []string) error {
	r._outerIds = _outerIds
	r.Set("outer_ids", _outerIds)
	return nil
}

// GetOuterIds OuterIds Getter
func (r TaobaoAlitripTravelItemElementQueryAPIRequest) GetOuterIds() []string {
	return r._outerIds
}

var poolTaobaoAlitripTravelItemElementQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAlitripTravelItemElementQueryRequest()
	},
}

// GetTaobaoAlitripTravelItemElementQueryRequest 从 sync.Pool 获取 TaobaoAlitripTravelItemElementQueryAPIRequest
func GetTaobaoAlitripTravelItemElementQueryAPIRequest() *TaobaoAlitripTravelItemElementQueryAPIRequest {
	return poolTaobaoAlitripTravelItemElementQueryAPIRequest.Get().(*TaobaoAlitripTravelItemElementQueryAPIRequest)
}

// ReleaseTaobaoAlitripTravelItemElementQueryAPIRequest 将 TaobaoAlitripTravelItemElementQueryAPIRequest 放入 sync.Pool
func ReleaseTaobaoAlitripTravelItemElementQueryAPIRequest(v *TaobaoAlitripTravelItemElementQueryAPIRequest) {
	v.Reset()
	poolTaobaoAlitripTravelItemElementQueryAPIRequest.Put(v)
}
