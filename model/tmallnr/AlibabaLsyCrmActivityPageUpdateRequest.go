package tmallnr

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
ISV活动页面创建修改 API请求
alibaba.lsy.crm.activity.page.update

ISV活动页面创建修改
*/
type AlibabaLsyCrmActivityPageUpdateRequest struct {
    model.Params
    // 入参
    nrtCrmActivityPageCreateReq   *NrtCrmActivityPageCreateReq
}

// 初始化AlibabaLsyCrmActivityPageUpdateRequest对象
func NewAlibabaLsyCrmActivityPageUpdateRequest() *AlibabaLsyCrmActivityPageUpdateRequest{
    return &AlibabaLsyCrmActivityPageUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLsyCrmActivityPageUpdateRequest) GetApiMethodName() string {
    return "alibaba.lsy.crm.activity.page.update"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLsyCrmActivityPageUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// NrtCrmActivityPageCreateReq Setter
// 入参
func (r *AlibabaLsyCrmActivityPageUpdateRequest) SetNrtCrmActivityPageCreateReq(nrtCrmActivityPageCreateReq *NrtCrmActivityPageCreateReq) error {
    r.nrtCrmActivityPageCreateReq = nrtCrmActivityPageCreateReq
    r.Set("nrt_crm_activity_page_create_req", nrtCrmActivityPageCreateReq)
    return nil
}

// NrtCrmActivityPageCreateReq Getter
func (r AlibabaLsyCrmActivityPageUpdateRequest) GetNrtCrmActivityPageCreateReq() *NrtCrmActivityPageCreateReq {
    return r.nrtCrmActivityPageCreateReq
}
