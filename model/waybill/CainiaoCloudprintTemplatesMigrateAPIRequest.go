package waybill

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoCloudprintTemplatesMigrateAPIRequest 云打印模板迁移接口 API请求
// cainiao.cloudprint.templates.migrate
//
// 云打印模板迁移接口
type CainiaoCloudprintTemplatesMigrateAPIRequest struct {
	model.Params
	// 自定义区名称
	_customAreaName string
	// 自定义区内容
	_customAreaContent string
	// 标准电子面单模板的id
	_tempalteId int64
}

// NewCainiaoCloudprintTemplatesMigrateRequest 初始化CainiaoCloudprintTemplatesMigrateAPIRequest对象
func NewCainiaoCloudprintTemplatesMigrateRequest() *CainiaoCloudprintTemplatesMigrateAPIRequest {
	return &CainiaoCloudprintTemplatesMigrateAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *CainiaoCloudprintTemplatesMigrateAPIRequest) Reset() {
	r._customAreaName = ""
	r._customAreaContent = ""
	r._tempalteId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoCloudprintTemplatesMigrateAPIRequest) GetApiMethodName() string {
	return "cainiao.cloudprint.templates.migrate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoCloudprintTemplatesMigrateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoCloudprintTemplatesMigrateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCustomAreaName is CustomAreaName Setter
// 自定义区名称
func (r *CainiaoCloudprintTemplatesMigrateAPIRequest) SetCustomAreaName(_customAreaName string) error {
	r._customAreaName = _customAreaName
	r.Set("custom_area_name", _customAreaName)
	return nil
}

// GetCustomAreaName CustomAreaName Getter
func (r CainiaoCloudprintTemplatesMigrateAPIRequest) GetCustomAreaName() string {
	return r._customAreaName
}

// SetCustomAreaContent is CustomAreaContent Setter
// 自定义区内容
func (r *CainiaoCloudprintTemplatesMigrateAPIRequest) SetCustomAreaContent(_customAreaContent string) error {
	r._customAreaContent = _customAreaContent
	r.Set("custom_area_content", _customAreaContent)
	return nil
}

// GetCustomAreaContent CustomAreaContent Getter
func (r CainiaoCloudprintTemplatesMigrateAPIRequest) GetCustomAreaContent() string {
	return r._customAreaContent
}

// SetTempalteId is TempalteId Setter
// 标准电子面单模板的id
func (r *CainiaoCloudprintTemplatesMigrateAPIRequest) SetTempalteId(_tempalteId int64) error {
	r._tempalteId = _tempalteId
	r.Set("tempalte_id", _tempalteId)
	return nil
}

// GetTempalteId TempalteId Getter
func (r CainiaoCloudprintTemplatesMigrateAPIRequest) GetTempalteId() int64 {
	return r._tempalteId
}

var poolCainiaoCloudprintTemplatesMigrateAPIRequest = sync.Pool{
	New: func() any {
		return NewCainiaoCloudprintTemplatesMigrateRequest()
	},
}

// GetCainiaoCloudprintTemplatesMigrateRequest 从 sync.Pool 获取 CainiaoCloudprintTemplatesMigrateAPIRequest
func GetCainiaoCloudprintTemplatesMigrateAPIRequest() *CainiaoCloudprintTemplatesMigrateAPIRequest {
	return poolCainiaoCloudprintTemplatesMigrateAPIRequest.Get().(*CainiaoCloudprintTemplatesMigrateAPIRequest)
}

// ReleaseCainiaoCloudprintTemplatesMigrateAPIRequest 将 CainiaoCloudprintTemplatesMigrateAPIRequest 放入 sync.Pool
func ReleaseCainiaoCloudprintTemplatesMigrateAPIRequest(v *CainiaoCloudprintTemplatesMigrateAPIRequest) {
	v.Reset()
	poolCainiaoCloudprintTemplatesMigrateAPIRequest.Put(v)
}
