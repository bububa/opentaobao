package westcrm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询是否是亲橙里会员 API请求
alibaba.westcrm.is.member.get

根据淘宝Id查询是否是亲橙里会员
*/
type AlibabaWestcrmIsMemberGetRequest struct {
    model.Params
}

// 初始化AlibabaWestcrmIsMemberGetRequest对象
func NewAlibabaWestcrmIsMemberGetRequest() *AlibabaWestcrmIsMemberGetRequest{
    return &AlibabaWestcrmIsMemberGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWestcrmIsMemberGetRequest) GetApiMethodName() string {
    return "alibaba.westcrm.is.member.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWestcrmIsMemberGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
