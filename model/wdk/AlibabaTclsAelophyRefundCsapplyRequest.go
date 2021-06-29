package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家代客售后提交逆向申请 API请求
alibaba.tcls.aelophy.refund.csapply

商家代客售后提交逆向申请
*/
type AlibabaTclsAelophyRefundCsapplyRequest struct {
    model.Params
    // 逆向申请入参
    _refundCsApplyDTO   *RefundCsApplyDTO
}

// 初始化AlibabaTclsAelophyRefundCsapplyRequest对象
func NewAlibabaTclsAelophyRefundCsapplyRequest() *AlibabaTclsAelophyRefundCsapplyRequest{
    return &AlibabaTclsAelophyRefundCsapplyRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaTclsAelophyRefundCsapplyRequest) GetApiMethodName() string {
    return "alibaba.tcls.aelophy.refund.csapply"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaTclsAelophyRefundCsapplyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefundCsApplyDTO Setter
// 逆向申请入参
func (r *AlibabaTclsAelophyRefundCsapplyRequest) SetRefundCsApplyDTO(_refundCsApplyDTO *RefundCsApplyDTO) error {
    r._refundCsApplyDTO = _refundCsApplyDTO
    r.Set("refund_cs_apply_d_t_o", _refundCsApplyDTO)
    return nil
}

// RefundCsApplyDTO Getter
func (r AlibabaTclsAelophyRefundCsapplyRequest) GetRefundCsApplyDTO() *RefundCsApplyDTO {
    return r._refundCsApplyDTO
}
