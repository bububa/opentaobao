package middleclaims

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMiddleClaimsacceptReceiveAPIRequest 国际化中台服务域接收保险公司理赔受理结果 API请求
// alibaba.middle.claimsaccept.receive
//
// 国际化中台服务域与保险公司交互对接一个订单在保险公司方对该订单进行理赔受理结果的处理后，将该结果返回至服务域
type AlibabaMiddleClaimsacceptReceiveAPIRequest struct {
	model.Params
	// 理赔受理结果实体类
	_claimsAcceptDto *ClaimsAcceptDto
}

// NewAlibabaMiddleClaimsacceptReceiveRequest 初始化AlibabaMiddleClaimsacceptReceiveAPIRequest对象
func NewAlibabaMiddleClaimsacceptReceiveRequest() *AlibabaMiddleClaimsacceptReceiveAPIRequest {
	return &AlibabaMiddleClaimsacceptReceiveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMiddleClaimsacceptReceiveAPIRequest) GetApiMethodName() string {
	return "alibaba.middle.claimsaccept.receive"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMiddleClaimsacceptReceiveAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ClaimsAcceptDto Setter
// 理赔受理结果实体类
func (r *AlibabaMiddleClaimsacceptReceiveAPIRequest) SetClaimsAcceptDto(_claimsAcceptDto *ClaimsAcceptDto) error {
	r._claimsAcceptDto = _claimsAcceptDto
	r.Set("claims_accept_dto", _claimsAcceptDto)
	return nil
}

// Get ClaimsAcceptDto Getter
func (r AlibabaMiddleClaimsacceptReceiveAPIRequest) GetClaimsAcceptDto() *ClaimsAcceptDto {
	return r._claimsAcceptDto
}
