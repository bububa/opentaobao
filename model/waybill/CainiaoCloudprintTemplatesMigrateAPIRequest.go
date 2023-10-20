package waybill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaocloudprinttemplatesmigrateAPIRequest 云打印模板迁移接口 API请求
// cainiao.cloudprint.templates.migrate
//
// 云打印模板迁移接口
type CainiaocloudprinttemplatesmigrateAPIRequest struct {
	model.Params
	// 自定义区名称
	_customAreaName string
	// 自定义区内容
	_customAreaContent string
	// 标准电子面单模板的id
	_tempalteId int64
}

// NewCainiaocloudprinttemplatesmigrateRequest 初始化CainiaocloudprinttemplatesmigrateAPIRequest对象
func NewCainiaocloudprinttemplatesmigrateRequest() *CainiaocloudprinttemplatesmigrateAPIRequest {
	return &CainiaocloudprinttemplatesmigrateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaocloudprinttemplatesmigrateAPIRequest) GetApiMethodName() string {
	return "cainiao.cloudprint.templates.migrate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaocloudprinttemplatesmigrateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaocloudprinttemplatesmigrateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCustomAreaName is CustomAreaName Setter
// 自定义区名称
func (r *CainiaocloudprinttemplatesmigrateAPIRequest) SetCustomAreaName(_customAreaName string) error {
	r._customAreaName = _customAreaName
	r.Set("custom_area_name", _customAreaName)
	return nil
}

// GetCustomAreaName CustomAreaName Getter
func (r CainiaocloudprinttemplatesmigrateAPIRequest) GetCustomAreaName() string {
	return r._customAreaName
}

// SetCustomAreaContent is CustomAreaContent Setter
// 自定义区内容
func (r *CainiaocloudprinttemplatesmigrateAPIRequest) SetCustomAreaContent(_customAreaContent string) error {
	r._customAreaContent = _customAreaContent
	r.Set("custom_area_content", _customAreaContent)
	return nil
}

// GetCustomAreaContent CustomAreaContent Getter
func (r CainiaocloudprinttemplatesmigrateAPIRequest) GetCustomAreaContent() string {
	return r._customAreaContent
}

// SetTempalteId is TempalteId Setter
// 标准电子面单模板的id
func (r *CainiaocloudprinttemplatesmigrateAPIRequest) SetTempalteId(_tempalteId int64) error {
	r._tempalteId = _tempalteId
	r.Set("tempalte_id", _tempalteId)
	return nil
}

// GetTempalteId TempalteId Getter
func (r CainiaocloudprinttemplatesmigrateAPIRequest) GetTempalteId() int64 {
	return r._tempalteId
}
