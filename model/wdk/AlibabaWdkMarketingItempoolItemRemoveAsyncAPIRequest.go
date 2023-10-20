package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingItempoolItemRemoveAsyncAPIRequest 商品池删除商品 API请求
// alibaba.wdk.marketing.itempool.item.remove.async
//
// 新模型下删除商品
type AlibabaWdkMarketingItempoolItemRemoveAsyncAPIRequest struct {
	model.Params
	// sku信息
	_param0 []ItemPoolSku
	// 活动信息
	_param1 *CommonActivityParam
	// alibaba.wdk.marketing.version.generate接口生成
	_version int64
}

// NewAlibabaWdkMarketingItempoolItemRemoveAsyncRequest 初始化AlibabaWdkMarketingItempoolItemRemoveAsyncAPIRequest对象
func NewAlibabaWdkMarketingItempoolItemRemoveAsyncRequest() *AlibabaWdkMarketingItempoolItemRemoveAsyncAPIRequest {
	return &AlibabaWdkMarketingItempoolItemRemoveAsyncAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkMarketingItempoolItemRemoveAsyncAPIRequest) Reset() {
	r._param0 = r._param0[:0]
	r._param1 = nil
	r._version = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingItempoolItemRemoveAsyncAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.itempool.item.remove.async"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingItempoolItemRemoveAsyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkMarketingItempoolItemRemoveAsyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// sku信息
func (r *AlibabaWdkMarketingItempoolItemRemoveAsyncAPIRequest) SetParam0(_param0 []ItemPoolSku) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaWdkMarketingItempoolItemRemoveAsyncAPIRequest) GetParam0() []ItemPoolSku {
	return r._param0
}

// SetParam1 is Param1 Setter
// 活动信息
func (r *AlibabaWdkMarketingItempoolItemRemoveAsyncAPIRequest) SetParam1(_param1 *CommonActivityParam) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AlibabaWdkMarketingItempoolItemRemoveAsyncAPIRequest) GetParam1() *CommonActivityParam {
	return r._param1
}

// SetVersion is Version Setter
// alibaba.wdk.marketing.version.generate接口生成
func (r *AlibabaWdkMarketingItempoolItemRemoveAsyncAPIRequest) SetVersion(_version int64) error {
	r._version = _version
	r.Set("version", _version)
	return nil
}

// GetVersion Version Getter
func (r AlibabaWdkMarketingItempoolItemRemoveAsyncAPIRequest) GetVersion() int64 {
	return r._version
}

var poolAlibabaWdkMarketingItempoolItemRemoveAsyncAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkMarketingItempoolItemRemoveAsyncRequest()
	},
}

// GetAlibabaWdkMarketingItempoolItemRemoveAsyncRequest 从 sync.Pool 获取 AlibabaWdkMarketingItempoolItemRemoveAsyncAPIRequest
func GetAlibabaWdkMarketingItempoolItemRemoveAsyncAPIRequest() *AlibabaWdkMarketingItempoolItemRemoveAsyncAPIRequest {
	return poolAlibabaWdkMarketingItempoolItemRemoveAsyncAPIRequest.Get().(*AlibabaWdkMarketingItempoolItemRemoveAsyncAPIRequest)
}

// ReleaseAlibabaWdkMarketingItempoolItemRemoveAsyncAPIRequest 将 AlibabaWdkMarketingItempoolItemRemoveAsyncAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkMarketingItempoolItemRemoveAsyncAPIRequest(v *AlibabaWdkMarketingItempoolItemRemoveAsyncAPIRequest) {
	v.Reset()
	poolAlibabaWdkMarketingItempoolItemRemoveAsyncAPIRequest.Put(v)
}
