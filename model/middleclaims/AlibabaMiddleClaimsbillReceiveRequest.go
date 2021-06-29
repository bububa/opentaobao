package middleclaims

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
国际化中台服务域接收理赔账单 API请求
alibaba.middle.claimsbill.receive

国际化中台服务域与保险公司交互对接一个订单在保险公司方对该订单进行理赔打款的处理后，将该打款结果返回至服务域
*/
type AlibabaMiddleClaimsbillReceiveRequest struct {
    model.Params
    // 理赔账单实体
    _claimsBillDto   *ClaimsBillDto
}

// 初始化AlibabaMiddleClaimsbillReceiveRequest对象
func NewAlibabaMiddleClaimsbillReceiveRequest() *AlibabaMiddleClaimsbillReceiveRequest{
    return &AlibabaMiddleClaimsbillReceiveRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMiddleClaimsbillReceiveRequest) GetApiMethodName() string {
    return "alibaba.middle.claimsbill.receive"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMiddleClaimsbillReceiveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ClaimsBillDto Setter
// 理赔账单实体
func (r *AlibabaMiddleClaimsbillReceiveRequest) SetClaimsBillDto(_claimsBillDto *ClaimsBillDto) error {
    r._claimsBillDto = _claimsBillDto
    r.Set("claims_bill_dto", _claimsBillDto)
    return nil
}

// ClaimsBillDto Getter
func (r AlibabaMiddleClaimsbillReceiveRequest) GetClaimsBillDto() *ClaimsBillDto {
    return r._claimsBillDto
}
