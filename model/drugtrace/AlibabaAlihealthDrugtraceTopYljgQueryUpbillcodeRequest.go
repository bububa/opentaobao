package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
通过一个码，查询这个码对应的上游企业出库单的单据号 API请求
alibaba.alihealth.drugtrace.top.yljg.query.upbillcode

一个查询上游出库单号的接口。企业在扫码入库时，接口通过扫到的码判定这个码对应的上游企业所属的出库单据号
*/
type AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeRequest struct {
    model.Params
    // 追溯码
    _code   string
    // 企业ID （一般为要查询单据的收货企业）
    _refEntId   string
}

// 初始化AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeRequest对象
func NewAlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeRequest() *AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeRequest{
    return &AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drugtrace.top.yljg.query.upbillcode"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Code Setter
// 追溯码
func (r *AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeRequest) SetCode(_code string) error {
    r._code = _code
    r.Set("code", _code)
    return nil
}

// Code Getter
func (r AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeRequest) GetCode() string {
    return r._code
}
// RefEntId Setter
// 企业ID （一般为要查询单据的收货企业）
func (r *AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeRequest) SetRefEntId(_refEntId string) error {
    r._refEntId = _refEntId
    r.Set("ref_ent_id", _refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeRequest) GetRefEntId() string {
    return r._refEntId
}
