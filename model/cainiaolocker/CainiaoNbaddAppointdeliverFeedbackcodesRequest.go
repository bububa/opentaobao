package cainiaolocker

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务质量反馈编码列表 API请求
cainiao.nbadd.appointdeliver.feedbackcodes

服务质量反馈编码列表，建议获取数据后缓存在本地，定时刷新即可
*/
type CainiaoNbaddAppointdeliverFeedbackcodesRequest struct {
    model.Params
}

// 初始化CainiaoNbaddAppointdeliverFeedbackcodesRequest对象
func NewCainiaoNbaddAppointdeliverFeedbackcodesRequest() *CainiaoNbaddAppointdeliverFeedbackcodesRequest{
    return &CainiaoNbaddAppointdeliverFeedbackcodesRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoNbaddAppointdeliverFeedbackcodesRequest) GetApiMethodName() string {
    return "cainiao.nbadd.appointdeliver.feedbackcodes"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoNbaddAppointdeliverFeedbackcodesRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
