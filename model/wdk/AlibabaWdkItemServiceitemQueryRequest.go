package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询服务商品 API请求
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

// 初始化AlibabaWdkItemServiceitemQueryRequest对象
func NewAlibabaWdkItemServiceitemQueryRequest() *AlibabaWdkItemServiceitemQueryRequest{
    return &AlibabaWdkItemServiceitemQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkItemServiceitemQueryRequest) GetApiMethodName() string {
    return "alibaba.wdk.item.serviceitem.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkItemServiceitemQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// HemaCategoryId Setter
// 后台类目
func (r *AlibabaWdkItemServiceitemQueryRequest) SetHemaCategoryId(hemaCategoryId string) error {
    r.hemaCategoryId = hemaCategoryId
    r.Set("hema_category_id", hemaCategoryId)
    return nil
}

// HemaCategoryId Getter
func (r AlibabaWdkItemServiceitemQueryRequest) GetHemaCategoryId() string {
    return r.hemaCategoryId
}
// OrgNo Setter
// 机构编码
func (r *AlibabaWdkItemServiceitemQueryRequest) SetOrgNo(orgNo string) error {
    r.orgNo = orgNo
    r.Set("org_no", orgNo)
    return nil
}

// OrgNo Getter
func (r AlibabaWdkItemServiceitemQueryRequest) GetOrgNo() string {
    return r.orgNo
}
