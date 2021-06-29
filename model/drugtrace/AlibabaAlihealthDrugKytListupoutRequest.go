package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询货主/本企业上游企业出库单据信息 API请求
alibaba.alihealth.drug.kyt.listupout

查询货主/本企业上游企业出库单据信息
*/
type AlibabaAlihealthDrugKytListupoutRequest struct {
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

// 初始化AlibabaAlihealthDrugKytListupoutRequest对象
func NewAlibabaAlihealthDrugKytListupoutRequest() *AlibabaAlihealthDrugKytListupoutRequest{
    return &AlibabaAlihealthDrugKytListupoutRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytListupoutRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.listupout"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytListupoutRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefEntId Setter
// 企业ID
func (r *AlibabaAlihealthDrugKytListupoutRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugKytListupoutRequest) GetRefEntId() string {
    return r.refEntId
}
// BeginDate Setter
// 开始日期（不写时分秒）
func (r *AlibabaAlihealthDrugKytListupoutRequest) SetBeginDate(beginDate string) error {
    r.beginDate = beginDate
    r.Set("begin_date", beginDate)
    return nil
}

// BeginDate Getter
func (r AlibabaAlihealthDrugKytListupoutRequest) GetBeginDate() string {
    return r.beginDate
}
// EndDate Setter
// 结束日期（不写时分秒）
func (r *AlibabaAlihealthDrugKytListupoutRequest) SetEndDate(endDate string) error {
    r.endDate = endDate
    r.Set("end_date", endDate)
    return nil
}

// EndDate Getter
func (r AlibabaAlihealthDrugKytListupoutRequest) GetEndDate() string {
    return r.endDate
}
// ProduceBatchNo Setter
// 生产批号
func (r *AlibabaAlihealthDrugKytListupoutRequest) SetProduceBatchNo(produceBatchNo string) error {
    r.produceBatchNo = produceBatchNo
    r.Set("produce_batch_no", produceBatchNo)
    return nil
}

// ProduceBatchNo Getter
func (r AlibabaAlihealthDrugKytListupoutRequest) GetProduceBatchNo() string {
    return r.produceBatchNo
}
// DrugEntBaseInfoId Setter
// 药品ID
func (r *AlibabaAlihealthDrugKytListupoutRequest) SetDrugEntBaseInfoId(drugEntBaseInfoId string) error {
    r.drugEntBaseInfoId = drugEntBaseInfoId
    r.Set("drug_ent_base_info_id", drugEntBaseInfoId)
    return nil
}

// DrugEntBaseInfoId Getter
func (r AlibabaAlihealthDrugKytListupoutRequest) GetDrugEntBaseInfoId() string {
    return r.drugEntBaseInfoId
}
// BillType Setter
// 单据类型
func (r *AlibabaAlihealthDrugKytListupoutRequest) SetBillType(billType string) error {
    r.billType = billType
    r.Set("bill_type", billType)
    return nil
}

// BillType Getter
func (r AlibabaAlihealthDrugKytListupoutRequest) GetBillType() string {
    return r.billType
}
// PhysicType Setter
// 药品类型
func (r *AlibabaAlihealthDrugKytListupoutRequest) SetPhysicType(physicType string) error {
    r.physicType = physicType
    r.Set("physic_type", physicType)
    return nil
}

// PhysicType Getter
func (r AlibabaAlihealthDrugKytListupoutRequest) GetPhysicType() string {
    return r.physicType
}
// Status Setter
// 状态
func (r *AlibabaAlihealthDrugKytListupoutRequest) SetStatus(status string) error {
    r.status = status
    r.Set("status", status)
    return nil
}

// Status Getter
func (r AlibabaAlihealthDrugKytListupoutRequest) GetStatus() string {
    return r.status
}
// BillCode Setter
// 单据号
func (r *AlibabaAlihealthDrugKytListupoutRequest) SetBillCode(billCode string) error {
    r.billCode = billCode
    r.Set("bill_code", billCode)
    return nil
}

// BillCode Getter
func (r AlibabaAlihealthDrugKytListupoutRequest) GetBillCode() string {
    return r.billCode
}
// PageSize Setter
// 页大小
func (r *AlibabaAlihealthDrugKytListupoutRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r AlibabaAlihealthDrugKytListupoutRequest) GetPageSize() int64 {
    return r.pageSize
}
// Page Setter
// 页码
func (r *AlibabaAlihealthDrugKytListupoutRequest) SetPage(page int64) error {
    r.page = page
    r.Set("page", page)
    return nil
}

// Page Getter
func (r AlibabaAlihealthDrugKytListupoutRequest) GetPage() int64 {
    return r.page
}
// FromUserId Setter
// 发货单位
func (r *AlibabaAlihealthDrugKytListupoutRequest) SetFromUserId(fromUserId string) error {
    r.fromUserId = fromUserId
    r.Set("from_user_id", fromUserId)
    return nil
}

// FromUserId Getter
func (r AlibabaAlihealthDrugKytListupoutRequest) GetFromUserId() string {
    return r.fromUserId
}
