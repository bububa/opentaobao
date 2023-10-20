package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabahmmarketingitembuygiftqueryactivityAPIRequest 查询买赠活动 API请求
// alibaba.hm.marketing.itembuygift.queryactivity
//
// 查询买赠活动
type AlibabahmmarketingitembuygiftqueryactivityAPIRequest struct {
	model.Params
	// 查询入参
	_param *CommonActivityParam
}

// NewAlibabahmmarketingitembuygiftqueryactivityRequest 初始化AlibabahmmarketingitembuygiftqueryactivityAPIRequest对象
func NewAlibabahmmarketingitembuygiftqueryactivityRequest() *AlibabahmmarketingitembuygiftqueryactivityAPIRequest {
	return &AlibabahmmarketingitembuygiftqueryactivityAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabahmmarketingitembuygiftqueryactivityAPIRequest) GetApiMethodName() string {
	return "alibaba.hm.marketing.itembuygift.queryactivity"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabahmmarketingitembuygiftqueryactivityAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabahmmarketingitembuygiftqueryactivityAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 查询入参
func (r *AlibabahmmarketingitembuygiftqueryactivityAPIRequest) SetParam(_param *CommonActivityParam) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabahmmarketingitembuygiftqueryactivityAPIRequest) GetParam() *CommonActivityParam {
	return r._param
}
