package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
同城零售履约异常中心异常单处理结果回调接口 APIRequest
tmall.cityretail.fulfill.abnormal.center.abnormal.status.change

同城零售履约异常中心异常单处理结果回调接口
*/
type TmallCityretailFulfillAbnormalCenterAbnormalStatusChangeRequest struct {
    model.Params

    // 入参
    abnormalStatusChangeDto   []AbnormalStatusChangeDto 

}

func NewTmallCityretailFulfillAbnormalCenterAbnormalStatusChangeRequest() *TmallCityretailFulfillAbnormalCenterAbnormalStatusChangeRequest{
    return &TmallCityretailFulfillAbnormalCenterAbnormalStatusChangeRequest{
        Params: model.NewParams(),
    }
}

func (r TmallCityretailFulfillAbnormalCenterAbnormalStatusChangeRequest) GetApiMethodName() string {
    return "tmall.cityretail.fulfill.abnormal.center.abnormal.status.change"
}

func (r TmallCityretailFulfillAbnormalCenterAbnormalStatusChangeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallCityretailFulfillAbnormalCenterAbnormalStatusChangeRequest) SetAbnormalStatusChangeDto(abnormalStatusChangeDto []AbnormalStatusChangeDto) error {
    r.abnormalStatusChangeDto = abnormalStatusChangeDto
    r.Set("abnormal_status_change_dto", abnormalStatusChangeDto)
    return nil
}

func (r TmallCityretailFulfillAbnormalCenterAbnormalStatusChangeRequest) GetAbnormalStatusChangeDto() []AbnormalStatusChangeDto {
    return r.abnormalStatusChangeDto
}

