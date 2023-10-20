package tmallcar

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallAliautoServiceItemGetAPIRequest 查询服务商门店已上架服务商品列表 API请求
// tmall.aliauto.service.item.get
//
// 根据门店自定义门店编码查询门店【已上架】服务商品列表
type TmallAliautoServiceItemGetAPIRequest struct {
	model.Params
	// 商家自定义门店编码
	_outerShopId string
}

// NewTmallAliautoServiceItemGetRequest 初始化TmallAliautoServiceItemGetAPIRequest对象
func NewTmallAliautoServiceItemGetRequest() *TmallAliautoServiceItemGetAPIRequest {
	return &TmallAliautoServiceItemGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallAliautoServiceItemGetAPIRequest) Reset() {
	r._outerShopId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallAliautoServiceItemGetAPIRequest) GetApiMethodName() string {
	return "tmall.aliauto.service.item.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallAliautoServiceItemGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallAliautoServiceItemGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOuterShopId is OuterShopId Setter
// 商家自定义门店编码
func (r *TmallAliautoServiceItemGetAPIRequest) SetOuterShopId(_outerShopId string) error {
	r._outerShopId = _outerShopId
	r.Set("outer_shop_id", _outerShopId)
	return nil
}

// GetOuterShopId OuterShopId Getter
func (r TmallAliautoServiceItemGetAPIRequest) GetOuterShopId() string {
	return r._outerShopId
}

var poolTmallAliautoServiceItemGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallAliautoServiceItemGetRequest()
	},
}

// GetTmallAliautoServiceItemGetRequest 从 sync.Pool 获取 TmallAliautoServiceItemGetAPIRequest
func GetTmallAliautoServiceItemGetAPIRequest() *TmallAliautoServiceItemGetAPIRequest {
	return poolTmallAliautoServiceItemGetAPIRequest.Get().(*TmallAliautoServiceItemGetAPIRequest)
}

// ReleaseTmallAliautoServiceItemGetAPIRequest 将 TmallAliautoServiceItemGetAPIRequest 放入 sync.Pool
func ReleaseTmallAliautoServiceItemGetAPIRequest(v *TmallAliautoServiceItemGetAPIRequest) {
	v.Reset()
	poolTmallAliautoServiceItemGetAPIRequest.Put(v)
}
