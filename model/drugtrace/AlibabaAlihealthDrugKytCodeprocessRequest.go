package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
关联关系处理查询 APIRequest
alibaba.alihealth.drug.kyt.codeprocess

关联关系处理查询
*/
type AlibabaAlihealthDrugKytCodeprocessRequest struct {
    model.Params

    // 开始时间
    startDate   string 

    // 结束时间
    endDate   string 

    // 上传标识
    uploadFlag   string 

    // 处理状态
    processFlag   string 

    // 批次号
    produceBatchNo   string 

    // 查询标识
    queryFlag   string 

    // 药品类型
    physicType   string 

    // 生产企业ID
    prodSeqNo   string 

    // 药品ID
    drugEntBaseInfoId   string 

    // 包装规格
    pkgSpec   string 

    // 页数
    page   int64 

    // 条数
    pageSize   int64 

    // 客户端
    clientType   string 

    // 企业ID
    refEntId   string 

}

func NewAlibabaAlihealthDrugKytCodeprocessRequest() *AlibabaAlihealthDrugKytCodeprocessRequest{
    return &AlibabaAlihealthDrugKytCodeprocessRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugKytCodeprocessRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.codeprocess"
}

func (r AlibabaAlihealthDrugKytCodeprocessRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugKytCodeprocessRequest) SetStartDate(startDate string) error {
    r.startDate = startDate
    r.Set("start_date", startDate)
    return nil
}

func (r AlibabaAlihealthDrugKytCodeprocessRequest) GetStartDate() string {
    return r.startDate
}

func (r *AlibabaAlihealthDrugKytCodeprocessRequest) SetEndDate(endDate string) error {
    r.endDate = endDate
    r.Set("end_date", endDate)
    return nil
}

func (r AlibabaAlihealthDrugKytCodeprocessRequest) GetEndDate() string {
    return r.endDate
}

func (r *AlibabaAlihealthDrugKytCodeprocessRequest) SetUploadFlag(uploadFlag string) error {
    r.uploadFlag = uploadFlag
    r.Set("upload_flag", uploadFlag)
    return nil
}

func (r AlibabaAlihealthDrugKytCodeprocessRequest) GetUploadFlag() string {
    return r.uploadFlag
}

func (r *AlibabaAlihealthDrugKytCodeprocessRequest) SetProcessFlag(processFlag string) error {
    r.processFlag = processFlag
    r.Set("process_flag", processFlag)
    return nil
}

func (r AlibabaAlihealthDrugKytCodeprocessRequest) GetProcessFlag() string {
    return r.processFlag
}

func (r *AlibabaAlihealthDrugKytCodeprocessRequest) SetProduceBatchNo(produceBatchNo string) error {
    r.produceBatchNo = produceBatchNo
    r.Set("produce_batch_no", produceBatchNo)
    return nil
}

func (r AlibabaAlihealthDrugKytCodeprocessRequest) GetProduceBatchNo() string {
    return r.produceBatchNo
}

func (r *AlibabaAlihealthDrugKytCodeprocessRequest) SetQueryFlag(queryFlag string) error {
    r.queryFlag = queryFlag
    r.Set("query_flag", queryFlag)
    return nil
}

func (r AlibabaAlihealthDrugKytCodeprocessRequest) GetQueryFlag() string {
    return r.queryFlag
}

func (r *AlibabaAlihealthDrugKytCodeprocessRequest) SetPhysicType(physicType string) error {
    r.physicType = physicType
    r.Set("physic_type", physicType)
    return nil
}

func (r AlibabaAlihealthDrugKytCodeprocessRequest) GetPhysicType() string {
    return r.physicType
}

func (r *AlibabaAlihealthDrugKytCodeprocessRequest) SetProdSeqNo(prodSeqNo string) error {
    r.prodSeqNo = prodSeqNo
    r.Set("prod_seq_no", prodSeqNo)
    return nil
}

func (r AlibabaAlihealthDrugKytCodeprocessRequest) GetProdSeqNo() string {
    return r.prodSeqNo
}

func (r *AlibabaAlihealthDrugKytCodeprocessRequest) SetDrugEntBaseInfoId(drugEntBaseInfoId string) error {
    r.drugEntBaseInfoId = drugEntBaseInfoId
    r.Set("drug_ent_base_info_id", drugEntBaseInfoId)
    return nil
}

func (r AlibabaAlihealthDrugKytCodeprocessRequest) GetDrugEntBaseInfoId() string {
    return r.drugEntBaseInfoId
}

func (r *AlibabaAlihealthDrugKytCodeprocessRequest) SetPkgSpec(pkgSpec string) error {
    r.pkgSpec = pkgSpec
    r.Set("pkg_spec", pkgSpec)
    return nil
}

func (r AlibabaAlihealthDrugKytCodeprocessRequest) GetPkgSpec() string {
    return r.pkgSpec
}

func (r *AlibabaAlihealthDrugKytCodeprocessRequest) SetPage(page int64) error {
    r.page = page
    r.Set("page", page)
    return nil
}

func (r AlibabaAlihealthDrugKytCodeprocessRequest) GetPage() int64 {
    return r.page
}

func (r *AlibabaAlihealthDrugKytCodeprocessRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r AlibabaAlihealthDrugKytCodeprocessRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *AlibabaAlihealthDrugKytCodeprocessRequest) SetClientType(clientType string) error {
    r.clientType = clientType
    r.Set("client_type", clientType)
    return nil
}

func (r AlibabaAlihealthDrugKytCodeprocessRequest) GetClientType() string {
    return r.clientType
}

func (r *AlibabaAlihealthDrugKytCodeprocessRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

func (r AlibabaAlihealthDrugKytCodeprocessRequest) GetRefEntId() string {
    return r.refEntId
}

