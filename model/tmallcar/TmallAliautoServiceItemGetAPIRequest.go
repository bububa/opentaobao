package tmallcar

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallAliautoServiceItemGetAPIRequest) GetApiMethodName() string {
	return "tmall.aliauto.service.item.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallAliautoServiceItemGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
