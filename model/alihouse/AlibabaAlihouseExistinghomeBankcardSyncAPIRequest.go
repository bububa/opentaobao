package alihouse

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseExistinghomeBankcardSyncAPIRequest) Reset() {
	r._existingHomeBankCardDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseExistinghomeBankcardSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.bankcard.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseExistinghomeBankcardSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseExistinghomeBankcardSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaAlihouseExistinghomeBankcardSyncAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseExistinghomeBankcardSyncRequest()
	},
}

// GetAlibabaAlihouseExistinghomeBankcardSyncRequest 从 sync.Pool 获取 AlibabaAlihouseExistinghomeBankcardSyncAPIRequest
func GetAlibabaAlihouseExistinghomeBankcardSyncAPIRequest() *AlibabaAlihouseExistinghomeBankcardSyncAPIRequest {
	return poolAlibabaAlihouseExistinghomeBankcardSyncAPIRequest.Get().(*AlibabaAlihouseExistinghomeBankcardSyncAPIRequest)
}

// ReleaseAlibabaAlihouseExistinghomeBankcardSyncAPIRequest 将 AlibabaAlihouseExistinghomeBankcardSyncAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseExistinghomeBankcardSyncAPIRequest(v *AlibabaAlihouseExistinghomeBankcardSyncAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseExistinghomeBankcardSyncAPIRequest.Put(v)
}
