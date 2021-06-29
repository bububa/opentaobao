package alihouse

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
视频草稿状态更新 APIRequest
alibaba.alihouse.newhome.video.changestatus

视频草稿状态更新
*/
type AlibabaAlihouseNewhomeVideoChangestatusRequest struct {
    model.Params

    // 外部视频id
    outerId   string 

    // 0 失效 1 有效
    status   int64 

}

func NewAlibabaAlihouseNewhomeVideoChangestatusRequest() *AlibabaAlihouseNewhomeVideoChangestatusRequest{
    return &AlibabaAlihouseNewhomeVideoChangestatusRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihouseNewhomeVideoChangestatusRequest) GetApiMethodName() string {
    return "alibaba.alihouse.newhome.video.changestatus"
}

func (r AlibabaAlihouseNewhomeVideoChangestatusRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihouseNewhomeVideoChangestatusRequest) SetOuterId(outerId string) error {
    r.outerId = outerId
    r.Set("outer_id", outerId)
    return nil
}

func (r AlibabaAlihouseNewhomeVideoChangestatusRequest) GetOuterId() string {
    return r.outerId
}

func (r *AlibabaAlihouseNewhomeVideoChangestatusRequest) SetStatus(status int64) error {
    r.status = status
    r.Set("status", status)
    return nil
}

func (r AlibabaAlihouseNewhomeVideoChangestatusRequest) GetStatus() int64 {
    return r.status
}

