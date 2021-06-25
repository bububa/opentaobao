package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询P4P产品 APIRequest
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

func NewAlibabaScbpProductListRequest() *AlibabaScbpProductListRequest{
    return &AlibabaScbpProductListRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaScbpProductListRequest) GetApiMethodName() string {
    return "alibaba.scbp.product.list"
}

func (r AlibabaScbpProductListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaScbpProductListRequest) SetGroupId(groupId string) error {
    r.groupId = groupId
    r.Set("group_id", groupId)
    return nil
}

func (r AlibabaScbpProductListRequest) GetGroupId() string {
    return r.groupId
}

func (r *AlibabaScbpProductListRequest) SetPerPageSize(perPageSize int64) error {
    r.perPageSize = perPageSize
    r.Set("per_page_size", perPageSize)
    return nil
}

func (r AlibabaScbpProductListRequest) GetPerPageSize() int64 {
    return r.perPageSize
}

func (r *AlibabaScbpProductListRequest) SetToPage(toPage int64) error {
    r.toPage = toPage
    r.Set("to_page", toPage)
    return nil
}

func (r AlibabaScbpProductListRequest) GetToPage() int64 {
    return r.toPage
}

