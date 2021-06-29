package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询指定产品分组的下一层子分组 API请求
alibaba.scbp.product.group.get

查询指定产品分组的下一层子分组
*/
type AlibabaScbpProductGroupGetRequest struct {
    model.Params
    // 产品分组标识，null表示查询第一层分组
    groupId   string
}

// 初始化AlibabaScbpProductGroupGetRequest对象
func NewAlibabaScbpProductGroupGetRequest() *AlibabaScbpProductGroupGetRequest{
    return &AlibabaScbpProductGroupGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpProductGroupGetRequest) GetApiMethodName() string {
    return "alibaba.scbp.product.group.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpProductGroupGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// GroupId Setter
// 产品分组标识，null表示查询第一层分组
func (r *AlibabaScbpProductGroupGetRequest) SetGroupId(groupId string) error {
    r.groupId = groupId
    r.Set("group_id", groupId)
    return nil
}

// GroupId Getter
func (r AlibabaScbpProductGroupGetRequest) GetGroupId() string {
    return r.groupId
}
