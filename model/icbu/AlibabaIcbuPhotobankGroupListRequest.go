package icbu

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
图片银行分组信息获取 APIRequest
alibaba.icbu.photobank.group.list

图片银行分组信息获取
*/
type AlibabaIcbuPhotobankGroupListRequest struct {
    model.Params

    // 补充信息
    extraContext   string 

    // 查询图片分组信息，如果传入id，则获取当前分组和所有子分组信息，否则获取所有一级分组信息
    id   int64 

}

func NewAlibabaIcbuPhotobankGroupListRequest() *AlibabaIcbuPhotobankGroupListRequest{
    return &AlibabaIcbuPhotobankGroupListRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIcbuPhotobankGroupListRequest) GetApiMethodName() string {
    return "alibaba.icbu.photobank.group.list"
}

func (r AlibabaIcbuPhotobankGroupListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIcbuPhotobankGroupListRequest) SetExtraContext(extraContext string) error {
    r.extraContext = extraContext
    r.Set("extra_context", extraContext)
    return nil
}

func (r AlibabaIcbuPhotobankGroupListRequest) GetExtraContext() string {
    return r.extraContext
}

func (r *AlibabaIcbuPhotobankGroupListRequest) SetId(id int64) error {
    r.id = id
    r.Set("id", id)
    return nil
}

func (r AlibabaIcbuPhotobankGroupListRequest) GetId() int64 {
    return r.id
}

