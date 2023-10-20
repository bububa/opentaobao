package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaretailmarketingitemdiscountactivityupdateAPIRequest 更新单品特价活动【同城零售】 API请求
// alibaba.retail.marketing.itemdiscount.activity.update
//
// 同城零售单品特价活动更新
type AlibabaretailmarketingitemdiscountactivityupdateAPIRequest struct {
	model.Params
	// 创建活动参数
	_param *ItemDiscountActivityOperateRequest
}

// NewAlibabaretailmarketingitemdiscountactivityupdateRequest 初始化AlibabaretailmarketingitemdiscountactivityupdateAPIRequest对象
func NewAlibabaretailmarketingitemdiscountactivityupdateRequest() *AlibabaretailmarketingitemdiscountactivityupdateAPIRequest {
	return &AlibabaretailmarketingitemdiscountactivityupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaretailmarketingitemdiscountactivityupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.retail.marketing.itemdiscount.activity.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaretailmarketingitemdiscountactivityupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaretailmarketingitemdiscountactivityupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 创建活动参数
func (r *AlibabaretailmarketingitemdiscountactivityupdateAPIRequest) SetParam(_param *ItemDiscountActivityOperateRequest) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaretailmarketingitemdiscountactivityupdateAPIRequest) GetParam() *ItemDiscountActivityOperateRequest {
	return r._param
}
