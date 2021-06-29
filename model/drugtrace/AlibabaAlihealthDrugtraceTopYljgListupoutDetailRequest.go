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
type AlibabaAlihealthDrugtraceTopYljgListupoutDetailRequest struct {
    model.Params
    // 企业id
    refEntId   string
    // 单据编码
    billCode   string
    // 发货企业renEntId
    fromRefUserId   string
    // 收货企业refEntId
    toRefUserId   string
}

// 初始化AlibabaAlihealthDrugtraceTopYljgListupoutDetailRequest对象
func NewAlibabaAlihealthDrugtraceTopYljgListupoutDetailRequest() *AlibabaAlihealthDrugtraceTopYljgListupoutDetailRequest{
    return &AlibabaAlihealthDrugtraceTopYljgListupoutDetailRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugtraceTopYljgListupoutDetailRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drugtrace.top.yljg.listupout.detail"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugtraceTopYljgListupoutDetailRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefEntId Setter
// 企业id
func (r *AlibabaAlihealthDrugtraceTopYljgListupoutDetailRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugtraceTopYljgListupoutDetailRequest) GetRefEntId() string {
    return r.refEntId
}
// BillCode Setter
// 单据编码
func (r *AlibabaAlihealthDrugtraceTopYljgListupoutDetailRequest) SetBillCode(billCode string) error {
    r.billCode = billCode
    r.Set("bill_code", billCode)
    return nil
}

// BillCode Getter
func (r AlibabaAlihealthDrugtraceTopYljgListupoutDetailRequest) GetBillCode() string {
    return r.billCode
}
// FromRefUserId Setter
// 发货企业renEntId
func (r *AlibabaAlihealthDrugtraceTopYljgListupoutDetailRequest) SetFromRefUserId(fromRefUserId string) error {
    r.fromRefUserId = fromRefUserId
    r.Set("from_ref_user_id", fromRefUserId)
    return nil
}

// FromRefUserId Getter
func (r AlibabaAlihealthDrugtraceTopYljgListupoutDetailRequest) GetFromRefUserId() string {
    return r.fromRefUserId
}
// ToRefUserId Setter
// 收货企业refEntId
func (r *AlibabaAlihealthDrugtraceTopYljgListupoutDetailRequest) SetToRefUserId(toRefUserId string) error {
    r.toRefUserId = toRefUserId
    r.Set("to_ref_user_id", toRefUserId)
    return nil
}

// ToRefUserId Getter
func (r AlibabaAlihealthDrugtraceTopYljgListupoutDetailRequest) GetToRefUserId() string {
    return r.toRefUserId
}
