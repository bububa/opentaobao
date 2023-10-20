package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingBuygiftItemRemoveAsyncAPIRequest 批量删除买赠商品 API请求
// alibaba.hm.marketing.buygift.item.remove.async
//
// 批量删除买赠商品
type AlibabaHmMarketingBuygiftItemRemoveAsyncAPIRequest struct {
	model.Params
	// sku信息
	_param0 []ItemBuyGiftSku
	// 系统自动生成
	_param1 *CommonActivityParam
	// alibaba.wdk.marketing.version.generate接口生成
	_version int64
}

// NewAlibabaHmMarketingBuygiftItemRemoveAsyncRequest 初始化AlibabaHmMarketingBuygiftItemRemoveAsyncAPIRequest对象
func NewAlibabaHmMarketingBuygiftItemRemoveAsyncRequest() *AlibabaHmMarketingBuygiftItemRemoveAsyncAPIRequest {
	return &AlibabaHmMarketingBuygiftItemRemoveAsyncAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaHmMarketingBuygiftItemRemoveAsyncAPIRequest) Reset() {
	r._param0 = r._param0[:0]
	r._param1 = nil
	r._version = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHmMarketingBuygiftItemRemoveAsyncAPIRequest) GetApiMethodName() string {
	return "alibaba.hm.marketing.buygift.item.remove.async"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHmMarketingBuygiftItemRemoveAsyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaHmMarketingBuygiftItemRemoveAsyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// sku信息
func (r *AlibabaHmMarketingBuygiftItemRemoveAsyncAPIRequest) SetParam0(_param0 []ItemBuyGiftSku) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaHmMarketingBuygiftItemRemoveAsyncAPIRequest) GetParam0() []ItemBuyGiftSku {
	return r._param0
}

// SetParam1 is Param1 Setter
// 系统自动生成
func (r *AlibabaHmMarketingBuygiftItemRemoveAsyncAPIRequest) SetParam1(_param1 *CommonActivityParam) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AlibabaHmMarketingBuygiftItemRemoveAsyncAPIRequest) GetParam1() *CommonActivityParam {
	return r._param1
}

// SetVersion is Version Setter
// alibaba.wdk.marketing.version.generate接口生成
func (r *AlibabaHmMarketingBuygiftItemRemoveAsyncAPIRequest) SetVersion(_version int64) error {
	r._version = _version
	r.Set("version", _version)
	return nil
}

// GetVersion Version Getter
func (r AlibabaHmMarketingBuygiftItemRemoveAsyncAPIRequest) GetVersion() int64 {
	return r._version
}

var poolAlibabaHmMarketingBuygiftItemRemoveAsyncAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaHmMarketingBuygiftItemRemoveAsyncRequest()
	},
}

// GetAlibabaHmMarketingBuygiftItemRemoveAsyncRequest 从 sync.Pool 获取 AlibabaHmMarketingBuygiftItemRemoveAsyncAPIRequest
func GetAlibabaHmMarketingBuygiftItemRemoveAsyncAPIRequest() *AlibabaHmMarketingBuygiftItemRemoveAsyncAPIRequest {
	return poolAlibabaHmMarketingBuygiftItemRemoveAsyncAPIRequest.Get().(*AlibabaHmMarketingBuygiftItemRemoveAsyncAPIRequest)
}

// ReleaseAlibabaHmMarketingBuygiftItemRemoveAsyncAPIRequest 将 AlibabaHmMarketingBuygiftItemRemoveAsyncAPIRequest 放入 sync.Pool
func ReleaseAlibabaHmMarketingBuygiftItemRemoveAsyncAPIRequest(v *AlibabaHmMarketingBuygiftItemRemoveAsyncAPIRequest) {
	v.Reset()
	poolAlibabaHmMarketingBuygiftItemRemoveAsyncAPIRequest.Put(v)
}
