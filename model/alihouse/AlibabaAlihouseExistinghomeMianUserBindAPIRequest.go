package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihouseexistinghomemianuserbindAPIRequest 主账号入驻 API请求
// alibaba.alihouse.existinghome.mian.user.bind
//
// 主账号入驻
type AlibabaalihouseexistinghomemianuserbindAPIRequest struct {
	model.Params
	// dto
	_mainAccountReqDto *MainAccountReqDto
}

// NewAlibabaalihouseexistinghomemianuserbindRequest 初始化AlibabaalihouseexistinghomemianuserbindAPIRequest对象
func NewAlibabaalihouseexistinghomemianuserbindRequest() *AlibabaalihouseexistinghomemianuserbindAPIRequest {
	return &AlibabaalihouseexistinghomemianuserbindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihouseexistinghomemianuserbindAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.mian.user.bind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihouseexistinghomemianuserbindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihouseexistinghomemianuserbindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMainAccountReqDto is MainAccountReqDto Setter
// dto
func (r *AlibabaalihouseexistinghomemianuserbindAPIRequest) SetMainAccountReqDto(_mainAccountReqDto *MainAccountReqDto) error {
	r._mainAccountReqDto = _mainAccountReqDto
	r.Set("main_account_req_dto", _mainAccountReqDto)
	return nil
}

// GetMainAccountReqDto MainAccountReqDto Getter
func (r AlibabaalihouseexistinghomemianuserbindAPIRequest) GetMainAccountReqDto() *MainAccountReqDto {
	return r._mainAccountReqDto
}
