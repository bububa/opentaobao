package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaRetailMarketingItemdiscountActivityDeleteAPIRequest 删除单品特价活动【同城零售】 API请求
// alibaba.retail.marketing.itemdiscount.activity.delete
//
// 同城零售单品特价活动删除
type AlibabaRetailMarketingItemdiscountActivityDeleteAPIRequest struct {
	model.Params
	// 删除活动参数
	_param *ItemDiscountActivityOperateRequest
}

// NewAlibabaRetailMarketingItemdiscountActivityDeleteRequest 初始化AlibabaRetailMarketingItemdiscountActivityDeleteAPIRequest对象
func NewAlibabaRetailMarketingItemdiscountActivityDeleteRequest() *AlibabaRetailMarketingItemdiscountActivityDeleteAPIRequest {
	return &AlibabaRetailMarketingItemdiscountActivityDeleteAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaRetailMarketingItemdiscountActivityDeleteAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaRetailMarketingItemdiscountActivityDeleteAPIRequest) GetApiMethodName() string {
	return "alibaba.retail.marketing.itemdiscount.activity.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaRetailMarketingItemdiscountActivityDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaRetailMarketingItemdiscountActivityDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 删除活动参数
func (r *AlibabaRetailMarketingItemdiscountActivityDeleteAPIRequest) SetParam(_param *ItemDiscountActivityOperateRequest) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaRetailMarketingItemdiscountActivityDeleteAPIRequest) GetParam() *ItemDiscountActivityOperateRequest {
	return r._param
}

var poolAlibabaRetailMarketingItemdiscountActivityDeleteAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaRetailMarketingItemdiscountActivityDeleteRequest()
	},
}

// GetAlibabaRetailMarketingItemdiscountActivityDeleteRequest 从 sync.Pool 获取 AlibabaRetailMarketingItemdiscountActivityDeleteAPIRequest
func GetAlibabaRetailMarketingItemdiscountActivityDeleteAPIRequest() *AlibabaRetailMarketingItemdiscountActivityDeleteAPIRequest {
	return poolAlibabaRetailMarketingItemdiscountActivityDeleteAPIRequest.Get().(*AlibabaRetailMarketingItemdiscountActivityDeleteAPIRequest)
}

// ReleaseAlibabaRetailMarketingItemdiscountActivityDeleteAPIRequest 将 AlibabaRetailMarketingItemdiscountActivityDeleteAPIRequest 放入 sync.Pool
func ReleaseAlibabaRetailMarketingItemdiscountActivityDeleteAPIRequest(v *AlibabaRetailMarketingItemdiscountActivityDeleteAPIRequest) {
	v.Reset()
	poolAlibabaRetailMarketingItemdiscountActivityDeleteAPIRequest.Put(v)
}
