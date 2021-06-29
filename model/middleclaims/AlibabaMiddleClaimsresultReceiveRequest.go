package middleclaims

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
国际化中台服务域接收理赔结果 API请求
alibaba.middle.claimsresult.receive

国际化中台服务域与保险公司交互对接一个订单在保险公司方对该订单进行理赔结果的处理后，将该结果返回至服务域
*/
type AlibabaMiddleClaimsresultReceiveRequest struct {
    model.Params
    // 理赔结果实体
    _claimsResultDTO   *ClaimsResultDto
}

// 初始化AlibabaMiddleClaimsresultReceiveRequest对象
func NewAlibabaMiddleClaimsresultReceiveRequest() *AlibabaMiddleClaimsresultReceiveRequest{
    return &AlibabaMiddleClaimsresultReceiveRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMiddleClaimsresultReceiveRequest) GetApiMethodName() string {
    return "alibaba.middle.claimsresult.receive"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMiddleClaimsresultReceiveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ClaimsResultDTO Setter
// 理赔结果实体
func (r *AlibabaMiddleClaimsresultReceiveRequest) SetClaimsResultDTO(_claimsResultDTO *ClaimsResultDto) error {
    r._claimsResultDTO = _claimsResultDTO
    r.Set("claims_result_d_t_o", _claimsResultDTO)
    return nil
}

// ClaimsResultDTO Getter
func (r AlibabaMiddleClaimsresultReceiveRequest) GetClaimsResultDTO() *ClaimsResultDto {
    return r._claimsResultDTO
}
