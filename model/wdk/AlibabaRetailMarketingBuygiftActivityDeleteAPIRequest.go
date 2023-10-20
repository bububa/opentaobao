package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaretailmarketingbuygiftactivitydeleteAPIRequest 删除单品买赠活动 API请求
// alibaba.retail.marketing.buygift.activity.delete
//
// 同城零售单品特价活动删除
type AlibabaretailmarketingbuygiftactivitydeleteAPIRequest struct {
	model.Params
	// 删除活动参数
	_param *ItemDiscountActivityOperateRequest
}

// NewAlibabaretailmarketingbuygiftactivitydeleteRequest 初始化AlibabaretailmarketingbuygiftactivitydeleteAPIRequest对象
func NewAlibabaretailmarketingbuygiftactivitydeleteRequest() *AlibabaretailmarketingbuygiftactivitydeleteAPIRequest {
	return &AlibabaretailmarketingbuygiftactivitydeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaretailmarketingbuygiftactivitydeleteAPIRequest) GetApiMethodName() string {
	return "alibaba.retail.marketing.buygift.activity.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaretailmarketingbuygiftactivitydeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaretailmarketingbuygiftactivitydeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 删除活动参数
func (r *AlibabaretailmarketingbuygiftactivitydeleteAPIRequest) SetParam(_param *ItemDiscountActivityOperateRequest) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaretailmarketingbuygiftactivitydeleteAPIRequest) GetParam() *ItemDiscountActivityOperateRequest {
	return r._param
}
