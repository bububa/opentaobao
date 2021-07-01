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
type AlibabaLsyCrmActivityGetbaseinfoAPIRequest struct {
    model.Params
    // 入参
    _nrtQueryActivityReq   *NrtQueryActivityReq
}

// 初始化AlibabaLsyCrmActivityGetbaseinfoAPIRequest对象
func NewAlibabaLsyCrmActivityGetbaseinfoRequest() *AlibabaLsyCrmActivityGetbaseinfoAPIRequest{
    return &AlibabaLsyCrmActivityGetbaseinfoAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLsyCrmActivityGetbaseinfoAPIRequest) GetApiMethodName() string {
    return "alibaba.lsy.crm.activity.getbaseinfo"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLsyCrmActivityGetbaseinfoAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// NrtQueryActivityReq Setter
// 入参
func (r *AlibabaLsyCrmActivityGetbaseinfoAPIRequest) SetNrtQueryActivityReq(_nrtQueryActivityReq *NrtQueryActivityReq) error {
    r._nrtQueryActivityReq = _nrtQueryActivityReq
    r.Set("nrt_query_activity_req", _nrtQueryActivityReq)
    return nil
}

// NrtQueryActivityReq Getter
func (r AlibabaLsyCrmActivityGetbaseinfoAPIRequest) GetNrtQueryActivityReq() *NrtQueryActivityReq {
    return r._nrtQueryActivityReq
}
