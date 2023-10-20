package waybill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaocloudprintcustomareaupdateAPIRequest 自定义区内容更新 API请求
// cainiao.cloudprint.customarea.update
//
// 自定义区内容更新
type CainiaocloudprintcustomareaupdateAPIRequest struct {
	model.Params
	// 自定义区名称（可修改）
	_customAreaName string
	// 自定义区内容（可修改）
	_customAreaContent string
	// 自定义区id（不可修改）
	_customAreaId int64
}

// NewCainiaocloudprintcustomareaupdateRequest 初始化CainiaocloudprintcustomareaupdateAPIRequest对象
func NewCainiaocloudprintcustomareaupdateRequest() *CainiaocloudprintcustomareaupdateAPIRequest {
	return &CainiaocloudprintcustomareaupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaocloudprintcustomareaupdateAPIRequest) GetApiMethodName() string {
	return "cainiao.cloudprint.customarea.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaocloudprintcustomareaupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaocloudprintcustomareaupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCustomAreaName is CustomAreaName Setter
// 自定义区名称（可修改）
func (r *CainiaocloudprintcustomareaupdateAPIRequest) SetCustomAreaName(_customAreaName string) error {
	r._customAreaName = _customAreaName
	r.Set("custom_area_name", _customAreaName)
	return nil
}

// GetCustomAreaName CustomAreaName Getter
func (r CainiaocloudprintcustomareaupdateAPIRequest) GetCustomAreaName() string {
	return r._customAreaName
}

// SetCustomAreaContent is CustomAreaContent Setter
// 自定义区内容（可修改）
func (r *CainiaocloudprintcustomareaupdateAPIRequest) SetCustomAreaContent(_customAreaContent string) error {
	r._customAreaContent = _customAreaContent
	r.Set("custom_area_content", _customAreaContent)
	return nil
}

// GetCustomAreaContent CustomAreaContent Getter
func (r CainiaocloudprintcustomareaupdateAPIRequest) GetCustomAreaContent() string {
	return r._customAreaContent
}

// SetCustomAreaId is CustomAreaId Setter
// 自定义区id（不可修改）
func (r *CainiaocloudprintcustomareaupdateAPIRequest) SetCustomAreaId(_customAreaId int64) error {
	r._customAreaId = _customAreaId
	r.Set("custom_area_id", _customAreaId)
	return nil
}

// GetCustomAreaId CustomAreaId Getter
func (r CainiaocloudprintcustomareaupdateAPIRequest) GetCustomAreaId() int64 {
	return r._customAreaId
}
