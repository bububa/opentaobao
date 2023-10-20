package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingItempoolItemAddAsyncAPIRequest 商品池新增商品 API请求
// alibaba.hm.marketing.itempool.item.add.async
//
// 新分组模型下新增商品
type AlibabaHmMarketingItempoolItemAddAsyncAPIRequest struct {
	model.Params
	// 阶梯商品信息
	_param0 []ItemPoolSku
	// 系统自动生成
	_param1 *CommonActivityParam
	// alibaba.wdk.marketing.version.generate接口生成
	_version int64
}

// NewAlibabaHmMarketingItempoolItemAddAsyncRequest 初始化AlibabaHmMarketingItempoolItemAddAsyncAPIRequest对象
func NewAlibabaHmMarketingItempoolItemAddAsyncRequest() *AlibabaHmMarketingItempoolItemAddAsyncAPIRequest {
	return &AlibabaHmMarketingItempoolItemAddAsyncAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaHmMarketingItempoolItemAddAsyncAPIRequest) Reset() {
	r._param0 = r._param0[:0]
	r._param1 = nil
	r._version = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHmMarketingItempoolItemAddAsyncAPIRequest) GetApiMethodName() string {
	return "alibaba.hm.marketing.itempool.item.add.async"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHmMarketingItempoolItemAddAsyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaHmMarketingItempoolItemAddAsyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 阶梯商品信息
func (r *AlibabaHmMarketingItempoolItemAddAsyncAPIRequest) SetParam0(_param0 []ItemPoolSku) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaHmMarketingItempoolItemAddAsyncAPIRequest) GetParam0() []ItemPoolSku {
	return r._param0
}

// SetParam1 is Param1 Setter
// 系统自动生成
func (r *AlibabaHmMarketingItempoolItemAddAsyncAPIRequest) SetParam1(_param1 *CommonActivityParam) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AlibabaHmMarketingItempoolItemAddAsyncAPIRequest) GetParam1() *CommonActivityParam {
	return r._param1
}

// SetVersion is Version Setter
// alibaba.wdk.marketing.version.generate接口生成
func (r *AlibabaHmMarketingItempoolItemAddAsyncAPIRequest) SetVersion(_version int64) error {
	r._version = _version
	r.Set("version", _version)
	return nil
}

// GetVersion Version Getter
func (r AlibabaHmMarketingItempoolItemAddAsyncAPIRequest) GetVersion() int64 {
	return r._version
}

var poolAlibabaHmMarketingItempoolItemAddAsyncAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaHmMarketingItempoolItemAddAsyncRequest()
	},
}

// GetAlibabaHmMarketingItempoolItemAddAsyncRequest 从 sync.Pool 获取 AlibabaHmMarketingItempoolItemAddAsyncAPIRequest
func GetAlibabaHmMarketingItempoolItemAddAsyncAPIRequest() *AlibabaHmMarketingItempoolItemAddAsyncAPIRequest {
	return poolAlibabaHmMarketingItempoolItemAddAsyncAPIRequest.Get().(*AlibabaHmMarketingItempoolItemAddAsyncAPIRequest)
}

// ReleaseAlibabaHmMarketingItempoolItemAddAsyncAPIRequest 将 AlibabaHmMarketingItempoolItemAddAsyncAPIRequest 放入 sync.Pool
func ReleaseAlibabaHmMarketingItempoolItemAddAsyncAPIRequest(v *AlibabaHmMarketingItempoolItemAddAsyncAPIRequest) {
	v.Reset()
	poolAlibabaHmMarketingItempoolItemAddAsyncAPIRequest.Put(v)
}
