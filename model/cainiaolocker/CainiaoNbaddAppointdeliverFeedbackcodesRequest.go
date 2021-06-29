package cainiaolocker

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务质量反馈编码列表 APIRequest
cainiao.nbadd.appointdeliver.feedbackcodes

服务质量反馈编码列表，建议获取数据后缓存在本地，定时刷新即可
*/
type CainiaoNbaddAppointdeliverFeedbackcodesRequest struct {
    model.Params

}

func NewCainiaoNbaddAppointdeliverFeedbackcodesRequest() *CainiaoNbaddAppointdeliverFeedbackcodesRequest{
    return &CainiaoNbaddAppointdeliverFeedbackcodesRequest{
        Params: model.NewParams(),
    }
}

func (r CainiaoNbaddAppointdeliverFeedbackcodesRequest) GetApiMethodName() string {
    return "cainiao.nbadd.appointdeliver.feedbackcodes"
}

func (r CainiaoNbaddAppointdeliverFeedbackcodesRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


