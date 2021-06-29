package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询P4P产品 API请求
alibaba.scbp.product.list

查询P4P产品
*/
type AlibabaScbpProductListRequest struct {
    model.Params
    // 产品分组标识
    groupId   string
    // 产品分页查询，每页个数，最大值20
    perPageSize   int64
    // 第几页
    toPage   int64
}

// 初始化AlibabaScbpProductListRequest对象
func NewAlibabaScbpProductListRequest() *AlibabaScbpProductListRequest{
    return &AlibabaScbpProductListRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpProductListRequest) GetApiMethodName() string {
    return "alibaba.scbp.product.list"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpProductListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// GroupId Setter
// 产品分组标识
func (r *AlibabaScbpProductListRequest) SetGroupId(groupId string) error {
    r.groupId = groupId
    r.Set("group_id", groupId)
    return nil
}

// GroupId Getter
func (r AlibabaScbpProductListRequest) GetGroupId() string {
    return r.groupId
}
// PerPageSize Setter
// 产品分页查询，每页个数，最大值20
func (r *AlibabaScbpProductListRequest) SetPerPageSize(perPageSize int64) error {
    r.perPageSize = perPageSize
    r.Set("per_page_size", perPageSize)
    return nil
}

// PerPageSize Getter
func (r AlibabaScbpProductListRequest) GetPerPageSize() int64 {
    return r.perPageSize
}
// ToPage Setter
// 第几页
func (r *AlibabaScbpProductListRequest) SetToPage(toPage int64) error {
    r.toPage = toPage
    r.Set("to_page", toPage)
    return nil
}

// ToPage Getter
func (r AlibabaScbpProductListRequest) GetToPage() int64 {
    return r.toPage
}
