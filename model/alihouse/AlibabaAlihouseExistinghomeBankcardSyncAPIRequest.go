package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihouseexistinghomebankcardsyncAPIRequest 结算账号同步 API请求
// alibaba.alihouse.existinghome.bankcard.sync
//
// 结算账号同步
type AlibabaalihouseexistinghomebankcardsyncAPIRequest struct {
	model.Params
	// 结构
	_existingHomeBankCardDto *ExistingHomeBankCardDto
}

// NewAlibabaalihouseexistinghomebankcardsyncRequest 初始化AlibabaalihouseexistinghomebankcardsyncAPIRequest对象
func NewAlibabaalihouseexistinghomebankcardsyncRequest() *AlibabaalihouseexistinghomebankcardsyncAPIRequest {
	return &AlibabaalihouseexistinghomebankcardsyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihouseexistinghomebankcardsyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.bankcard.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihouseexistinghomebankcardsyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihouseexistinghomebankcardsyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetExistingHomeBankCardDto is ExistingHomeBankCardDto Setter
// 结构
func (r *AlibabaalihouseexistinghomebankcardsyncAPIRequest) SetExistingHomeBankCardDto(_existingHomeBankCardDto *ExistingHomeBankCardDto) error {
	r._existingHomeBankCardDto = _existingHomeBankCardDto
	r.Set("existing_home_bank_card_dto", _existingHomeBankCardDto)
	return nil
}

// GetExistingHomeBankCardDto ExistingHomeBankCardDto Getter
func (r AlibabaalihouseexistinghomebankcardsyncAPIRequest) GetExistingHomeBankCardDto() *ExistingHomeBankCardDto {
	return r._existingHomeBankCardDto
}
