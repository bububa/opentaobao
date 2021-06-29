package middleclaims

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
国际化中台服务域接收保险公司理赔受理结果 API请求
alibaba.middle.claimsaccept.receive

国际化中台服务域与保险公司交互对接一个订单在保险公司方对该订单进行理赔受理结果的处理后，将该结果返回至服务域
*/
type AlibabaMiddleClaimsacceptReceiveRequest struct {
    model.Params
    // 理赔受理结果实体类
    claimsAcceptDto   *ClaimsAcceptDto
}

// 初始化AlibabaMiddleClaimsacceptReceiveRequest对象
func NewAlibabaMiddleClaimsacceptReceiveRequest() *AlibabaMiddleClaimsacceptReceiveRequest{
    return &AlibabaMiddleClaimsacceptReceiveRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMiddleClaimsacceptReceiveRequest) GetApiMethodName() string {
    return "alibaba.middle.claimsaccept.receive"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMiddleClaimsacceptReceiveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ClaimsAcceptDto Setter
// 理赔受理结果实体类
func (r *AlibabaMiddleClaimsacceptReceiveRequest) SetClaimsAcceptDto(claimsAcceptDto *ClaimsAcceptDto) error {
    r.claimsAcceptDto = claimsAcceptDto
    r.Set("claims_accept_dto", claimsAcceptDto)
    return nil
}

// ClaimsAcceptDto Getter
func (r AlibabaMiddleClaimsacceptReceiveRequest) GetClaimsAcceptDto() *ClaimsAcceptDto {
    return r.claimsAcceptDto
}
