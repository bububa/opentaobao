package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaRetailMarketingItemdiscountActivityUpdateAPIRequest 更新单品特价活动【同城零售】 API请求
// alibaba.retail.marketing.itemdiscount.activity.update
//
// 同城零售单品特价活动更新
type AlibabaRetailMarketingItemdiscountActivityUpdateAPIRequest struct {
	model.Params
	// 创建活动参数
	_param *ItemDiscountActivityOperateRequest
}

// NewAlibabaRetailMarketingItemdiscountActivityUpdateRequest 初始化AlibabaRetailMarketingItemdiscountActivityUpdateAPIRequest对象
func NewAlibabaRetailMarketingItemdiscountActivityUpdateRequest() *AlibabaRetailMarketingItemdiscountActivityUpdateAPIRequest {
	return &AlibabaRetailMarketingItemdiscountActivityUpdateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaRetailMarketingItemdiscountActivityUpdateAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaRetailMarketingItemdiscountActivityUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.retail.marketing.itemdiscount.activity.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaRetailMarketingItemdiscountActivityUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaRetailMarketingItemdiscountActivityUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 创建活动参数
func (r *AlibabaRetailMarketingItemdiscountActivityUpdateAPIRequest) SetParam(_param *ItemDiscountActivityOperateRequest) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaRetailMarketingItemdiscountActivityUpdateAPIRequest) GetParam() *ItemDiscountActivityOperateRequest {
	return r._param
}

var poolAlibabaRetailMarketingItemdiscountActivityUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaRetailMarketingItemdiscountActivityUpdateRequest()
	},
}

// GetAlibabaRetailMarketingItemdiscountActivityUpdateRequest 从 sync.Pool 获取 AlibabaRetailMarketingItemdiscountActivityUpdateAPIRequest
func GetAlibabaRetailMarketingItemdiscountActivityUpdateAPIRequest() *AlibabaRetailMarketingItemdiscountActivityUpdateAPIRequest {
	return poolAlibabaRetailMarketingItemdiscountActivityUpdateAPIRequest.Get().(*AlibabaRetailMarketingItemdiscountActivityUpdateAPIRequest)
}

// ReleaseAlibabaRetailMarketingItemdiscountActivityUpdateAPIRequest 将 AlibabaRetailMarketingItemdiscountActivityUpdateAPIRequest 放入 sync.Pool
func ReleaseAlibabaRetailMarketingItemdiscountActivityUpdateAPIRequest(v *AlibabaRetailMarketingItemdiscountActivityUpdateAPIRequest) {
	v.Reset()
	poolAlibabaRetailMarketingItemdiscountActivityUpdateAPIRequest.Put(v)
}
