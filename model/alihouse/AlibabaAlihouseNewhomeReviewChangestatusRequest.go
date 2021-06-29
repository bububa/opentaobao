package alihouse

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
楼盘测评草稿状态同步 APIRequest
alibaba.alihouse.newhome.review.changestatus

楼盘测评草稿状态更新
*/
type AlibabaAlihouseNewhomeReviewChangestatusRequest struct {
    model.Params

    // 外部测评id
    outerId   string 

    // 0 失效 1 有效
    status   int64 

}

func NewAlibabaAlihouseNewhomeReviewChangestatusRequest() *AlibabaAlihouseNewhomeReviewChangestatusRequest{
    return &AlibabaAlihouseNewhomeReviewChangestatusRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihouseNewhomeReviewChangestatusRequest) GetApiMethodName() string {
    return "alibaba.alihouse.newhome.review.changestatus"
}

func (r AlibabaAlihouseNewhomeReviewChangestatusRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihouseNewhomeReviewChangestatusRequest) SetOuterId(outerId string) error {
    r.outerId = outerId
    r.Set("outer_id", outerId)
    return nil
}

func (r AlibabaAlihouseNewhomeReviewChangestatusRequest) GetOuterId() string {
    return r.outerId
}

func (r *AlibabaAlihouseNewhomeReviewChangestatusRequest) SetStatus(status int64) error {
    r.status = status
    r.Set("status", status)
    return nil
}

func (r AlibabaAlihouseNewhomeReviewChangestatusRequest) GetStatus() int64 {
    return r.status
}

