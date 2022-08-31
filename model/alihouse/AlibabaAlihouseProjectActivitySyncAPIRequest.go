package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseProjectActivitySyncAPIRequest 电商券数据同步 API请求
// alibaba.alihouse.project.activity.sync
//
// 电商券数据同步
type AlibabaAlihouseProjectActivitySyncAPIRequest struct {
	model.Params
	// 数据
	_businessActivityDataDto *BusinessActivityDataDto
}

// NewAlibabaAlihouseProjectActivitySyncRequest 初始化AlibabaAlihouseProjectActivitySyncAPIRequest对象
func NewAlibabaAlihouseProjectActivitySyncRequest() *AlibabaAlihouseProjectActivitySyncAPIRequest {
	return &AlibabaAlihouseProjectActivitySyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseProjectActivitySyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.project.activity.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseProjectActivitySyncAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetBusinessActivityDataDto is BusinessActivityDataDto Setter
// 数据
func (r *AlibabaAlihouseProjectActivitySyncAPIRequest) SetBusinessActivityDataDto(_businessActivityDataDto *BusinessActivityDataDto) error {
	r._businessActivityDataDto = _businessActivityDataDto
	r.Set("business_activity_data_dto", _businessActivityDataDto)
	return nil
}

// GetBusinessActivityDataDto BusinessActivityDataDto Getter
func (r AlibabaAlihouseProjectActivitySyncAPIRequest) GetBusinessActivityDataDto() *BusinessActivityDataDto {
	return r._businessActivityDataDto
}
