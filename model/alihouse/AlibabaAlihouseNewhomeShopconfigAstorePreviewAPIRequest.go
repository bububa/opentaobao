package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomeshopconfigastorepreviewAPIRequest 天猫好房店铺装修-地址预览 API请求
// alibaba.alihouse.newhome.shopconfig.astore.preview
//
// 天猫好房店铺装修-Astore上翻
type AlibabaalihousenewhomeshopconfigastorepreviewAPIRequest struct {
	model.Params
	// 装修配置信息
	_astoreShopConfig *AstoreShopConfigDto
}

// NewAlibabaalihousenewhomeshopconfigastorepreviewRequest 初始化AlibabaalihousenewhomeshopconfigastorepreviewAPIRequest对象
func NewAlibabaalihousenewhomeshopconfigastorepreviewRequest() *AlibabaalihousenewhomeshopconfigastorepreviewAPIRequest {
	return &AlibabaalihousenewhomeshopconfigastorepreviewAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihousenewhomeshopconfigastorepreviewAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.shopconfig.astore.preview"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihousenewhomeshopconfigastorepreviewAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihousenewhomeshopconfigastorepreviewAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAstoreShopConfig is AstoreShopConfig Setter
// 装修配置信息
func (r *AlibabaalihousenewhomeshopconfigastorepreviewAPIRequest) SetAstoreShopConfig(_astoreShopConfig *AstoreShopConfigDto) error {
	r._astoreShopConfig = _astoreShopConfig
	r.Set("astore_shop_config", _astoreShopConfig)
	return nil
}

// GetAstoreShopConfig AstoreShopConfig Getter
func (r AlibabaalihousenewhomeshopconfigastorepreviewAPIRequest) GetAstoreShopConfig() *AstoreShopConfigDto {
	return r._astoreShopConfig
}
