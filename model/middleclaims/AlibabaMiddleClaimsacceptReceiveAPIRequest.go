package middleclaims

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabamiddleclaimsacceptreceiveAPIRequest 国际化中台服务域接收保险公司理赔受理结果 API请求
// alibaba.middle.claimsaccept.receive
//
// 国际化中台服务域与保险公司交互对接一个订单在保险公司方对该订单进行理赔受理结果的处理后，将该结果返回至服务域
type AlibabamiddleclaimsacceptreceiveAPIRequest struct {
	model.Params
	// 理赔受理结果实体类
	_claimsAcceptDto *ClaimsAcceptDto
}

// NewAlibabamiddleclaimsacceptreceiveRequest 初始化AlibabamiddleclaimsacceptreceiveAPIRequest对象
func NewAlibabamiddleclaimsacceptreceiveRequest() *AlibabamiddleclaimsacceptreceiveAPIRequest {
	return &AlibabamiddleclaimsacceptreceiveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabamiddleclaimsacceptreceiveAPIRequest) GetApiMethodName() string {
	return "alibaba.middle.claimsaccept.receive"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabamiddleclaimsacceptreceiveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabamiddleclaimsacceptreceiveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetClaimsAcceptDto is ClaimsAcceptDto Setter
// 理赔受理结果实体类
func (r *AlibabamiddleclaimsacceptreceiveAPIRequest) SetClaimsAcceptDto(_claimsAcceptDto *ClaimsAcceptDto) error {
	r._claimsAcceptDto = _claimsAcceptDto
	r.Set("claims_accept_dto", _claimsAcceptDto)
	return nil
}

// GetClaimsAcceptDto ClaimsAcceptDto Getter
func (r AlibabamiddleclaimsacceptreceiveAPIRequest) GetClaimsAcceptDto() *ClaimsAcceptDto {
	return r._claimsAcceptDto
}
