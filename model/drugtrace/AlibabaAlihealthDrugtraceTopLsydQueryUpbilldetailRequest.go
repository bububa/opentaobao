package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据单据号查询单据的详情信息【注意：查询的是本企业的单据】 API请求
alibaba.alihealth.drugtrace.top.lsyd.query.upbilldetail

根据单据号查询单据的详情信息【注意：这个接口查询的是本企业的单据，如果是查询上游的单据明细信息，使用xxxxxxx.listupout.detail接口】。
*/
type AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailRequest struct {
    model.Params
    // 单据号码
    billCode   string
    // 本企业refEntId
    refEntId   string
}

// 初始化AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailRequest对象
func NewAlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailRequest() *AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailRequest{
    return &AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drugtrace.top.lsyd.query.upbilldetail"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BillCode Setter
// 单据号码
func (r *AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailRequest) SetBillCode(billCode string) error {
    r.billCode = billCode
    r.Set("bill_code", billCode)
    return nil
}

// BillCode Getter
func (r AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailRequest) GetBillCode() string {
    return r.billCode
}
// RefEntId Setter
// 本企业refEntId
func (r *AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailRequest) GetRefEntId() string {
    return r.refEntId
}
