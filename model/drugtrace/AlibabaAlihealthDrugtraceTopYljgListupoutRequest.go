package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
医疗机构查询本企业上游企业出库单据信息 APIRequest
alibaba.alihealth.drugtrace.top.yljg.listupout

查询货主/本企业上游企业出库单据信息
*/
type AlibabaAlihealthDrugtraceTopYljgListupoutRequest struct {
    model.Params

    // 企业ID
    refEntId   string 

    // 开始日期（不写时分秒）
    beginDate   string 

    // 结束日期（不写时分秒）
    endDate   string 

    // 生产批号
    produceBatchNo   string 

    // 药品ID
    drugEntBaseInfoId   string 

    // 单据类型
    billType   string 

    // 药品类型
    physicType   string 

    // 状态
    status   string 

    // 单据号
    billCode   string 

    // 页大小
    pageSize   int64 

    // 页码
    page   int64 

    // 发货单位
    fromUserId   string 

}

func NewAlibabaAlihealthDrugtraceTopYljgListupoutRequest() *AlibabaAlihealthDrugtraceTopYljgListupoutRequest{
    return &AlibabaAlihealthDrugtraceTopYljgListupoutRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugtraceTopYljgListupoutRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drugtrace.top.yljg.listupout"
}

func (r AlibabaAlihealthDrugtraceTopYljgListupoutRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugtraceTopYljgListupoutRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopYljgListupoutRequest) GetRefEntId() string {
    return r.refEntId
}

func (r *AlibabaAlihealthDrugtraceTopYljgListupoutRequest) SetBeginDate(beginDate string) error {
    r.beginDate = beginDate
    r.Set("begin_date", beginDate)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopYljgListupoutRequest) GetBeginDate() string {
    return r.beginDate
}

func (r *AlibabaAlihealthDrugtraceTopYljgListupoutRequest) SetEndDate(endDate string) error {
    r.endDate = endDate
    r.Set("end_date", endDate)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopYljgListupoutRequest) GetEndDate() string {
    return r.endDate
}

func (r *AlibabaAlihealthDrugtraceTopYljgListupoutRequest) SetProduceBatchNo(produceBatchNo string) error {
    r.produceBatchNo = produceBatchNo
    r.Set("produce_batch_no", produceBatchNo)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopYljgListupoutRequest) GetProduceBatchNo() string {
    return r.produceBatchNo
}

func (r *AlibabaAlihealthDrugtraceTopYljgListupoutRequest) SetDrugEntBaseInfoId(drugEntBaseInfoId string) error {
    r.drugEntBaseInfoId = drugEntBaseInfoId
    r.Set("drug_ent_base_info_id", drugEntBaseInfoId)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopYljgListupoutRequest) GetDrugEntBaseInfoId() string {
    return r.drugEntBaseInfoId
}

func (r *AlibabaAlihealthDrugtraceTopYljgListupoutRequest) SetBillType(billType string) error {
    r.billType = billType
    r.Set("bill_type", billType)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopYljgListupoutRequest) GetBillType() string {
    return r.billType
}

func (r *AlibabaAlihealthDrugtraceTopYljgListupoutRequest) SetPhysicType(physicType string) error {
    r.physicType = physicType
    r.Set("physic_type", physicType)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopYljgListupoutRequest) GetPhysicType() string {
    return r.physicType
}

func (r *AlibabaAlihealthDrugtraceTopYljgListupoutRequest) SetStatus(status string) error {
    r.status = status
    r.Set("status", status)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopYljgListupoutRequest) GetStatus() string {
    return r.status
}

func (r *AlibabaAlihealthDrugtraceTopYljgListupoutRequest) SetBillCode(billCode string) error {
    r.billCode = billCode
    r.Set("bill_code", billCode)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopYljgListupoutRequest) GetBillCode() string {
    return r.billCode
}

func (r *AlibabaAlihealthDrugtraceTopYljgListupoutRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopYljgListupoutRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *AlibabaAlihealthDrugtraceTopYljgListupoutRequest) SetPage(page int64) error {
    r.page = page
    r.Set("page", page)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopYljgListupoutRequest) GetPage() int64 {
    return r.page
}

func (r *AlibabaAlihealthDrugtraceTopYljgListupoutRequest) SetFromUserId(fromUserId string) error {
    r.fromUserId = fromUserId
    r.Set("from_user_id", fromUserId)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopYljgListupoutRequest) GetFromUserId() string {
    return r.fromUserId
}

