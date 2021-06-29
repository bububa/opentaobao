package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
企业搜索自己授权的物流企业 API请求
alibaba.alihealth.drug.kyt.listauths

企业搜索自己授权的物流企业
*/
type AlibabaAlihealthDrugKytListauthsRequest struct {
    model.Params
    // 企业ID
    refEntId   string
    // 企业名称
    entName   string
    // 页码
    page   int64
    // 页大小
    pageSize   int64
}

// 初始化AlibabaAlihealthDrugKytListauthsRequest对象
func NewAlibabaAlihealthDrugKytListauthsRequest() *AlibabaAlihealthDrugKytListauthsRequest{
    return &AlibabaAlihealthDrugKytListauthsRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytListauthsRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.listauths"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytListauthsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefEntId Setter
// 企业ID
func (r *AlibabaAlihealthDrugKytListauthsRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugKytListauthsRequest) GetRefEntId() string {
    return r.refEntId
}
// EntName Setter
// 企业名称
func (r *AlibabaAlihealthDrugKytListauthsRequest) SetEntName(entName string) error {
    r.entName = entName
    r.Set("ent_name", entName)
    return nil
}

// EntName Getter
func (r AlibabaAlihealthDrugKytListauthsRequest) GetEntName() string {
    return r.entName
}
// Page Setter
// 页码
func (r *AlibabaAlihealthDrugKytListauthsRequest) SetPage(page int64) error {
    r.page = page
    r.Set("page", page)
    return nil
}

// Page Getter
func (r AlibabaAlihealthDrugKytListauthsRequest) GetPage() int64 {
    return r.page
}
// PageSize Setter
// 页大小
func (r *AlibabaAlihealthDrugKytListauthsRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r AlibabaAlihealthDrugKytListauthsRequest) GetPageSize() int64 {
    return r.pageSize
}
