package waybill

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoCloudprintCustomareaUpdateAPIRequest 自定义区内容更新 API请求
// cainiao.cloudprint.customarea.update
//
// 自定义区内容更新
type CainiaoCloudprintCustomareaUpdateAPIRequest struct {
	model.Params
	// 自定义区名称（可修改）
	_customAreaName string
	// 自定义区内容（可修改）
	_customAreaContent string
	// 自定义区id（不可修改）
	_customAreaId int64
}

// NewCainiaoCloudprintCustomareaUpdateRequest 初始化CainiaoCloudprintCustomareaUpdateAPIRequest对象
func NewCainiaoCloudprintCustomareaUpdateRequest() *CainiaoCloudprintCustomareaUpdateAPIRequest {
	return &CainiaoCloudprintCustomareaUpdateAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *CainiaoCloudprintCustomareaUpdateAPIRequest) Reset() {
	r._customAreaName = ""
	r._customAreaContent = ""
	r._customAreaId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoCloudprintCustomareaUpdateAPIRequest) GetApiMethodName() string {
	return "cainiao.cloudprint.customarea.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoCloudprintCustomareaUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoCloudprintCustomareaUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCustomAreaName is CustomAreaName Setter
// 自定义区名称（可修改）
func (r *CainiaoCloudprintCustomareaUpdateAPIRequest) SetCustomAreaName(_customAreaName string) error {
	r._customAreaName = _customAreaName
	r.Set("custom_area_name", _customAreaName)
	return nil
}

// GetCustomAreaName CustomAreaName Getter
func (r CainiaoCloudprintCustomareaUpdateAPIRequest) GetCustomAreaName() string {
	return r._customAreaName
}

// SetCustomAreaContent is CustomAreaContent Setter
// 自定义区内容（可修改）
func (r *CainiaoCloudprintCustomareaUpdateAPIRequest) SetCustomAreaContent(_customAreaContent string) error {
	r._customAreaContent = _customAreaContent
	r.Set("custom_area_content", _customAreaContent)
	return nil
}

// GetCustomAreaContent CustomAreaContent Getter
func (r CainiaoCloudprintCustomareaUpdateAPIRequest) GetCustomAreaContent() string {
	return r._customAreaContent
}

// SetCustomAreaId is CustomAreaId Setter
// 自定义区id（不可修改）
func (r *CainiaoCloudprintCustomareaUpdateAPIRequest) SetCustomAreaId(_customAreaId int64) error {
	r._customAreaId = _customAreaId
	r.Set("custom_area_id", _customAreaId)
	return nil
}

// GetCustomAreaId CustomAreaId Getter
func (r CainiaoCloudprintCustomareaUpdateAPIRequest) GetCustomAreaId() int64 {
	return r._customAreaId
}

var poolCainiaoCloudprintCustomareaUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewCainiaoCloudprintCustomareaUpdateRequest()
	},
}

// GetCainiaoCloudprintCustomareaUpdateRequest 从 sync.Pool 获取 CainiaoCloudprintCustomareaUpdateAPIRequest
func GetCainiaoCloudprintCustomareaUpdateAPIRequest() *CainiaoCloudprintCustomareaUpdateAPIRequest {
	return poolCainiaoCloudprintCustomareaUpdateAPIRequest.Get().(*CainiaoCloudprintCustomareaUpdateAPIRequest)
}

// ReleaseCainiaoCloudprintCustomareaUpdateAPIRequest 将 CainiaoCloudprintCustomareaUpdateAPIRequest 放入 sync.Pool
func ReleaseCainiaoCloudprintCustomareaUpdateAPIRequest(v *CainiaoCloudprintCustomareaUpdateAPIRequest) {
	v.Reset()
	poolCainiaoCloudprintCustomareaUpdateAPIRequest.Put(v)
}
