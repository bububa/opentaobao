package waybill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoCloudprintIsvResourcesGetAPIRequest
isv资源查询 API请求
cainiao.cloudprint.isv.resources.get

isv资源查询，包括isv模板、打印项、预设的自定义区等 */
type CainiaoCloudprintIsvResourcesGetAPIRequest struct {
	model.Params
	// isv资源类型，分为：TEMPLATE（表示模板），PRINT_ITEM（打印项），CUSTOM_AREA（预设自定义区）
	_isvResourceType string
}

// NewCainiaoCloudprintIsvResourcesGetRequest 初始化CainiaoCloudprintIsvResourcesGetAPIRequest对象
func NewCainiaoCloudprintIsvResourcesGetRequest() *CainiaoCloudprintIsvResourcesGetAPIRequest {
	return &CainiaoCloudprintIsvResourcesGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoCloudprintIsvResourcesGetAPIRequest) GetApiMethodName() string {
	return "cainiao.cloudprint.isv.resources.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoCloudprintIsvResourcesGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is IsvResourceType Setter
// isv资源类型，分为：TEMPLATE（表示模板），PRINT_ITEM（打印项），CUSTOM_AREA（预设自定义区）
func (r *CainiaoCloudprintIsvResourcesGetAPIRequest) SetIsvResourceType(_isvResourceType string) error {
	r._isvResourceType = _isvResourceType
	r.Set("isv_resource_type", _isvResourceType)
	return nil
}

// Get IsvResourceType Getter
func (r CainiaoCloudprintIsvResourcesGetAPIRequest) GetIsvResourceType() string {
	return r._isvResourceType
}
