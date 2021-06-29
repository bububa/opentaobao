package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
医疗机构查询本企业上游企业出库单据信息 API请求
alibaba.alihealth.drugtrace.top.yljg.listupout

查询货主/本企业上游企业出库单据信息
*/
type AlibabaAlihealthDrugtraceTopYljgListupoutRequest struct {
    model.Params
    // 企业ID
    _refEntId   string
    // 开始日期（不写时分秒）
    _beginDate   string
    // 结束日期（不写时分秒）
    _endDate   string
    // 生产批号
    _produceBatchNo   string
    // 药品ID
    _drugEntBaseInfoId   string
    // 单据类型
    _billType   string
    // 药品类型
    _physicType   string
    // 状态
    _status   string
    // 单据号
    _billCode   string
    // 页大小
    _pageSize   int64
    // 页码
    _page   int64
    // 发货单位
    _fromUserId   string
}

// 初始化AlibabaAlihealthDrugtraceTopYljgListupoutRequest对象
func NewAlibabaAlihealthDrugtraceTopYljgListupoutRequest() *AlibabaAlihealthDrugtraceTopYljgListupoutRequest{
    return &AlibabaAlihealthDrugtraceTopYljgListupoutRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugtraceTopYljgListupoutRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drugtrace.top.yljg.listupout"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugtraceTopYljgListupoutRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefEntId Setter
// 企业ID
func (r *AlibabaAlihealthDrugtraceTopYljgListupoutRequest) SetRefEntId(_refEntId string) error {
    r._refEntId = _refEntId
    r.Set("ref_ent_id", _refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugtraceTopYljgListupoutRequest) GetRefEntId() string {
    return r._refEntId
}
// BeginDate Setter
// 开始日期（不写时分秒）
func (r *AlibabaAlihealthDrugtraceTopYljgListupoutRequest) SetBeginDate(_beginDate string) error {
    r._beginDate = _beginDate
    r.Set("begin_date", _beginDate)
    return nil
}

// BeginDate Getter
func (r AlibabaAlihealthDrugtraceTopYljgListupoutRequest) GetBeginDate() string {
    return r._beginDate
}
// EndDate Setter
// 结束日期（不写时分秒）
func (r *AlibabaAlihealthDrugtraceTopYljgListupoutRequest) SetEndDate(_endDate string) error {
    r._endDate = _endDate
    r.Set("end_date", _endDate)
    return nil
}

// EndDate Getter
func (r AlibabaAlihealthDrugtraceTopYljgListupoutRequest) GetEndDate() string {
    return r._endDate
}
// ProduceBatchNo Setter
// 生产批号
func (r *AlibabaAlihealthDrugtraceTopYljgListupoutRequest) SetProduceBatchNo(_produceBatchNo string) error {
    r._produceBatchNo = _produceBatchNo
    r.Set("produce_batch_no", _produceBatchNo)
    return nil
}

// ProduceBatchNo Getter
func (r AlibabaAlihealthDrugtraceTopYljgListupoutRequest) GetProduceBatchNo() string {
    return r._produceBatchNo
}
// DrugEntBaseInfoId Setter
// 药品ID
func (r *AlibabaAlihealthDrugtraceTopYljgListupoutRequest) SetDrugEntBaseInfoId(_drugEntBaseInfoId string) error {
    r._drugEntBaseInfoId = _drugEntBaseInfoId
    r.Set("drug_ent_base_info_id", _drugEntBaseInfoId)
    return nil
}

// DrugEntBaseInfoId Getter
func (r AlibabaAlihealthDrugtraceTopYljgListupoutRequest) GetDrugEntBaseInfoId() string {
    return r._drugEntBaseInfoId
}
// BillType Setter
// 单据类型
func (r *AlibabaAlihealthDrugtraceTopYljgListupoutRequest) SetBillType(_billType string) error {
    r._billType = _billType
    r.Set("bill_type", _billType)
    return nil
}

// BillType Getter
func (r AlibabaAlihealthDrugtraceTopYljgListupoutRequest) GetBillType() string {
    return r._billType
}
// PhysicType Setter
// 药品类型
func (r *AlibabaAlihealthDrugtraceTopYljgListupoutRequest) SetPhysicType(_physicType string) error {
    r._physicType = _physicType
    r.Set("physic_type", _physicType)
    return nil
}

// PhysicType Getter
func (r AlibabaAlihealthDrugtraceTopYljgListupoutRequest) GetPhysicType() string {
    return r._physicType
}
// Status Setter
// 状态
func (r *AlibabaAlihealthDrugtraceTopYljgListupoutRequest) SetStatus(_status string) error {
    r._status = _status
    r.Set("status", _status)
    return nil
}

// Status Getter
func (r AlibabaAlihealthDrugtraceTopYljgListupoutRequest) GetStatus() string {
    return r._status
}
// BillCode Setter
// 单据号
func (r *AlibabaAlihealthDrugtraceTopYljgListupoutRequest) SetBillCode(_billCode string) error {
    r._billCode = _billCode
    r.Set("bill_code", _billCode)
    return nil
}

// BillCode Getter
func (r AlibabaAlihealthDrugtraceTopYljgListupoutRequest) GetBillCode() string {
    return r._billCode
}
// PageSize Setter
// 页大小
func (r *AlibabaAlihealthDrugtraceTopYljgListupoutRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r AlibabaAlihealthDrugtraceTopYljgListupoutRequest) GetPageSize() int64 {
    return r._pageSize
}
// Page Setter
// 页码
func (r *AlibabaAlihealthDrugtraceTopYljgListupoutRequest) SetPage(_page int64) error {
    r._page = _page
    r.Set("page", _page)
    return nil
}

// Page Getter
func (r AlibabaAlihealthDrugtraceTopYljgListupoutRequest) GetPage() int64 {
    return r._page
}
// FromUserId Setter
// 发货单位
func (r *AlibabaAlihealthDrugtraceTopYljgListupoutRequest) SetFromUserId(_fromUserId string) error {
    r._fromUserId = _fromUserId
    r.Set("from_user_id", _fromUserId)
    return nil
}

// FromUserId Getter
func (r AlibabaAlihealthDrugtraceTopYljgListupoutRequest) GetFromUserId() string {
    return r._fromUserId
}
