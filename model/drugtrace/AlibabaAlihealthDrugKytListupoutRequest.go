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
func (r *AlibabaAlihealthDrugKytListupoutRequest) SetRefEntId(_refEntId string) error {
    r._refEntId = _refEntId
    r.Set("ref_ent_id", _refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugKytListupoutRequest) GetRefEntId() string {
    return r._refEntId
}
// BeginDate Setter
// 开始日期（不写时分秒）
func (r *AlibabaAlihealthDrugKytListupoutRequest) SetBeginDate(_beginDate string) error {
    r._beginDate = _beginDate
    r.Set("begin_date", _beginDate)
    return nil
}

// BeginDate Getter
func (r AlibabaAlihealthDrugKytListupoutRequest) GetBeginDate() string {
    return r._beginDate
}
// EndDate Setter
// 结束日期（不写时分秒）
func (r *AlibabaAlihealthDrugKytListupoutRequest) SetEndDate(_endDate string) error {
    r._endDate = _endDate
    r.Set("end_date", _endDate)
    return nil
}

// EndDate Getter
func (r AlibabaAlihealthDrugKytListupoutRequest) GetEndDate() string {
    return r._endDate
}
// ProduceBatchNo Setter
// 生产批号
func (r *AlibabaAlihealthDrugKytListupoutRequest) SetProduceBatchNo(_produceBatchNo string) error {
    r._produceBatchNo = _produceBatchNo
    r.Set("produce_batch_no", _produceBatchNo)
    return nil
}

// ProduceBatchNo Getter
func (r AlibabaAlihealthDrugKytListupoutRequest) GetProduceBatchNo() string {
    return r._produceBatchNo
}
// DrugEntBaseInfoId Setter
// 药品ID
func (r *AlibabaAlihealthDrugKytListupoutRequest) SetDrugEntBaseInfoId(_drugEntBaseInfoId string) error {
    r._drugEntBaseInfoId = _drugEntBaseInfoId
    r.Set("drug_ent_base_info_id", _drugEntBaseInfoId)
    return nil
}

// DrugEntBaseInfoId Getter
func (r AlibabaAlihealthDrugKytListupoutRequest) GetDrugEntBaseInfoId() string {
    return r._drugEntBaseInfoId
}
// BillType Setter
// 单据类型
func (r *AlibabaAlihealthDrugKytListupoutRequest) SetBillType(_billType string) error {
    r._billType = _billType
    r.Set("bill_type", _billType)
    return nil
}

// BillType Getter
func (r AlibabaAlihealthDrugKytListupoutRequest) GetBillType() string {
    return r._billType
}
// PhysicType Setter
// 药品类型
func (r *AlibabaAlihealthDrugKytListupoutRequest) SetPhysicType(_physicType string) error {
    r._physicType = _physicType
    r.Set("physic_type", _physicType)
    return nil
}

// PhysicType Getter
func (r AlibabaAlihealthDrugKytListupoutRequest) GetPhysicType() string {
    return r._physicType
}
// Status Setter
// 状态
func (r *AlibabaAlihealthDrugKytListupoutRequest) SetStatus(_status string) error {
    r._status = _status
    r.Set("status", _status)
    return nil
}

// Status Getter
func (r AlibabaAlihealthDrugKytListupoutRequest) GetStatus() string {
    return r._status
}
// BillCode Setter
// 单据号
func (r *AlibabaAlihealthDrugKytListupoutRequest) SetBillCode(_billCode string) error {
    r._billCode = _billCode
    r.Set("bill_code", _billCode)
    return nil
}

// BillCode Getter
func (r AlibabaAlihealthDrugKytListupoutRequest) GetBillCode() string {
    return r._billCode
}
// PageSize Setter
// 页大小
func (r *AlibabaAlihealthDrugKytListupoutRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r AlibabaAlihealthDrugKytListupoutRequest) GetPageSize() int64 {
    return r._pageSize
}
// Page Setter
// 页码
func (r *AlibabaAlihealthDrugKytListupoutRequest) SetPage(_page int64) error {
    r._page = _page
    r.Set("page", _page)
    return nil
}

// Page Getter
func (r AlibabaAlihealthDrugKytListupoutRequest) GetPage() int64 {
    return r._page
}
// FromUserId Setter
// 发货单位
func (r *AlibabaAlihealthDrugKytListupoutRequest) SetFromUserId(_fromUserId string) error {
    r._fromUserId = _fromUserId
    r.Set("from_user_id", _fromUserId)
    return nil
}

// FromUserId Getter
func (r AlibabaAlihealthDrugKytListupoutRequest) GetFromUserId() string {
    return r._fromUserId
}
