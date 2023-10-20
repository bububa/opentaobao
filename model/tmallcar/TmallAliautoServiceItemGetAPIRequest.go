package tmallcar

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallaliautoserviceitemgetAPIRequest 查询服务商门店已上架服务商品列表 API请求
// tmall.aliauto.service.item.get
//
// 根据门店自定义门店编码查询门店【已上架】服务商品列表
type TmallaliautoserviceitemgetAPIRequest struct {
	model.Params
	// 商家自定义门店编码
	_outerShopId string
}

// NewTmallaliautoserviceitemgetRequest 初始化TmallaliautoserviceitemgetAPIRequest对象
func NewTmallaliautoserviceitemgetRequest() *TmallaliautoserviceitemgetAPIRequest {
	return &TmallaliautoserviceitemgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallaliautoserviceitemgetAPIRequest) GetApiMethodName() string {
	return "tmall.aliauto.service.item.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallaliautoserviceitemgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallaliautoserviceitemgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOuterShopId is OuterShopId Setter
// 商家自定义门店编码
func (r *TmallaliautoserviceitemgetAPIRequest) SetOuterShopId(_outerShopId string) error {
	r._outerShopId = _outerShopId
	r.Set("outer_shop_id", _outerShopId)
	return nil
}

// GetOuterShopId OuterShopId Getter
func (r TmallaliautoserviceitemgetAPIRequest) GetOuterShopId() string {
	return r._outerShopId
}
