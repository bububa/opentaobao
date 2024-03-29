package alihouse

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseNewhomeShopconfigAstorePreviewAPIRequest) Reset() {
	r._astoreShopConfig = nil
	r.Params.ToZero()
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

var poolAlibabaAlihouseNewhomeShopconfigAstorePreviewAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseNewhomeShopconfigAstorePreviewRequest()
	},
}

// GetAlibabaAlihouseNewhomeShopconfigAstorePreviewRequest 从 sync.Pool 获取 AlibabaAlihouseNewhomeShopconfigAstorePreviewAPIRequest
func GetAlibabaAlihouseNewhomeShopconfigAstorePreviewAPIRequest() *AlibabaAlihouseNewhomeShopconfigAstorePreviewAPIRequest {
	return poolAlibabaAlihouseNewhomeShopconfigAstorePreviewAPIRequest.Get().(*AlibabaAlihouseNewhomeShopconfigAstorePreviewAPIRequest)
}

// ReleaseAlibabaAlihouseNewhomeShopconfigAstorePreviewAPIRequest 将 AlibabaAlihouseNewhomeShopconfigAstorePreviewAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseNewhomeShopconfigAstorePreviewAPIRequest(v *AlibabaAlihouseNewhomeShopconfigAstorePreviewAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseNewhomeShopconfigAstorePreviewAPIRequest.Put(v)
}
