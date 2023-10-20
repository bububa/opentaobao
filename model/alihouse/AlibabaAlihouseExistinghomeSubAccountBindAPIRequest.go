package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihouseexistinghomesubaccountbindAPIRequest 子账号入驻 API请求
// alibaba.alihouse.existinghome.sub.account.bind
//
// 子账号入驻
type AlibabaalihouseexistinghomesubaccountbindAPIRequest struct {
	model.Params
	// dto
	_subAccountReqDto *SubAccountReqDto
}

// NewAlibabaalihouseexistinghomesubaccountbindRequest 初始化AlibabaalihouseexistinghomesubaccountbindAPIRequest对象
func NewAlibabaalihouseexistinghomesubaccountbindRequest() *AlibabaalihouseexistinghomesubaccountbindAPIRequest {
	return &AlibabaalihouseexistinghomesubaccountbindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihouseexistinghomesubaccountbindAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.sub.account.bind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihouseexistinghomesubaccountbindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihouseexistinghomesubaccountbindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSubAccountReqDto is SubAccountReqDto Setter
// dto
func (r *AlibabaalihouseexistinghomesubaccountbindAPIRequest) SetSubAccountReqDto(_subAccountReqDto *SubAccountReqDto) error {
	r._subAccountReqDto = _subAccountReqDto
	r.Set("sub_account_req_dto", _subAccountReqDto)
	return nil
}

// GetSubAccountReqDto SubAccountReqDto Getter
func (r AlibabaalihouseexistinghomesubaccountbindAPIRequest) GetSubAccountReqDto() *SubAccountReqDto {
	return r._subAccountReqDto
}
