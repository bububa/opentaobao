package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeShopconfigAstorePreviewAPIRequest 天猫好房店铺装修-地址预览 API请求
// alibaba.alihouse.newhome.shopconfig.astore.preview
//
// 天猫好房店铺装修-Astore上翻
type AlibabaAlihouseNewhomeShopconfigAstorePreviewAPIRequest struct {
	model.Params
	// 装修配置信息
	_astoreShopConfig *AstoreShopConfigDto
}

// NewAlibabaAlihouseNewhomeShopconfigAstorePreviewRequest 初始化AlibabaAlihouseNewhomeShopconfigAstorePreviewAPIRequest对象
func NewAlibabaAlihouseNewhomeShopconfigAstorePreviewRequest() *AlibabaAlihouseNewhomeShopconfigAstorePreviewAPIRequest {
	return &AlibabaAlihouseNewhomeShopconfigAstorePreviewAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeShopconfigAstorePreviewAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.shopconfig.astore.preview"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeShopconfigAstorePreviewAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseNewhomeShopconfigAstorePreviewAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAstoreShopConfig is AstoreShopConfig Setter
// 装修配置信息
func (r *AlibabaAlihouseNewhomeShopconfigAstorePreviewAPIRequest) SetAstoreShopConfig(_astoreShopConfig *AstoreShopConfigDto) error {
	r._astoreShopConfig = _astoreShopConfig
	r.Set("astore_shop_config", _astoreShopConfig)
	return nil
}

// GetAstoreShopConfig AstoreShopConfig Getter
func (r AlibabaAlihouseNewhomeShopconfigAstorePreviewAPIRequest) GetAstoreShopConfig() *AstoreShopConfigDto {
	return r._astoreShopConfig
}
