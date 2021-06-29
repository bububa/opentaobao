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
    _refEntId   string
    // 企业名称
    _entName   string
    // 页码
    _page   int64
    // 页大小
    _pageSize   int64
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
func (r *AlibabaAlihealthDrugKytListauthsRequest) SetRefEntId(_refEntId string) error {
    r._refEntId = _refEntId
    r.Set("ref_ent_id", _refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugKytListauthsRequest) GetRefEntId() string {
    return r._refEntId
}
// EntName Setter
// 企业名称
func (r *AlibabaAlihealthDrugKytListauthsRequest) SetEntName(_entName string) error {
    r._entName = _entName
    r.Set("ent_name", _entName)
    return nil
}

// EntName Getter
func (r AlibabaAlihealthDrugKytListauthsRequest) GetEntName() string {
    return r._entName
}
// Page Setter
// 页码
func (r *AlibabaAlihealthDrugKytListauthsRequest) SetPage(_page int64) error {
    r._page = _page
    r.Set("page", _page)
    return nil
}

// Page Getter
func (r AlibabaAlihealthDrugKytListauthsRequest) GetPage() int64 {
    return r._page
}
// PageSize Setter
// 页大小
func (r *AlibabaAlihealthDrugKytListauthsRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r AlibabaAlihealthDrugKytListauthsRequest) GetPageSize() int64 {
    return r._pageSize
}
