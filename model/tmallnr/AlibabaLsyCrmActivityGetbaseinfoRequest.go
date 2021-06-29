package tmallnr

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
ISV查询活动 APIRequest
alibaba.lsy.crm.activity.getbaseinfo

ISV查询活动
*/
type AlibabaLsyCrmActivityGetbaseinfoRequest struct {
    model.Params

    // 入参
    nrtQueryActivityReq   *NrtQueryActivityReq 

}

func NewAlibabaLsyCrmActivityGetbaseinfoRequest() *AlibabaLsyCrmActivityGetbaseinfoRequest{
    return &AlibabaLsyCrmActivityGetbaseinfoRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaLsyCrmActivityGetbaseinfoRequest) GetApiMethodName() string {
    return "alibaba.lsy.crm.activity.getbaseinfo"
}

func (r AlibabaLsyCrmActivityGetbaseinfoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaLsyCrmActivityGetbaseinfoRequest) SetNrtQueryActivityReq(nrtQueryActivityReq *NrtQueryActivityReq) error {
    r.nrtQueryActivityReq = nrtQueryActivityReq
    r.Set("nrt_query_activity_req", nrtQueryActivityReq)
    return nil
}

func (r AlibabaLsyCrmActivityGetbaseinfoRequest) GetNrtQueryActivityReq() *NrtQueryActivityReq {
    return r.nrtQueryActivityReq
}

