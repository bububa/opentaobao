package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeBankcardSyncAPIRequest 结算账号同步 API请求
// alibaba.alihouse.existinghome.bankcard.sync
//
// 结算账号同步
type AlibabaAlihouseExistinghomeBankcardSyncAPIRequest struct {
	model.Params
	// 结构
	_existingHomeBankCardDto *ExistingHomeBankCardDto
}

// NewAlibabaAlihouseExistinghomeBankcardSyncRequest 初始化AlibabaAlihouseExistinghomeBankcardSyncAPIRequest对象
func NewAlibabaAlihouseExistinghomeBankcardSyncRequest() *AlibabaAlihouseExistinghomeBankcardSyncAPIRequest {
	return &AlibabaAlihouseExistinghomeBankcardSyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseExistinghomeBankcardSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.bankcard.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseExistinghomeBankcardSyncAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetExistingHomeBankCardDto is ExistingHomeBankCardDto Setter
// 结构
func (r *AlibabaAlihouseExistinghomeBankcardSyncAPIRequest) SetExistingHomeBankCardDto(_existingHomeBankCardDto *ExistingHomeBankCardDto) error {
	r._existingHomeBankCardDto = _existingHomeBankCardDto
	r.Set("existing_home_bank_card_dto", _existingHomeBankCardDto)
	return nil
}

// GetExistingHomeBankCardDto ExistingHomeBankCardDto Getter
func (r AlibabaAlihouseExistinghomeBankcardSyncAPIRequest) GetExistingHomeBankCardDto() *ExistingHomeBankCardDto {
	return r._existingHomeBankCardDto
}
