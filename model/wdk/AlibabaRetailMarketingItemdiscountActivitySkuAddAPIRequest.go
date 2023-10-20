package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaretailmarketingitemdiscountactivityskuaddAPIRequest 特价活动新增商品 API请求
// alibaba.retail.marketing.itemdiscount.activity.sku.add
//
// 新增或更新活动商品信息【同城零售】
type AlibabaretailmarketingitemdiscountactivityskuaddAPIRequest struct {
	model.Params
	// 添加活动商品参数
	_param *ItemDiscountActivityElementOperateRequest
}

// NewAlibabaretailmarketingitemdiscountactivityskuaddRequest 初始化AlibabaretailmarketingitemdiscountactivityskuaddAPIRequest对象
func NewAlibabaretailmarketingitemdiscountactivityskuaddRequest() *AlibabaretailmarketingitemdiscountactivityskuaddAPIRequest {
	return &AlibabaretailmarketingitemdiscountactivityskuaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaretailmarketingitemdiscountactivityskuaddAPIRequest) GetApiMethodName() string {
	return "alibaba.retail.marketing.itemdiscount.activity.sku.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaretailmarketingitemdiscountactivityskuaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaretailmarketingitemdiscountactivityskuaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 添加活动商品参数
func (r *AlibabaretailmarketingitemdiscountactivityskuaddAPIRequest) SetParam(_param *ItemDiscountActivityElementOperateRequest) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaretailmarketingitemdiscountactivityskuaddAPIRequest) GetParam() *ItemDiscountActivityElementOperateRequest {
	return r._param
}
