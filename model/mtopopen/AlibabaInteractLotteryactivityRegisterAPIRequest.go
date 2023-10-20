package mtopopen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabainteractlotteryactivityregisterAPIRequest 回传抽奖相关参数 API请求
// alibaba.interact.lotteryactivity.register
//
// 提供接口供三方应用将数据回传到平台
type AlibabainteractlotteryactivityregisterAPIRequest struct {
	model.Params
	// 入参
	_paramTopUpdateActivityLotteryInfoParam *TopUpdateActivityLotteryInfoParam
}

// NewAlibabainteractlotteryactivityregisterRequest 初始化AlibabainteractlotteryactivityregisterAPIRequest对象
func NewAlibabainteractlotteryactivityregisterRequest() *AlibabainteractlotteryactivityregisterAPIRequest {
	return &AlibabainteractlotteryactivityregisterAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabainteractlotteryactivityregisterAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.lotteryactivity.register"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabainteractlotteryactivityregisterAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabainteractlotteryactivityregisterAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamTopUpdateActivityLotteryInfoParam is ParamTopUpdateActivityLotteryInfoParam Setter
// 入参
func (r *AlibabainteractlotteryactivityregisterAPIRequest) SetParamTopUpdateActivityLotteryInfoParam(_paramTopUpdateActivityLotteryInfoParam *TopUpdateActivityLotteryInfoParam) error {
	r._paramTopUpdateActivityLotteryInfoParam = _paramTopUpdateActivityLotteryInfoParam
	r.Set("param_top_update_activity_lottery_info_param", _paramTopUpdateActivityLotteryInfoParam)
	return nil
}

// GetParamTopUpdateActivityLotteryInfoParam ParamTopUpdateActivityLotteryInfoParam Getter
func (r AlibabainteractlotteryactivityregisterAPIRequest) GetParamTopUpdateActivityLotteryInfoParam() *TopUpdateActivityLotteryInfoParam {
	return r._paramTopUpdateActivityLotteryInfoParam
}
