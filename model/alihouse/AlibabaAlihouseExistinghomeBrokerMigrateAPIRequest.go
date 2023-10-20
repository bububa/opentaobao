package alihouse

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeBrokerMigrateAPIRequest 融合店经纪人迁移 API请求
// alibaba.alihouse.existinghome.broker.migrate
//
// 融合店经纪人迁移
type AlibabaAlihouseExistinghomeBrokerMigrateAPIRequest struct {
	model.Params
	// 1
	_brokerMigrateDto *BrokerMigrateDto
}

// NewAlibabaAlihouseExistinghomeBrokerMigrateRequest 初始化AlibabaAlihouseExistinghomeBrokerMigrateAPIRequest对象
func NewAlibabaAlihouseExistinghomeBrokerMigrateRequest() *AlibabaAlihouseExistinghomeBrokerMigrateAPIRequest {
	return &AlibabaAlihouseExistinghomeBrokerMigrateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseExistinghomeBrokerMigrateAPIRequest) Reset() {
	r._brokerMigrateDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseExistinghomeBrokerMigrateAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.broker.migrate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseExistinghomeBrokerMigrateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseExistinghomeBrokerMigrateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBrokerMigrateDto is BrokerMigrateDto Setter
// 1
func (r *AlibabaAlihouseExistinghomeBrokerMigrateAPIRequest) SetBrokerMigrateDto(_brokerMigrateDto *BrokerMigrateDto) error {
	r._brokerMigrateDto = _brokerMigrateDto
	r.Set("broker_migrate_dto", _brokerMigrateDto)
	return nil
}

// GetBrokerMigrateDto BrokerMigrateDto Getter
func (r AlibabaAlihouseExistinghomeBrokerMigrateAPIRequest) GetBrokerMigrateDto() *BrokerMigrateDto {
	return r._brokerMigrateDto
}

var poolAlibabaAlihouseExistinghomeBrokerMigrateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseExistinghomeBrokerMigrateRequest()
	},
}

// GetAlibabaAlihouseExistinghomeBrokerMigrateRequest 从 sync.Pool 获取 AlibabaAlihouseExistinghomeBrokerMigrateAPIRequest
func GetAlibabaAlihouseExistinghomeBrokerMigrateAPIRequest() *AlibabaAlihouseExistinghomeBrokerMigrateAPIRequest {
	return poolAlibabaAlihouseExistinghomeBrokerMigrateAPIRequest.Get().(*AlibabaAlihouseExistinghomeBrokerMigrateAPIRequest)
}

// ReleaseAlibabaAlihouseExistinghomeBrokerMigrateAPIRequest 将 AlibabaAlihouseExistinghomeBrokerMigrateAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseExistinghomeBrokerMigrateAPIRequest(v *AlibabaAlihouseExistinghomeBrokerMigrateAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseExistinghomeBrokerMigrateAPIRequest.Put(v)
}
