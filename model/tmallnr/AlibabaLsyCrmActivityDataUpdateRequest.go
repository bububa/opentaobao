package tmallnr

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
私域导购数据回流接口 APIRequest
alibaba.lsy.crm.activity.data.update

私域导购数据回流接口
*/
type AlibabaLsyCrmActivityDataUpdateRequest struct {
    model.Params

    // 入参对象
    reqDTO   *NrtCrmActivityStatisticsDataReq 

}

func NewAlibabaLsyCrmActivityDataUpdateRequest() *AlibabaLsyCrmActivityDataUpdateRequest{
    return &AlibabaLsyCrmActivityDataUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaLsyCrmActivityDataUpdateRequest) GetApiMethodName() string {
    return "alibaba.lsy.crm.activity.data.update"
}

func (r AlibabaLsyCrmActivityDataUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaLsyCrmActivityDataUpdateRequest) SetReqDTO(reqDTO *NrtCrmActivityStatisticsDataReq) error {
    r.reqDTO = reqDTO
    r.Set("req_d_t_o", reqDTO)
    return nil
}

func (r AlibabaLsyCrmActivityDataUpdateRequest) GetReqDTO() *NrtCrmActivityStatisticsDataReq {
    return r.reqDTO
}

