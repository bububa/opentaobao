package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaRetailMarketingItemdiscountActivitySaveAPIRequest 【同城零售】单品活动保存 API请求
// alibaba.retail.marketing.itemdiscount.activity.save
//
// 【同城零售】单品活动保存
type AlibabaRetailMarketingItemdiscountActivitySaveAPIRequest struct {
	model.Params
	// 保存活动参数
	_param *ItemDiscountActivityOperateRequest
}

// NewAlibabaRetailMarketingItemdiscountActivitySaveRequest 初始化AlibabaRetailMarketingItemdiscountActivitySaveAPIRequest对象
func NewAlibabaRetailMarketingItemdiscountActivitySaveRequest() *AlibabaRetailMarketingItemdiscountActivitySaveAPIRequest {
	return &AlibabaRetailMarketingItemdiscountActivitySaveAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaRetailMarketingItemdiscountActivitySaveAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaRetailMarketingItemdiscountActivitySaveAPIRequest) GetApiMethodName() string {
	return "alibaba.retail.marketing.itemdiscount.activity.save"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaRetailMarketingItemdiscountActivitySaveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaRetailMarketingItemdiscountActivitySaveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 保存活动参数
func (r *AlibabaRetailMarketingItemdiscountActivitySaveAPIRequest) SetParam(_param *ItemDiscountActivityOperateRequest) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaRetailMarketingItemdiscountActivitySaveAPIRequest) GetParam() *ItemDiscountActivityOperateRequest {
	return r._param
}

var poolAlibabaRetailMarketingItemdiscountActivitySaveAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaRetailMarketingItemdiscountActivitySaveRequest()
	},
}

// GetAlibabaRetailMarketingItemdiscountActivitySaveRequest 从 sync.Pool 获取 AlibabaRetailMarketingItemdiscountActivitySaveAPIRequest
func GetAlibabaRetailMarketingItemdiscountActivitySaveAPIRequest() *AlibabaRetailMarketingItemdiscountActivitySaveAPIRequest {
	return poolAlibabaRetailMarketingItemdiscountActivitySaveAPIRequest.Get().(*AlibabaRetailMarketingItemdiscountActivitySaveAPIRequest)
}

// ReleaseAlibabaRetailMarketingItemdiscountActivitySaveAPIRequest 将 AlibabaRetailMarketingItemdiscountActivitySaveAPIRequest 放入 sync.Pool
func ReleaseAlibabaRetailMarketingItemdiscountActivitySaveAPIRequest(v *AlibabaRetailMarketingItemdiscountActivitySaveAPIRequest) {
	v.Reset()
	poolAlibabaRetailMarketingItemdiscountActivitySaveAPIRequest.Put(v)
}
