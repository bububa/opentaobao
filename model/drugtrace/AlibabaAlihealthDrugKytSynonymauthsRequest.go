package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
物流企业查询货主企业信息 APIRequest
alibaba.alihealth.drug.kyt.synonymauths

物流企业查询货主企业信息
*/
type AlibabaAlihealthDrugKytSynonymauthsRequest struct {
    model.Params

    // 企业ID
    refEntId   string 

    // 企业名称
    entName   string 

    // 货主自定义编号
    synOwnEntId   string 

    // 页码
    pageSize   int64 

    // 页面大小
    page   int64 

}

func NewAlibabaAlihealthDrugKytSynonymauthsRequest() *AlibabaAlihealthDrugKytSynonymauthsRequest{
    return &AlibabaAlihealthDrugKytSynonymauthsRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugKytSynonymauthsRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.synonymauths"
}

func (r AlibabaAlihealthDrugKytSynonymauthsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugKytSynonymauthsRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

func (r AlibabaAlihealthDrugKytSynonymauthsRequest) GetRefEntId() string {
    return r.refEntId
}

func (r *AlibabaAlihealthDrugKytSynonymauthsRequest) SetEntName(entName string) error {
    r.entName = entName
    r.Set("ent_name", entName)
    return nil
}

func (r AlibabaAlihealthDrugKytSynonymauthsRequest) GetEntName() string {
    return r.entName
}

func (r *AlibabaAlihealthDrugKytSynonymauthsRequest) SetSynOwnEntId(synOwnEntId string) error {
    r.synOwnEntId = synOwnEntId
    r.Set("syn_own_ent_id", synOwnEntId)
    return nil
}

func (r AlibabaAlihealthDrugKytSynonymauthsRequest) GetSynOwnEntId() string {
    return r.synOwnEntId
}

func (r *AlibabaAlihealthDrugKytSynonymauthsRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r AlibabaAlihealthDrugKytSynonymauthsRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *AlibabaAlihealthDrugKytSynonymauthsRequest) SetPage(page int64) error {
    r.page = page
    r.Set("page", page)
    return nil
}

func (r AlibabaAlihealthDrugKytSynonymauthsRequest) GetPage() int64 {
    return r.page
}

