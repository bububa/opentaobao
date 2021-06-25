package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家代客售后逆向申请渲染获取 APIRequest
alibaba.tcls.aelophy.refund.csapplyrender

提供商家代客售后逆向申请渲染获取的接口
*/
type AlibabaTclsAelophyRefundCsapplyrenderRequest struct {
    model.Params

    // 系统自动生成
    refundCsApplyRenderDTO   *RefundCsApplyRenderDto 

}

func NewAlibabaTclsAelophyRefundCsapplyrenderRequest() *AlibabaTclsAelophyRefundCsapplyrenderRequest{
    return &AlibabaTclsAelophyRefundCsapplyrenderRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaTclsAelophyRefundCsapplyrenderRequest) GetApiMethodName() string {
    return "alibaba.tcls.aelophy.refund.csapplyrender"
}

func (r AlibabaTclsAelophyRefundCsapplyrenderRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaTclsAelophyRefundCsapplyrenderRequest) SetRefundCsApplyRenderDTO(refundCsApplyRenderDTO *RefundCsApplyRenderDto) error {
    r.refundCsApplyRenderDTO = refundCsApplyRenderDTO
    r.Set("refund_cs_apply_render_d_t_o", refundCsApplyRenderDTO)
    return nil
}

func (r AlibabaTclsAelophyRefundCsapplyrenderRequest) GetRefundCsApplyRenderDTO() *RefundCsApplyRenderDto {
    return r.refundCsApplyRenderDTO
}

