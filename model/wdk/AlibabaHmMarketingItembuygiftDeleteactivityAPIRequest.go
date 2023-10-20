package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingItembuygiftDeleteactivityAPIRequest 删除买赠活动 API请求
// alibaba.hm.marketing.itembuygift.deleteactivity
//
// 删除买赠活动
type AlibabaHmMarketingItembuygiftDeleteactivityAPIRequest struct {
	model.Params
	// 要删除的活动信息
	_param *CommonActivityParam
}

// NewAlibabaHmMarketingItembuygiftDeleteactivityRequest 初始化AlibabaHmMarketingItembuygiftDeleteactivityAPIRequest对象
func NewAlibabaHmMarketingItembuygiftDeleteactivityRequest() *AlibabaHmMarketingItembuygiftDeleteactivityAPIRequest {
	return &AlibabaHmMarketingItembuygiftDeleteactivityAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHmMarketingItembuygiftDeleteactivityAPIRequest) GetApiMethodName() string {
	return "alibaba.hm.marketing.itembuygift.deleteactivity"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHmMarketingItembuygiftDeleteactivityAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaHmMarketingItembuygiftDeleteactivityAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 要删除的活动信息
func (r *AlibabaHmMarketingItembuygiftDeleteactivityAPIRequest) SetParam(_param *CommonActivityParam) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaHmMarketingItembuygiftDeleteactivityAPIRequest) GetParam() *CommonActivityParam {
	return r._param
}
