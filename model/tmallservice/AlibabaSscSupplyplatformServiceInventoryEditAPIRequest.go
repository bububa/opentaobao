package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSscSupplyplatformServiceInventoryEditAPIRequest 编辑服务库存 API请求
// alibaba.ssc.supplyplatform.service.inventory.edit
//
// 实时编辑服务库存。只支持增加或减少库存，不支持设置绝对库存值。
// 需要自己处理好幂等逻辑。
// 要先查询当前库存值，并基于返回结果做编辑操作。
// 参考alibaba.ssc.supplyplatform.service.inventory.query和alibaba.ssc.supplyplatform.servicecapacity.save
type AlibabaSscSupplyplatformServiceInventoryEditAPIRequest struct {
	model.Params
	// 服务提供者类型。参考alibaba.ssc.supplyplatform.servicecapacity.save入参
	_providerType string
	// 服务提供者id。参考alibaba.ssc.supplyplatform.servicecapacity.save入参
	_providerId int64
	// 业务幂等键。该字段主要用于远程调用失败后的重试的场景，例如接口超时，调用方感知到失败，但服务端可能实际上已经成功了，这时如果发起一次重试请求，服务端需要通过bizId来识别是同一个请求，这样才不会重复增加库存值。对于同一个bizId，多次请求只会生效一次，后续的重复请求不会生效。对于批量操作时，如果部分key成功，部分key失败，重试请求时只会对未成功的key生效。
	_bizId string
	// 库存编辑列表。每次不超过100条
	_editDetails []EditDetailInventoryRequest
}

// NewAlibabaSscSupplyplatformServiceInventoryEditRequest 初始化AlibabaSscSupplyplatformServiceInventoryEditAPIRequest对象
func NewAlibabaSscSupplyplatformServiceInventoryEditRequest() *AlibabaSscSupplyplatformServiceInventoryEditAPIRequest {
	return &AlibabaSscSupplyplatformServiceInventoryEditAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSscSupplyplatformServiceInventoryEditAPIRequest) GetApiMethodName() string {
	return "alibaba.ssc.supplyplatform.service.inventory.edit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSscSupplyplatformServiceInventoryEditAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ProviderType Setter
// 服务提供者类型。参考alibaba.ssc.supplyplatform.servicecapacity.save入参
func (r *AlibabaSscSupplyplatformServiceInventoryEditAPIRequest) SetProviderType(_providerType string) error {
	r._providerType = _providerType
	r.Set("provider_type", _providerType)
	return nil
}

// Get ProviderType Getter
func (r AlibabaSscSupplyplatformServiceInventoryEditAPIRequest) GetProviderType() string {
	return r._providerType
}

// Set is ProviderId Setter
// 服务提供者id。参考alibaba.ssc.supplyplatform.servicecapacity.save入参
func (r *AlibabaSscSupplyplatformServiceInventoryEditAPIRequest) SetProviderId(_providerId int64) error {
	r._providerId = _providerId
	r.Set("provider_id", _providerId)
	return nil
}

// Get ProviderId Getter
func (r AlibabaSscSupplyplatformServiceInventoryEditAPIRequest) GetProviderId() int64 {
	return r._providerId
}

// Set is BizId Setter
// 业务幂等键。该字段主要用于远程调用失败后的重试的场景，例如接口超时，调用方感知到失败，但服务端可能实际上已经成功了，这时如果发起一次重试请求，服务端需要通过bizId来识别是同一个请求，这样才不会重复增加库存值。对于同一个bizId，多次请求只会生效一次，后续的重复请求不会生效。对于批量操作时，如果部分key成功，部分key失败，重试请求时只会对未成功的key生效。
func (r *AlibabaSscSupplyplatformServiceInventoryEditAPIRequest) SetBizId(_bizId string) error {
	r._bizId = _bizId
	r.Set("biz_id", _bizId)
	return nil
}

// Get BizId Getter
func (r AlibabaSscSupplyplatformServiceInventoryEditAPIRequest) GetBizId() string {
	return r._bizId
}

// Set is EditDetails Setter
// 库存编辑列表。每次不超过100条
func (r *AlibabaSscSupplyplatformServiceInventoryEditAPIRequest) SetEditDetails(_editDetails []EditDetailInventoryRequest) error {
	r._editDetails = _editDetails
	r.Set("edit_details", _editDetails)
	return nil
}

// Get EditDetails Getter
func (r AlibabaSscSupplyplatformServiceInventoryEditAPIRequest) GetEditDetails() []EditDetailInventoryRequest {
	return r._editDetails
}
