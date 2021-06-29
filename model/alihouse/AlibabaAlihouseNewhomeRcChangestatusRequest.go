package alihouse

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
图文草稿状态更新 APIRequest
alibaba.alihouse.newhome.rc.changestatus

图文草稿状态更新
*/
type AlibabaAlihouseNewhomeRcChangestatusRequest struct {
    model.Params

    // 外部图文id
    outerId   string 

    // 0 失效 1 有效
    status   int64 

}

func NewAlibabaAlihouseNewhomeRcChangestatusRequest() *AlibabaAlihouseNewhomeRcChangestatusRequest{
    return &AlibabaAlihouseNewhomeRcChangestatusRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihouseNewhomeRcChangestatusRequest) GetApiMethodName() string {
    return "alibaba.alihouse.newhome.rc.changestatus"
}

func (r AlibabaAlihouseNewhomeRcChangestatusRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihouseNewhomeRcChangestatusRequest) SetOuterId(outerId string) error {
    r.outerId = outerId
    r.Set("outer_id", outerId)
    return nil
}

func (r AlibabaAlihouseNewhomeRcChangestatusRequest) GetOuterId() string {
    return r.outerId
}

func (r *AlibabaAlihouseNewhomeRcChangestatusRequest) SetStatus(status int64) error {
    r.status = status
    r.Set("status", status)
    return nil
}

func (r AlibabaAlihouseNewhomeRcChangestatusRequest) GetStatus() int64 {
    return r.status
}

