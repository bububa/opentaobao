package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihouseexistinghomebrokermigrateAPIRequest 融合店经纪人迁移 API请求
// alibaba.alihouse.existinghome.broker.migrate
//
// 融合店经纪人迁移
type AlibabaalihouseexistinghomebrokermigrateAPIRequest struct {
	model.Params
	// 1
	_brokerMigrateDto *BrokerMigrateDto
}

// NewAlibabaalihouseexistinghomebrokermigrateRequest 初始化AlibabaalihouseexistinghomebrokermigrateAPIRequest对象
func NewAlibabaalihouseexistinghomebrokermigrateRequest() *AlibabaalihouseexistinghomebrokermigrateAPIRequest {
	return &AlibabaalihouseexistinghomebrokermigrateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihouseexistinghomebrokermigrateAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.broker.migrate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihouseexistinghomebrokermigrateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihouseexistinghomebrokermigrateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBrokerMigrateDto is BrokerMigrateDto Setter
// 1
func (r *AlibabaalihouseexistinghomebrokermigrateAPIRequest) SetBrokerMigrateDto(_brokerMigrateDto *BrokerMigrateDto) error {
	r._brokerMigrateDto = _brokerMigrateDto
	r.Set("broker_migrate_dto", _brokerMigrateDto)
	return nil
}

// GetBrokerMigrateDto BrokerMigrateDto Getter
func (r AlibabaalihouseexistinghomebrokermigrateAPIRequest) GetBrokerMigrateDto() *BrokerMigrateDto {
	return r._brokerMigrateDto
}
