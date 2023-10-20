package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaretailmarketingbuygiftactivityskudeleteAPIRequest 删除单品买赠活动商品 API请求
// alibaba.retail.marketing.buygift.activity.sku.delete
//
// 删除单品买赠活动商品信息【同城零售】
type AlibabaretailmarketingbuygiftactivityskudeleteAPIRequest struct {
	model.Params
	// 删除买赠活动商品参数
	_param *BuyGiftActivitySkuOperateRequest
}

// NewAlibabaretailmarketingbuygiftactivityskudeleteRequest 初始化AlibabaretailmarketingbuygiftactivityskudeleteAPIRequest对象
func NewAlibabaretailmarketingbuygiftactivityskudeleteRequest() *AlibabaretailmarketingbuygiftactivityskudeleteAPIRequest {
	return &AlibabaretailmarketingbuygiftactivityskudeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaretailmarketingbuygiftactivityskudeleteAPIRequest) GetApiMethodName() string {
	return "alibaba.retail.marketing.buygift.activity.sku.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaretailmarketingbuygiftactivityskudeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaretailmarketingbuygiftactivityskudeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 删除买赠活动商品参数
func (r *AlibabaretailmarketingbuygiftactivityskudeleteAPIRequest) SetParam(_param *BuyGiftActivitySkuOperateRequest) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaretailmarketingbuygiftactivityskudeleteAPIRequest) GetParam() *BuyGiftActivitySkuOperateRequest {
	return r._param
}
