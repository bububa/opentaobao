package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
上游出库单单据明细查询 API请求
alibaba.alihealth.drugtrace.top.yljg.listupout.detail

查询上游出库单明细(带追溯码信息)
*/
type AlibabaAlihealthDrugtraceTopYljgListupoutDetailAPIRequest struct {
    model.Params
    // 企业id
    _refEntId   string
    // 单据编码
    _billCode   string
    // 发货企业renEntId
    _fromRefUserId   string
    // 收货企业refEntId
    _toRefUserId   string
}

// 初始化AlibabaAlihealthDrugtraceTopYljgListupoutDetailAPIRequest对象
func NewAlibabaAlihealthDrugtraceTopYljgListupoutDetailRequest() *AlibabaAlihealthDrugtraceTopYljgListupoutDetailAPIRequest{
    return &AlibabaAlihealthDrugtraceTopYljgListupoutDetailAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugtraceTopYljgListupoutDetailAPIRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drugtrace.top.yljg.listupout.detail"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugtraceTopYljgListupoutDetailAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefEntId Setter
// 企业id
func (r *AlibabaAlihealthDrugtraceTopYljgListupoutDetailAPIRequest) SetRefEntId(_refEntId string) error {
    r._refEntId = _refEntId
    r.Set("ref_ent_id", _refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugtraceTopYljgListupoutDetailAPIRequest) GetRefEntId() string {
    return r._refEntId
}
// BillCode Setter
// 单据编码
func (r *AlibabaAlihealthDrugtraceTopYljgListupoutDetailAPIRequest) SetBillCode(_billCode string) error {
    r._billCode = _billCode
    r.Set("bill_code", _billCode)
    return nil
}

// BillCode Getter
func (r AlibabaAlihealthDrugtraceTopYljgListupoutDetailAPIRequest) GetBillCode() string {
    return r._billCode
}
// FromRefUserId Setter
// 发货企业renEntId
func (r *AlibabaAlihealthDrugtraceTopYljgListupoutDetailAPIRequest) SetFromRefUserId(_fromRefUserId string) error {
    r._fromRefUserId = _fromRefUserId
    r.Set("from_ref_user_id", _fromRefUserId)
    return nil
}

// FromRefUserId Getter
func (r AlibabaAlihealthDrugtraceTopYljgListupoutDetailAPIRequest) GetFromRefUserId() string {
    return r._fromRefUserId
}
// ToRefUserId Setter
// 收货企业refEntId
func (r *AlibabaAlihealthDrugtraceTopYljgListupoutDetailAPIRequest) SetToRefUserId(_toRefUserId string) error {
    r._toRefUserId = _toRefUserId
    r.Set("to_ref_user_id", _toRefUserId)
    return nil
}

// ToRefUserId Getter
func (r AlibabaAlihealthDrugtraceTopYljgListupoutDetailAPIRequest) GetToRefUserId() string {
    return r._toRefUserId
}
