package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
物流企业查询货主企业信息 API请求
alibaba.alihealth.drug.kyt.synonymauths

物流企业查询货主企业信息
*/
type AlibabaAlihealthDrugKytSynonymauthsRequest struct {
    model.Params
    // 企业ID
    _refEntId   string
    // 企业名称
    _entName   string
    // 货主自定义编号
    _synOwnEntId   string
    // 页码
    _pageSize   int64
    // 页面大小
    _page   int64
}

// 初始化AlibabaAlihealthDrugKytSynonymauthsRequest对象
func NewAlibabaAlihealthDrugKytSynonymauthsRequest() *AlibabaAlihealthDrugKytSynonymauthsRequest{
    return &AlibabaAlihealthDrugKytSynonymauthsRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytSynonymauthsRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.synonymauths"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytSynonymauthsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefEntId Setter
// 企业ID
func (r *AlibabaAlihealthDrugKytSynonymauthsRequest) SetRefEntId(_refEntId string) error {
    r._refEntId = _refEntId
    r.Set("ref_ent_id", _refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugKytSynonymauthsRequest) GetRefEntId() string {
    return r._refEntId
}
// EntName Setter
// 企业名称
func (r *AlibabaAlihealthDrugKytSynonymauthsRequest) SetEntName(_entName string) error {
    r._entName = _entName
    r.Set("ent_name", _entName)
    return nil
}

// EntName Getter
func (r AlibabaAlihealthDrugKytSynonymauthsRequest) GetEntName() string {
    return r._entName
}
// SynOwnEntId Setter
// 货主自定义编号
func (r *AlibabaAlihealthDrugKytSynonymauthsRequest) SetSynOwnEntId(_synOwnEntId string) error {
    r._synOwnEntId = _synOwnEntId
    r.Set("syn_own_ent_id", _synOwnEntId)
    return nil
}

// SynOwnEntId Getter
func (r AlibabaAlihealthDrugKytSynonymauthsRequest) GetSynOwnEntId() string {
    return r._synOwnEntId
}
// PageSize Setter
// 页码
func (r *AlibabaAlihealthDrugKytSynonymauthsRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r AlibabaAlihealthDrugKytSynonymauthsRequest) GetPageSize() int64 {
    return r._pageSize
}
// Page Setter
// 页面大小
func (r *AlibabaAlihealthDrugKytSynonymauthsRequest) SetPage(_page int64) error {
    r._page = _page
    r.Set("page", _page)
    return nil
}

// Page Getter
func (r AlibabaAlihealthDrugKytSynonymauthsRequest) GetPage() int64 {
    return r._page
}
