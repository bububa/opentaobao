package wdk

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaRetailMarketingItemdiscountActivitySaveAPIRequest) GetApiMethodName() string {
	return "alibaba.retail.marketing.itemdiscount.activity.save"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaRetailMarketingItemdiscountActivitySaveAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
