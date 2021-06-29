package tmallnr

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
ISV活动页面创建修改 APIRequest
alibaba.lsy.crm.activity.page.update

ISV活动页面创建修改
*/
type AlibabaLsyCrmActivityPageUpdateRequest struct {
    model.Params

    // 入参
    nrtCrmActivityPageCreateReq   *NrtCrmActivityPageCreateReq 

}

func NewAlibabaLsyCrmActivityPageUpdateRequest() *AlibabaLsyCrmActivityPageUpdateRequest{
    return &AlibabaLsyCrmActivityPageUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaLsyCrmActivityPageUpdateRequest) GetApiMethodName() string {
    return "alibaba.lsy.crm.activity.page.update"
}

func (r AlibabaLsyCrmActivityPageUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaLsyCrmActivityPageUpdateRequest) SetNrtCrmActivityPageCreateReq(nrtCrmActivityPageCreateReq *NrtCrmActivityPageCreateReq) error {
    r.nrtCrmActivityPageCreateReq = nrtCrmActivityPageCreateReq
    r.Set("nrt_crm_activity_page_create_req", nrtCrmActivityPageCreateReq)
    return nil
}

func (r AlibabaLsyCrmActivityPageUpdateRequest) GetNrtCrmActivityPageCreateReq() *NrtCrmActivityPageCreateReq {
    return r.nrtCrmActivityPageCreateReq
}

