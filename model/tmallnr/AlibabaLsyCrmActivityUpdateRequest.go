package tmallnr

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
ISV活动修改 APIRequest
alibaba.lsy.crm.activity.update

ISV活动修改
*/
type AlibabaLsyCrmActivityUpdateRequest struct {
    model.Params

    // 入参
    nrtUpdateActivityReq   *NrtUpdateActivityReq 

}

func NewAlibabaLsyCrmActivityUpdateRequest() *AlibabaLsyCrmActivityUpdateRequest{
    return &AlibabaLsyCrmActivityUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaLsyCrmActivityUpdateRequest) GetApiMethodName() string {
    return "alibaba.lsy.crm.activity.update"
}

func (r AlibabaLsyCrmActivityUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaLsyCrmActivityUpdateRequest) SetNrtUpdateActivityReq(nrtUpdateActivityReq *NrtUpdateActivityReq) error {
    r.nrtUpdateActivityReq = nrtUpdateActivityReq
    r.Set("nrt_update_activity_req", nrtUpdateActivityReq)
    return nil
}

func (r AlibabaLsyCrmActivityUpdateRequest) GetNrtUpdateActivityReq() *NrtUpdateActivityReq {
    return r.nrtUpdateActivityReq
}

