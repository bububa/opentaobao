package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingFullrangeAddexchangeitemAPIRequest 全场增加换购品 API请求
// alibaba.wdk.marketing.fullrange.addexchangeitem
//
// 全场增加换购品
type AlibabaWdkMarketingFullrangeAddexchangeitemAPIRequest struct {
	model.Params
	// 系统自动生成
	_param0 *ItemStairSku
	// 系统自动生成
	_param1 *CommonActivityParam
}

// NewAlibabaWdkMarketingFullrangeAddexchangeitemRequest 初始化AlibabaWdkMarketingFullrangeAddexchangeitemAPIRequest对象
func NewAlibabaWdkMarketingFullrangeAddexchangeitemRequest() *AlibabaWdkMarketingFullrangeAddexchangeitemAPIRequest {
	return &AlibabaWdkMarketingFullrangeAddexchangeitemAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkMarketingFullrangeAddexchangeitemAPIRequest) Reset() {
	r._param0 = nil
	r._param1 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingFullrangeAddexchangeitemAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.fullrange.addexchangeitem"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingFullrangeAddexchangeitemAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkMarketingFullrangeAddexchangeitemAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 系统自动生成
func (r *AlibabaWdkMarketingFullrangeAddexchangeitemAPIRequest) SetParam0(_param0 *ItemStairSku) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaWdkMarketingFullrangeAddexchangeitemAPIRequest) GetParam0() *ItemStairSku {
	return r._param0
}

// SetParam1 is Param1 Setter
// 系统自动生成
func (r *AlibabaWdkMarketingFullrangeAddexchangeitemAPIRequest) SetParam1(_param1 *CommonActivityParam) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AlibabaWdkMarketingFullrangeAddexchangeitemAPIRequest) GetParam1() *CommonActivityParam {
	return r._param1
}

var poolAlibabaWdkMarketingFullrangeAddexchangeitemAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkMarketingFullrangeAddexchangeitemRequest()
	},
}

// GetAlibabaWdkMarketingFullrangeAddexchangeitemRequest 从 sync.Pool 获取 AlibabaWdkMarketingFullrangeAddexchangeitemAPIRequest
func GetAlibabaWdkMarketingFullrangeAddexchangeitemAPIRequest() *AlibabaWdkMarketingFullrangeAddexchangeitemAPIRequest {
	return poolAlibabaWdkMarketingFullrangeAddexchangeitemAPIRequest.Get().(*AlibabaWdkMarketingFullrangeAddexchangeitemAPIRequest)
}

// ReleaseAlibabaWdkMarketingFullrangeAddexchangeitemAPIRequest 将 AlibabaWdkMarketingFullrangeAddexchangeitemAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkMarketingFullrangeAddexchangeitemAPIRequest(v *AlibabaWdkMarketingFullrangeAddexchangeitemAPIRequest) {
	v.Reset()
	poolAlibabaWdkMarketingFullrangeAddexchangeitemAPIRequest.Put(v)
}
