package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家代客售后逆向申请渲染获取 API请求
alibaba.tcls.aelophy.refund.csapplyrender

提供商家代客售后逆向申请渲染获取的接口
*/
type AlibabaTclsAelophyRefundCsapplyrenderRequest struct {
    model.Params
    // 系统自动生成
    _refundCsApplyRenderDTO   *RefundCsApplyRenderDTO
}

// 初始化AlibabaTclsAelophyRefundCsapplyrenderRequest对象
func NewAlibabaTclsAelophyRefundCsapplyrenderRequest() *AlibabaTclsAelophyRefundCsapplyrenderRequest{
    return &AlibabaTclsAelophyRefundCsapplyrenderRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaTclsAelophyRefundCsapplyrenderRequest) GetApiMethodName() string {
    return "alibaba.tcls.aelophy.refund.csapplyrender"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaTclsAelophyRefundCsapplyrenderRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefundCsApplyRenderDTO Setter
// 系统自动生成
func (r *AlibabaTclsAelophyRefundCsapplyrenderRequest) SetRefundCsApplyRenderDTO(_refundCsApplyRenderDTO *RefundCsApplyRenderDTO) error {
    r._refundCsApplyRenderDTO = _refundCsApplyRenderDTO
    r.Set("refund_cs_apply_render_d_t_o", _refundCsApplyRenderDTO)
    return nil
}

// RefundCsApplyRenderDTO Getter
func (r AlibabaTclsAelophyRefundCsapplyrenderRequest) GetRefundCsApplyRenderDTO() *RefundCsApplyRenderDTO {
    return r._refundCsApplyRenderDTO
}
