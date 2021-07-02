package wdk

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaRetailMarketingItemdiscountActivityUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.retail.marketing.itemdiscount.activity.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaRetailMarketingItemdiscountActivityUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Param Setter
// 创建活动参数
func (r *AlibabaRetailMarketingItemdiscountActivityUpdateAPIRequest) SetParam(_param *ItemDiscountActivityOperateRequest) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// Get Param Getter
func (r AlibabaRetailMarketingItemdiscountActivityUpdateAPIRequest) GetParam() *ItemDiscountActivityOperateRequest {
	return r._param
}
