package alihouse

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseNewhomeShopconfigAstoreSyncAPIRequest) Reset() {
	r._astoreShopConfig = nil
	r.Params.ToZero()
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

var poolAlibabaAlihouseNewhomeShopconfigAstoreSyncAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseNewhomeShopconfigAstoreSyncRequest()
	},
}

// GetAlibabaAlihouseNewhomeShopconfigAstoreSyncRequest 从 sync.Pool 获取 AlibabaAlihouseNewhomeShopconfigAstoreSyncAPIRequest
func GetAlibabaAlihouseNewhomeShopconfigAstoreSyncAPIRequest() *AlibabaAlihouseNewhomeShopconfigAstoreSyncAPIRequest {
	return poolAlibabaAlihouseNewhomeShopconfigAstoreSyncAPIRequest.Get().(*AlibabaAlihouseNewhomeShopconfigAstoreSyncAPIRequest)
}

// ReleaseAlibabaAlihouseNewhomeShopconfigAstoreSyncAPIRequest 将 AlibabaAlihouseNewhomeShopconfigAstoreSyncAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseNewhomeShopconfigAstoreSyncAPIRequest(v *AlibabaAlihouseNewhomeShopconfigAstoreSyncAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseNewhomeShopconfigAstoreSyncAPIRequest.Put(v)
}
