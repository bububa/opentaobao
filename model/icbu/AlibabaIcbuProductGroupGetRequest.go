package icbu

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分组信息获取 API请求
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

// 初始化AlibabaIcbuProductGroupGetRequest对象
func NewAlibabaIcbuProductGroupGetRequest() *AlibabaIcbuProductGroupGetRequest{
    return &AlibabaIcbuProductGroupGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIcbuProductGroupGetRequest) GetApiMethodName() string {
    return "alibaba.icbu.product.group.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIcbuProductGroupGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// GroupId Setter
// 分组ID，传-1获得所有一级分组
func (r *AlibabaIcbuProductGroupGetRequest) SetGroupId(groupId int64) error {
    r.groupId = groupId
    r.Set("group_id", groupId)
    return nil
}

// GroupId Getter
func (r AlibabaIcbuProductGroupGetRequest) GetGroupId() int64 {
    return r.groupId
}
// ExtraContext Setter
// 补充信息
func (r *AlibabaIcbuProductGroupGetRequest) SetExtraContext(extraContext string) error {
    r.extraContext = extraContext
    r.Set("extra_context", extraContext)
    return nil
}

// ExtraContext Getter
func (r AlibabaIcbuProductGroupGetRequest) GetExtraContext() string {
    return r.extraContext
}
