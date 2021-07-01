package mtopopen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
回传抽奖相关参数 API请求
alibaba.interact.lotteryactivity.register

提供接口供三方应用将数据回传到平台
*/
type AlibabaInteractLotteryactivityRegisterAPIRequest struct {
    model.Params
    // 入参
    _paramTopUpdateActivityLotteryInfoParam   *TopUpdateActivityLotteryInfoParam
}

// 初始化AlibabaInteractLotteryactivityRegisterAPIRequest对象
func NewAlibabaInteractLotteryactivityRegisterRequest() *AlibabaInteractLotteryactivityRegisterAPIRequest{
    return &AlibabaInteractLotteryactivityRegisterAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaInteractLotteryactivityRegisterAPIRequest) GetApiMethodName() string {
    return "alibaba.interact.lotteryactivity.register"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaInteractLotteryactivityRegisterAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamTopUpdateActivityLotteryInfoParam Setter
// 入参
func (r *AlibabaInteractLotteryactivityRegisterAPIRequest) SetParamTopUpdateActivityLotteryInfoParam(_paramTopUpdateActivityLotteryInfoParam *TopUpdateActivityLotteryInfoParam) error {
    r._paramTopUpdateActivityLotteryInfoParam = _paramTopUpdateActivityLotteryInfoParam
    r.Set("param_top_update_activity_lottery_info_param", _paramTopUpdateActivityLotteryInfoParam)
    return nil
}

// ParamTopUpdateActivityLotteryInfoParam Getter
func (r AlibabaInteractLotteryactivityRegisterAPIRequest) GetParamTopUpdateActivityLotteryInfoParam() *TopUpdateActivityLotteryInfoParam {
    return r._paramTopUpdateActivityLotteryInfoParam
}
