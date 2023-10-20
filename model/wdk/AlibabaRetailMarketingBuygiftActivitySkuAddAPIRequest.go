package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaretailmarketingbuygiftactivityskuaddAPIRequest 添加单品买赠活动商品 API请求
// alibaba.retail.marketing.buygift.activity.sku.add
//
// 新增或更新单品买赠活动商品信息【同城零售】
type AlibabaretailmarketingbuygiftactivityskuaddAPIRequest struct {
	model.Params
	// 添加活动商品参数
	_param *BuyGiftActivitySkuOperateRequest
}

// NewAlibabaretailmarketingbuygiftactivityskuaddRequest 初始化AlibabaretailmarketingbuygiftactivityskuaddAPIRequest对象
func NewAlibabaretailmarketingbuygiftactivityskuaddRequest() *AlibabaretailmarketingbuygiftactivityskuaddAPIRequest {
	return &AlibabaretailmarketingbuygiftactivityskuaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaretailmarketingbuygiftactivityskuaddAPIRequest) GetApiMethodName() string {
	return "alibaba.retail.marketing.buygift.activity.sku.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaretailmarketingbuygiftactivityskuaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaretailmarketingbuygiftactivityskuaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 添加活动商品参数
func (r *AlibabaretailmarketingbuygiftactivityskuaddAPIRequest) SetParam(_param *BuyGiftActivitySkuOperateRequest) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaretailmarketingbuygiftactivityskuaddAPIRequest) GetParam() *BuyGiftActivitySkuOperateRequest {
	return r._param
}
