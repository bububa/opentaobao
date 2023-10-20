package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaretailmarketingbuygiftactivityupdateAPIRequest 更新单品买赠活动 API请求
// alibaba.retail.marketing.buygift.activity.update
//
// 同城零售单品买赠活动更新
type AlibabaretailmarketingbuygiftactivityupdateAPIRequest struct {
	model.Params
	// 更新单品买赠活动参数
	_param *BuyGiftActivityOperateRequest
}

// NewAlibabaretailmarketingbuygiftactivityupdateRequest 初始化AlibabaretailmarketingbuygiftactivityupdateAPIRequest对象
func NewAlibabaretailmarketingbuygiftactivityupdateRequest() *AlibabaretailmarketingbuygiftactivityupdateAPIRequest {
	return &AlibabaretailmarketingbuygiftactivityupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaretailmarketingbuygiftactivityupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.retail.marketing.buygift.activity.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaretailmarketingbuygiftactivityupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaretailmarketingbuygiftactivityupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 更新单品买赠活动参数
func (r *AlibabaretailmarketingbuygiftactivityupdateAPIRequest) SetParam(_param *BuyGiftActivityOperateRequest) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaretailmarketingbuygiftactivityupdateAPIRequest) GetParam() *BuyGiftActivityOperateRequest {
	return r._param
}
