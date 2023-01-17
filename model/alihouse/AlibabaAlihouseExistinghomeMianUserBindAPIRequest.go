package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeMianUserBindAPIRequest 主账号入驻 API请求
// alibaba.alihouse.existinghome.mian.user.bind
//
// 主账号入驻
type AlibabaAlihouseExistinghomeMianUserBindAPIRequest struct {
	model.Params
	// dto
	_mainAccountReqDto *MainAccountReqDto
}

// NewAlibabaAlihouseExistinghomeMianUserBindRequest 初始化AlibabaAlihouseExistinghomeMianUserBindAPIRequest对象
func NewAlibabaAlihouseExistinghomeMianUserBindRequest() *AlibabaAlihouseExistinghomeMianUserBindAPIRequest {
	return &AlibabaAlihouseExistinghomeMianUserBindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseExistinghomeMianUserBindAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.mian.user.bind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseExistinghomeMianUserBindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseExistinghomeMianUserBindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMainAccountReqDto is MainAccountReqDto Setter
// dto
func (r *AlibabaAlihouseExistinghomeMianUserBindAPIRequest) SetMainAccountReqDto(_mainAccountReqDto *MainAccountReqDto) error {
	r._mainAccountReqDto = _mainAccountReqDto
	r.Set("main_account_req_dto", _mainAccountReqDto)
	return nil
}

// GetMainAccountReqDto MainAccountReqDto Getter
func (r AlibabaAlihouseExistinghomeMianUserBindAPIRequest) GetMainAccountReqDto() *MainAccountReqDto {
	return r._mainAccountReqDto
}
