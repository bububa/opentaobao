package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaretailmarketingbuygiftactivitysaveAPIRequest 【同城零售】单品买赠活动保存 API请求
// alibaba.retail.marketing.buygift.activity.save
//
// 同城零售单品买赠活动保存
type AlibabaretailmarketingbuygiftactivitysaveAPIRequest struct {
	model.Params
	// 保存单品买赠活动参数
	_param *BuyGiftActivityOperateRequest
}

// NewAlibabaretailmarketingbuygiftactivitysaveRequest 初始化AlibabaretailmarketingbuygiftactivitysaveAPIRequest对象
func NewAlibabaretailmarketingbuygiftactivitysaveRequest() *AlibabaretailmarketingbuygiftactivitysaveAPIRequest {
	return &AlibabaretailmarketingbuygiftactivitysaveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaretailmarketingbuygiftactivitysaveAPIRequest) GetApiMethodName() string {
	return "alibaba.retail.marketing.buygift.activity.save"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaretailmarketingbuygiftactivitysaveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaretailmarketingbuygiftactivitysaveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 保存单品买赠活动参数
func (r *AlibabaretailmarketingbuygiftactivitysaveAPIRequest) SetParam(_param *BuyGiftActivityOperateRequest) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaretailmarketingbuygiftactivitysaveAPIRequest) GetParam() *BuyGiftActivityOperateRequest {
	return r._param
}
