package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据单据号查询单据的详情信息【注意：查询的是本企业的单据】 API请求
alibaba.alihealth.drugtrace.top.yljg.query.upbilldetail

根据单据号查询单据的详情信息【注意：这个接口查询的是本企业的单据，如果是查询上游的单据明细信息，使用xxxxxxx.listupout.detail接口】。
*/
type AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailRequest struct {
    model.Params
    // 单据号码
    _billCode   string
    // 本企业refEntId
    _refEntId   string
}

// 初始化AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailRequest对象
func NewAlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailRequest() *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailRequest{
    return &AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drugtrace.top.yljg.query.upbilldetail"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BillCode Setter
// 单据号码
func (r *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailRequest) SetBillCode(_billCode string) error {
    r._billCode = _billCode
    r.Set("bill_code", _billCode)
    return nil
}

// BillCode Getter
func (r AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailRequest) GetBillCode() string {
    return r._billCode
}
// RefEntId Setter
// 本企业refEntId
func (r *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailRequest) SetRefEntId(_refEntId string) error {
    r._refEntId = _refEntId
    r.Set("ref_ent_id", _refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailRequest) GetRefEntId() string {
    return r._refEntId
}