package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
退款详情 APIRequest
alibaba.wdk.reverse.reversedetail

退款详情
*/
type AlibabaWdkReverseReversedetailRequest struct {
    model.Params

    // 退款单id
    reverseId   string 

}

func NewAlibabaWdkReverseReversedetailRequest() *AlibabaWdkReverseReversedetailRequest{
    return &AlibabaWdkReverseReversedetailRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkReverseReversedetailRequest) GetApiMethodName() string {
    return "alibaba.wdk.reverse.reversedetail"
}

func (r AlibabaWdkReverseReversedetailRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkReverseReversedetailRequest) SetReverseId(reverseId string) error {
    r.reverseId = reverseId
    r.Set("reverse_id", reverseId)
    return nil
}

func (r AlibabaWdkReverseReversedetailRequest) GetReverseId() string {
    return r.reverseId
}

