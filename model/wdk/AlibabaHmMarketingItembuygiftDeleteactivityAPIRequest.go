package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabahmmarketingitembuygiftdeleteactivityAPIRequest 删除买赠活动 API请求
// alibaba.hm.marketing.itembuygift.deleteactivity
//
// 删除买赠活动
type AlibabahmmarketingitembuygiftdeleteactivityAPIRequest struct {
	model.Params
	// 要删除的活动信息
	_param *CommonActivityParam
}

// NewAlibabahmmarketingitembuygiftdeleteactivityRequest 初始化AlibabahmmarketingitembuygiftdeleteactivityAPIRequest对象
func NewAlibabahmmarketingitembuygiftdeleteactivityRequest() *AlibabahmmarketingitembuygiftdeleteactivityAPIRequest {
	return &AlibabahmmarketingitembuygiftdeleteactivityAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabahmmarketingitembuygiftdeleteactivityAPIRequest) GetApiMethodName() string {
	return "alibaba.hm.marketing.itembuygift.deleteactivity"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabahmmarketingitembuygiftdeleteactivityAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabahmmarketingitembuygiftdeleteactivityAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 要删除的活动信息
func (r *AlibabahmmarketingitembuygiftdeleteactivityAPIRequest) SetParam(_param *CommonActivityParam) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabahmmarketingitembuygiftdeleteactivityAPIRequest) GetParam() *CommonActivityParam {
	return r._param
}
