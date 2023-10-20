package alihouse

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeAgreementSyncAPIRequest 二手房电子协议数据同步 API请求
// alibaba.alihouse.existinghome.agreement.sync
//
// 二手房电子协议数据同步
type AlibabaAlihouseExistinghomeAgreementSyncAPIRequest struct {
	model.Params
	// 数据结构
	_existingHomeElectricAgreementDto *ExistingHomeElectricAgreementDto
}

// NewAlibabaAlihouseExistinghomeAgreementSyncRequest 初始化AlibabaAlihouseExistinghomeAgreementSyncAPIRequest对象
func NewAlibabaAlihouseExistinghomeAgreementSyncRequest() *AlibabaAlihouseExistinghomeAgreementSyncAPIRequest {
	return &AlibabaAlihouseExistinghomeAgreementSyncAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseExistinghomeAgreementSyncAPIRequest) Reset() {
	r._existingHomeElectricAgreementDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseExistinghomeAgreementSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.agreement.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseExistinghomeAgreementSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseExistinghomeAgreementSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetExistingHomeElectricAgreementDto is ExistingHomeElectricAgreementDto Setter
// 数据结构
func (r *AlibabaAlihouseExistinghomeAgreementSyncAPIRequest) SetExistingHomeElectricAgreementDto(_existingHomeElectricAgreementDto *ExistingHomeElectricAgreementDto) error {
	r._existingHomeElectricAgreementDto = _existingHomeElectricAgreementDto
	r.Set("existing_home_electric_agreement_dto", _existingHomeElectricAgreementDto)
	return nil
}

// GetExistingHomeElectricAgreementDto ExistingHomeElectricAgreementDto Getter
func (r AlibabaAlihouseExistinghomeAgreementSyncAPIRequest) GetExistingHomeElectricAgreementDto() *ExistingHomeElectricAgreementDto {
	return r._existingHomeElectricAgreementDto
}

var poolAlibabaAlihouseExistinghomeAgreementSyncAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseExistinghomeAgreementSyncRequest()
	},
}

// GetAlibabaAlihouseExistinghomeAgreementSyncRequest 从 sync.Pool 获取 AlibabaAlihouseExistinghomeAgreementSyncAPIRequest
func GetAlibabaAlihouseExistinghomeAgreementSyncAPIRequest() *AlibabaAlihouseExistinghomeAgreementSyncAPIRequest {
	return poolAlibabaAlihouseExistinghomeAgreementSyncAPIRequest.Get().(*AlibabaAlihouseExistinghomeAgreementSyncAPIRequest)
}

// ReleaseAlibabaAlihouseExistinghomeAgreementSyncAPIRequest 将 AlibabaAlihouseExistinghomeAgreementSyncAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseExistinghomeAgreementSyncAPIRequest(v *AlibabaAlihouseExistinghomeAgreementSyncAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseExistinghomeAgreementSyncAPIRequest.Put(v)
}
