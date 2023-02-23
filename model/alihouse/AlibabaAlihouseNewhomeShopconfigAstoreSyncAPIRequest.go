package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeShopconfigAstoreSyncAPIRequest 天猫好房店铺装修-Astore上翻 API请求
// alibaba.alihouse.newhome.shopconfig.astore.sync
//
// 天猫好房店铺装修-Astore上翻
type AlibabaAlihouseNewhomeShopconfigAstoreSyncAPIRequest struct {
	model.Params
	// 装修配置信息
	_astoreShopConfig *AstoreShopConfigDto
}

// NewAlibabaAlihouseNewhomeShopconfigAstoreSyncRequest 初始化AlibabaAlihouseNewhomeShopconfigAstoreSyncAPIRequest对象
func NewAlibabaAlihouseNewhomeShopconfigAstoreSyncRequest() *AlibabaAlihouseNewhomeShopconfigAstoreSyncAPIRequest {
	return &AlibabaAlihouseNewhomeShopconfigAstoreSyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeShopconfigAstoreSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.shopconfig.astore.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeShopconfigAstoreSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseNewhomeShopconfigAstoreSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAstoreShopConfig is AstoreShopConfig Setter
// 装修配置信息
func (r *AlibabaAlihouseNewhomeShopconfigAstoreSyncAPIRequest) SetAstoreShopConfig(_astoreShopConfig *AstoreShopConfigDto) error {
	r._astoreShopConfig = _astoreShopConfig
	r.Set("astore_shop_config", _astoreShopConfig)
	return nil
}

// GetAstoreShopConfig AstoreShopConfig Getter
func (r AlibabaAlihouseNewhomeShopconfigAstoreSyncAPIRequest) GetAstoreShopConfig() *AstoreShopConfigDto {
	return r._astoreShopConfig
}
