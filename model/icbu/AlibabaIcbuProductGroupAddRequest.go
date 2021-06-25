package icbu

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
增加商品分组 APIRequest
alibaba.icbu.product.group.add

增加商品分组
*/
type AlibabaIcbuProductGroupAddRequest struct {
    model.Params

    // 分组名称
    groupName   string 

    // 上级分组ID，如果建立顶级分组设为-1
    parentId   int64 

    // 补充信息，如isv id
    extraContext   string 

}

func NewAlibabaIcbuProductGroupAddRequest() *AlibabaIcbuProductGroupAddRequest{
    return &AlibabaIcbuProductGroupAddRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIcbuProductGroupAddRequest) GetApiMethodName() string {
    return "alibaba.icbu.product.group.add"
}

func (r AlibabaIcbuProductGroupAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIcbuProductGroupAddRequest) SetGroupName(groupName string) error {
    r.groupName = groupName
    r.Set("group_name", groupName)
    return nil
}

func (r AlibabaIcbuProductGroupAddRequest) GetGroupName() string {
    return r.groupName
}

func (r *AlibabaIcbuProductGroupAddRequest) SetParentId(parentId int64) error {
    r.parentId = parentId
    r.Set("parent_id", parentId)
    return nil
}

func (r AlibabaIcbuProductGroupAddRequest) GetParentId() int64 {
    return r.parentId
}

func (r *AlibabaIcbuProductGroupAddRequest) SetExtraContext(extraContext string) error {
    r.extraContext = extraContext
    r.Set("extra_context", extraContext)
    return nil
}

func (r AlibabaIcbuProductGroupAddRequest) GetExtraContext() string {
    return r.extraContext
}

