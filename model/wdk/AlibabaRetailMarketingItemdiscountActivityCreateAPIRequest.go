package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaRetailMarketingItemdiscountActivityCreateAPIRequest 创建单品特价活动【同城零售】 API请求
// alibaba.retail.marketing.itemdiscount.activity.create
//
// 同城零售单品特价活动创建
type AlibabaRetailMarketingItemdiscountActivityCreateAPIRequest struct {
	model.Params
	// 创建活动参数
	_param *ItemDiscountActivityOperateRequest
}

// NewAlibabaRetailMarketingItemdiscountActivityCreateRequest 初始化AlibabaRetailMarketingItemdiscountActivityCreateAPIRequest对象
func NewAlibabaRetailMarketingItemdiscountActivityCreateRequest() *AlibabaRetailMarketingItemdiscountActivityCreateAPIRequest {
	return &AlibabaRetailMarketingItemdiscountActivityCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaRetailMarketingItemdiscountActivityCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.retail.marketing.itemdiscount.activity.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaRetailMarketingItemdiscountActivityCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaRetailMarketingItemdiscountActivityCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 创建活动参数
func (r *AlibabaRetailMarketingItemdiscountActivityCreateAPIRequest) SetParam(_param *ItemDiscountActivityOperateRequest) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaRetailMarketingItemdiscountActivityCreateAPIRequest) GetParam() *ItemDiscountActivityOperateRequest {
	return r._param
}
