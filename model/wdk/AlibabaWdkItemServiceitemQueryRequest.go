package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询服务商品 APIRequest
alibaba.wdk.item.serviceitem.query

查询服务商品
*/
type AlibabaWdkItemServiceitemQueryRequest struct {
    model.Params

    // 后台类目
    hemaCategoryId   string 

    // 机构编码
    orgNo   string 

}

func NewAlibabaWdkItemServiceitemQueryRequest() *AlibabaWdkItemServiceitemQueryRequest{
    return &AlibabaWdkItemServiceitemQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkItemServiceitemQueryRequest) GetApiMethodName() string {
    return "alibaba.wdk.item.serviceitem.query"
}

func (r AlibabaWdkItemServiceitemQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkItemServiceitemQueryRequest) SetHemaCategoryId(hemaCategoryId string) error {
    r.hemaCategoryId = hemaCategoryId
    r.Set("hema_category_id", hemaCategoryId)
    return nil
}

func (r AlibabaWdkItemServiceitemQueryRequest) GetHemaCategoryId() string {
    return r.hemaCategoryId
}

func (r *AlibabaWdkItemServiceitemQueryRequest) SetOrgNo(orgNo string) error {
    r.orgNo = orgNo
    r.Set("org_no", orgNo)
    return nil
}

func (r AlibabaWdkItemServiceitemQueryRequest) GetOrgNo() string {
    return r.orgNo
}

