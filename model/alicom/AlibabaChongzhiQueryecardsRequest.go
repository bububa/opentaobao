package alicom

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询指定商家的可用的话费宝贝接口 APIRequest
alibaba.chongzhi.queryecards

查询指定商家的可用的话费宝贝
*/
type AlibabaChongzhiQueryecardsRequest struct {
    model.Params

    // 号码
    mobile   int64 

    // 来源
    clientSource   string 

}

func NewAlibabaChongzhiQueryecardsRequest() *AlibabaChongzhiQueryecardsRequest{
    return &AlibabaChongzhiQueryecardsRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaChongzhiQueryecardsRequest) GetApiMethodName() string {
    return "alibaba.chongzhi.queryecards"
}

func (r AlibabaChongzhiQueryecardsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaChongzhiQueryecardsRequest) SetMobile(mobile int64) error {
    r.mobile = mobile
    r.Set("mobile", mobile)
    return nil
}

func (r AlibabaChongzhiQueryecardsRequest) GetMobile() int64 {
    return r.mobile
}

func (r *AlibabaChongzhiQueryecardsRequest) SetClientSource(clientSource string) error {
    r.clientSource = clientSource
    r.Set("client_source", clientSource)
    return nil
}

func (r AlibabaChongzhiQueryecardsRequest) GetClientSource() string {
    return r.clientSource
}

