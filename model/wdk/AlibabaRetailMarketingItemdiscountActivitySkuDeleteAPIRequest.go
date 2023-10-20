package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaretailmarketingitemdiscountactivityskudeleteAPIRequest 删除特价活动商品 API请求
// alibaba.retail.marketing.itemdiscount.activity.sku.delete
//
// 删除活动商品信息【同城零售】
type AlibabaretailmarketingitemdiscountactivityskudeleteAPIRequest struct {
	model.Params
	// 添加活动商品参数
	_param *ItemDiscountActivityElementOperateRequest
}

// NewAlibabaretailmarketingitemdiscountactivityskudeleteRequest 初始化AlibabaretailmarketingitemdiscountactivityskudeleteAPIRequest对象
func NewAlibabaretailmarketingitemdiscountactivityskudeleteRequest() *AlibabaretailmarketingitemdiscountactivityskudeleteAPIRequest {
	return &AlibabaretailmarketingitemdiscountactivityskudeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaretailmarketingitemdiscountactivityskudeleteAPIRequest) GetApiMethodName() string {
	return "alibaba.retail.marketing.itemdiscount.activity.sku.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaretailmarketingitemdiscountactivityskudeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaretailmarketingitemdiscountactivityskudeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 添加活动商品参数
func (r *AlibabaretailmarketingitemdiscountactivityskudeleteAPIRequest) SetParam(_param *ItemDiscountActivityElementOperateRequest) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaretailmarketingitemdiscountactivityskudeleteAPIRequest) GetParam() *ItemDiscountActivityElementOperateRequest {
	return r._param
}
