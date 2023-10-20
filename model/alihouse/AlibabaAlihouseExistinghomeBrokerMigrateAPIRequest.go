package alihouse

import (
	"net/url"

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
		Params: model.NewParams(),
	}
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
