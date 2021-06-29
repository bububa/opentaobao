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
type AlibabaInteractLotteryactivityRegisterRequest struct {
    model.Params
    // 入参
    paramTopUpdateActivityLotteryInfoParam   *TopUpdateActivityLotteryInfoParam
}

// 初始化AlibabaInteractLotteryactivityRegisterRequest对象
func NewAlibabaInteractLotteryactivityRegisterRequest() *AlibabaInteractLotteryactivityRegisterRequest{
    return &AlibabaInteractLotteryactivityRegisterRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaInteractLotteryactivityRegisterRequest) GetApiMethodName() string {
    return "alibaba.interact.lotteryactivity.register"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaInteractLotteryactivityRegisterRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamTopUpdateActivityLotteryInfoParam Setter
// 入参
func (r *AlibabaInteractLotteryactivityRegisterRequest) SetParamTopUpdateActivityLotteryInfoParam(paramTopUpdateActivityLotteryInfoParam *TopUpdateActivityLotteryInfoParam) error {
    r.paramTopUpdateActivityLotteryInfoParam = paramTopUpdateActivityLotteryInfoParam
    r.Set("param_top_update_activity_lottery_info_param", paramTopUpdateActivityLotteryInfoParam)
    return nil
}

// ParamTopUpdateActivityLotteryInfoParam Getter
func (r AlibabaInteractLotteryactivityRegisterRequest) GetParamTopUpdateActivityLotteryInfoParam() *TopUpdateActivityLotteryInfoParam {
    return r.paramTopUpdateActivityLotteryInfoParam
}
