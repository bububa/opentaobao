package tmallsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallWorkcardIdentifyAPIRequest 工单核销 API请求
// tmall.workcard.identify
//
// 工单核销，当工单完成以后，通过调用此接口核销
// 可以按照多维度核销工单，
// 电器预约安装按照工单维度核销，必选字段workcard_id,buyer_id,identify_code，可选字段attrs，通过扩展字段attrs回传机器码，格式{sn:'机器码'}
type TmallWorkcardIdentifyAPIRequest struct {
	model.Params
	// 核销dto
	_verifyRequestDTO *VerifyRequestDto
}

// NewTmallWorkcardIdentifyRequest 初始化TmallWorkcardIdentifyAPIRequest对象
func NewTmallWorkcardIdentifyRequest() *TmallWorkcardIdentifyAPIRequest {
	return &TmallWorkcardIdentifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallWorkcardIdentifyAPIRequest) GetApiMethodName() string {
	return "tmall.workcard.identify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallWorkcardIdentifyAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is VerifyRequestDTO Setter
// 核销dto
func (r *TmallWorkcardIdentifyAPIRequest) SetVerifyRequestDTO(_verifyRequestDTO *VerifyRequestDto) error {
	r._verifyRequestDTO = _verifyRequestDTO
	r.Set("verify_request_d_t_o", _verifyRequestDTO)
	return nil
}

// Get VerifyRequestDTO Getter
func (r TmallWorkcardIdentifyAPIRequest) GetVerifyRequestDTO() *VerifyRequestDto {
	return r._verifyRequestDTO
}
