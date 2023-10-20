package tmallsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallworkcardidentifyAPIRequest 工单核销 API请求
// tmall.workcard.identify
//
// 工单核销，当工单完成以后，通过调用此接口核销
// 可以按照多维度核销工单，
// 电器预约安装按照工单维度核销，必选字段workcard_id,buyer_id,identify_code，可选字段attrs，通过扩展字段attrs回传机器码，格式{sn:&#39;机器码&#39;}
type TmallworkcardidentifyAPIRequest struct {
	model.Params
	// 核销dto
	_verifyRequestDTO *VerifyRequestDto
}

// NewTmallworkcardidentifyRequest 初始化TmallworkcardidentifyAPIRequest对象
func NewTmallworkcardidentifyRequest() *TmallworkcardidentifyAPIRequest {
	return &TmallworkcardidentifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallworkcardidentifyAPIRequest) GetApiMethodName() string {
	return "tmall.workcard.identify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallworkcardidentifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallworkcardidentifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetVerifyRequestDTO is VerifyRequestDTO Setter
// 核销dto
func (r *TmallworkcardidentifyAPIRequest) SetVerifyRequestDTO(_verifyRequestDTO *VerifyRequestDto) error {
	r._verifyRequestDTO = _verifyRequestDTO
	r.Set("verify_request_d_t_o", _verifyRequestDTO)
	return nil
}

// GetVerifyRequestDTO VerifyRequestDTO Getter
func (r TmallworkcardidentifyAPIRequest) GetVerifyRequestDTO() *VerifyRequestDto {
	return r._verifyRequestDTO
}
