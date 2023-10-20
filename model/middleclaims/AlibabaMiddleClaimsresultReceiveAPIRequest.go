package middleclaims

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabamiddleclaimsresultreceiveAPIRequest 国际化中台服务域接收理赔结果 API请求
// alibaba.middle.claimsresult.receive
//
// 国际化中台服务域与保险公司交互对接一个订单在保险公司方对该订单进行理赔结果的处理后，将该结果返回至服务域
type AlibabamiddleclaimsresultreceiveAPIRequest struct {
	model.Params
	// 理赔结果实体
	_claimsResultDTO *ClaimsResultDto
}

// NewAlibabamiddleclaimsresultreceiveRequest 初始化AlibabamiddleclaimsresultreceiveAPIRequest对象
func NewAlibabamiddleclaimsresultreceiveRequest() *AlibabamiddleclaimsresultreceiveAPIRequest {
	return &AlibabamiddleclaimsresultreceiveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabamiddleclaimsresultreceiveAPIRequest) GetApiMethodName() string {
	return "alibaba.middle.claimsresult.receive"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabamiddleclaimsresultreceiveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabamiddleclaimsresultreceiveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetClaimsResultDTO is ClaimsResultDTO Setter
// 理赔结果实体
func (r *AlibabamiddleclaimsresultreceiveAPIRequest) SetClaimsResultDTO(_claimsResultDTO *ClaimsResultDto) error {
	r._claimsResultDTO = _claimsResultDTO
	r.Set("claims_result_d_t_o", _claimsResultDTO)
	return nil
}

// GetClaimsResultDTO ClaimsResultDTO Getter
func (r AlibabamiddleclaimsresultreceiveAPIRequest) GetClaimsResultDTO() *ClaimsResultDto {
	return r._claimsResultDTO
}
