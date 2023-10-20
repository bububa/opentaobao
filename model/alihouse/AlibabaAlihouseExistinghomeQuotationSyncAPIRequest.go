package alihouse

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeQuotationSyncAPIRequest 二手房行情数据同步 API请求
// alibaba.alihouse.existinghome.quotation.sync
//
// 二手房行情数据同步
type AlibabaAlihouseExistinghomeQuotationSyncAPIRequest struct {
	model.Params
	// dto
	_existingHouseQuotationReqDto *ExistingHouseQuotationReqDto
}

// NewAlibabaAlihouseExistinghomeQuotationSyncRequest 初始化AlibabaAlihouseExistinghomeQuotationSyncAPIRequest对象
func NewAlibabaAlihouseExistinghomeQuotationSyncRequest() *AlibabaAlihouseExistinghomeQuotationSyncAPIRequest {
	return &AlibabaAlihouseExistinghomeQuotationSyncAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseExistinghomeQuotationSyncAPIRequest) Reset() {
	r._existingHouseQuotationReqDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseExistinghomeQuotationSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.quotation.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseExistinghomeQuotationSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseExistinghomeQuotationSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetExistingHouseQuotationReqDto is ExistingHouseQuotationReqDto Setter
// dto
func (r *AlibabaAlihouseExistinghomeQuotationSyncAPIRequest) SetExistingHouseQuotationReqDto(_existingHouseQuotationReqDto *ExistingHouseQuotationReqDto) error {
	r._existingHouseQuotationReqDto = _existingHouseQuotationReqDto
	r.Set("existing_house_quotation_req_dto", _existingHouseQuotationReqDto)
	return nil
}

// GetExistingHouseQuotationReqDto ExistingHouseQuotationReqDto Getter
func (r AlibabaAlihouseExistinghomeQuotationSyncAPIRequest) GetExistingHouseQuotationReqDto() *ExistingHouseQuotationReqDto {
	return r._existingHouseQuotationReqDto
}

var poolAlibabaAlihouseExistinghomeQuotationSyncAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseExistinghomeQuotationSyncRequest()
	},
}

// GetAlibabaAlihouseExistinghomeQuotationSyncRequest 从 sync.Pool 获取 AlibabaAlihouseExistinghomeQuotationSyncAPIRequest
func GetAlibabaAlihouseExistinghomeQuotationSyncAPIRequest() *AlibabaAlihouseExistinghomeQuotationSyncAPIRequest {
	return poolAlibabaAlihouseExistinghomeQuotationSyncAPIRequest.Get().(*AlibabaAlihouseExistinghomeQuotationSyncAPIRequest)
}

// ReleaseAlibabaAlihouseExistinghomeQuotationSyncAPIRequest 将 AlibabaAlihouseExistinghomeQuotationSyncAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseExistinghomeQuotationSyncAPIRequest(v *AlibabaAlihouseExistinghomeQuotationSyncAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseExistinghomeQuotationSyncAPIRequest.Put(v)
}
