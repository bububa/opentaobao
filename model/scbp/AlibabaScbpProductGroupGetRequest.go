package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询指定产品分组的下一层子分组 APIRequest
alibaba.scbp.product.group.get

查询指定产品分组的下一层子分组
*/
type AlibabaScbpProductGroupGetRequest struct {
    model.Params

    // 产品分组标识，null表示查询第一层分组
    groupId   string 

}

func NewAlibabaScbpProductGroupGetRequest() *AlibabaScbpProductGroupGetRequest{
    return &AlibabaScbpProductGroupGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaScbpProductGroupGetRequest) GetApiMethodName() string {
    return "alibaba.scbp.product.group.get"
}

func (r AlibabaScbpProductGroupGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaScbpProductGroupGetRequest) SetGroupId(groupId string) error {
    r.groupId = groupId
    r.Set("group_id", groupId)
    return nil
}

func (r AlibabaScbpProductGroupGetRequest) GetGroupId() string {
    return r.groupId
}

