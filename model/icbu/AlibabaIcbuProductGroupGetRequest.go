package icbu

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分组信息获取 APIRequest
alibaba.icbu.product.group.get

分组信息获取
*/
type AlibabaIcbuProductGroupGetRequest struct {
    model.Params

    // 分组ID，传-1获得所有一级分组
    groupId   int64 

    // 补充信息
    extraContext   string 

}

func NewAlibabaIcbuProductGroupGetRequest() *AlibabaIcbuProductGroupGetRequest{
    return &AlibabaIcbuProductGroupGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIcbuProductGroupGetRequest) GetApiMethodName() string {
    return "alibaba.icbu.product.group.get"
}

func (r AlibabaIcbuProductGroupGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIcbuProductGroupGetRequest) SetGroupId(groupId int64) error {
    r.groupId = groupId
    r.Set("group_id", groupId)
    return nil
}

func (r AlibabaIcbuProductGroupGetRequest) GetGroupId() int64 {
    return r.groupId
}

func (r *AlibabaIcbuProductGroupGetRequest) SetExtraContext(extraContext string) error {
    r.extraContext = extraContext
    r.Set("extra_context", extraContext)
    return nil
}

func (r AlibabaIcbuProductGroupGetRequest) GetExtraContext() string {
    return r.extraContext
}

