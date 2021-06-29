package tmallnr

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
ISV活动修改 API请求
alibaba.lsy.crm.activity.update

ISV活动修改
*/
type AlibabaLsyCrmActivityUpdateRequest struct {
    model.Params
    // 入参
    nrtUpdateActivityReq   *NrtUpdateActivityReq
}

// 初始化AlibabaLsyCrmActivityUpdateRequest对象
func NewAlibabaLsyCrmActivityUpdateRequest() *AlibabaLsyCrmActivityUpdateRequest{
    return &AlibabaLsyCrmActivityUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLsyCrmActivityUpdateRequest) GetApiMethodName() string {
    return "alibaba.lsy.crm.activity.update"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLsyCrmActivityUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// NrtUpdateActivityReq Setter
// 入参
func (r *AlibabaLsyCrmActivityUpdateRequest) SetNrtUpdateActivityReq(nrtUpdateActivityReq *NrtUpdateActivityReq) error {
    r.nrtUpdateActivityReq = nrtUpdateActivityReq
    r.Set("nrt_update_activity_req", nrtUpdateActivityReq)
    return nil
}

// NrtUpdateActivityReq Getter
func (r AlibabaLsyCrmActivityUpdateRequest) GetNrtUpdateActivityReq() *NrtUpdateActivityReq {
    return r.nrtUpdateActivityReq
}
