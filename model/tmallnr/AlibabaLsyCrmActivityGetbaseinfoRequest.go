package tmallnr

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
ISV查询活动 API请求
alibaba.lsy.crm.activity.getbaseinfo

ISV查询活动
*/
type AlibabaLsyCrmActivityGetbaseinfoRequest struct {
    model.Params
    // 入参
    nrtQueryActivityReq   *NrtQueryActivityReq
}

// 初始化AlibabaLsyCrmActivityGetbaseinfoRequest对象
func NewAlibabaLsyCrmActivityGetbaseinfoRequest() *AlibabaLsyCrmActivityGetbaseinfoRequest{
    return &AlibabaLsyCrmActivityGetbaseinfoRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLsyCrmActivityGetbaseinfoRequest) GetApiMethodName() string {
    return "alibaba.lsy.crm.activity.getbaseinfo"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLsyCrmActivityGetbaseinfoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// NrtQueryActivityReq Setter
// 入参
func (r *AlibabaLsyCrmActivityGetbaseinfoRequest) SetNrtQueryActivityReq(nrtQueryActivityReq *NrtQueryActivityReq) error {
    r.nrtQueryActivityReq = nrtQueryActivityReq
    r.Set("nrt_query_activity_req", nrtQueryActivityReq)
    return nil
}

// NrtQueryActivityReq Getter
func (r AlibabaLsyCrmActivityGetbaseinfoRequest) GetNrtQueryActivityReq() *NrtQueryActivityReq {
    return r.nrtQueryActivityReq
}
