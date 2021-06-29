package westcrm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询是否是亲橙里会员 APIRequest
alibaba.westcrm.is.member.get

根据淘宝Id查询是否是亲橙里会员
*/
type AlibabaWestcrmIsMemberGetRequest struct {
    model.Params

}

func NewAlibabaWestcrmIsMemberGetRequest() *AlibabaWestcrmIsMemberGetRequest{
    return &AlibabaWestcrmIsMemberGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWestcrmIsMemberGetRequest) GetApiMethodName() string {
    return "alibaba.westcrm.is.member.get"
}

func (r AlibabaWestcrmIsMemberGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


