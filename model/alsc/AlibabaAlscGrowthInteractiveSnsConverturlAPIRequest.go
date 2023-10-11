package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscGrowthInteractiveSnsConverturlAPIRequest 防封接口 API请求
// alibaba.alsc.growth.interactive.sns.converturl
//
// 防封接口
type AlibabaAlscGrowthInteractiveSnsConverturlAPIRequest struct {
	model.Params
	// 入参
	_param *FcUrlRequest
}

// NewAlibabaAlscGrowthInteractiveSnsConverturlRequest 初始化AlibabaAlscGrowthInteractiveSnsConverturlAPIRequest对象
func NewAlibabaAlscGrowthInteractiveSnsConverturlRequest() *AlibabaAlscGrowthInteractiveSnsConverturlAPIRequest {
	return &AlibabaAlscGrowthInteractiveSnsConverturlAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscGrowthInteractiveSnsConverturlAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.growth.interactive.sns.converturl"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscGrowthInteractiveSnsConverturlAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlscGrowthInteractiveSnsConverturlAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 入参
func (r *AlibabaAlscGrowthInteractiveSnsConverturlAPIRequest) SetParam(_param *FcUrlRequest) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaAlscGrowthInteractiveSnsConverturlAPIRequest) GetParam() *FcUrlRequest {
	return r._param
}
