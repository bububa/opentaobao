package waybill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaocloudprintisvresourcesgetAPIRequest isv资源查询 API请求
// cainiao.cloudprint.isv.resources.get
//
// isv资源查询，包括isv模板、打印项、预设的自定义区等
type CainiaocloudprintisvresourcesgetAPIRequest struct {
	model.Params
	// isv资源类型，分为：TEMPLATE（表示模板），PRINT_ITEM（打印项），CUSTOM_AREA（预设自定义区）
	_isvResourceType string
}

// NewCainiaocloudprintisvresourcesgetRequest 初始化CainiaocloudprintisvresourcesgetAPIRequest对象
func NewCainiaocloudprintisvresourcesgetRequest() *CainiaocloudprintisvresourcesgetAPIRequest {
	return &CainiaocloudprintisvresourcesgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaocloudprintisvresourcesgetAPIRequest) GetApiMethodName() string {
	return "cainiao.cloudprint.isv.resources.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaocloudprintisvresourcesgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaocloudprintisvresourcesgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetIsvResourceType is IsvResourceType Setter
// isv资源类型，分为：TEMPLATE（表示模板），PRINT_ITEM（打印项），CUSTOM_AREA（预设自定义区）
func (r *CainiaocloudprintisvresourcesgetAPIRequest) SetIsvResourceType(_isvResourceType string) error {
	r._isvResourceType = _isvResourceType
	r.Set("isv_resource_type", _isvResourceType)
	return nil
}

// GetIsvResourceType IsvResourceType Getter
func (r CainiaocloudprintisvresourcesgetAPIRequest) GetIsvResourceType() string {
	return r._isvResourceType
}
