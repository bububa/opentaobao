package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWdkEquipmentWcsWcsinfoUploadAPIRequest 悬挂链业务信息上传 API请求
// taobao.wdk.equipment.wcs.wcsinfo.upload
//
// 五道口仓库悬挂链信息上传
type TaobaoWdkEquipmentWcsWcsinfoUploadAPIRequest struct {
	model.Params
	// 上传信息
	_param0 string
}

// NewTaobaoWdkEquipmentWcsWcsinfoUploadRequest 初始化TaobaoWdkEquipmentWcsWcsinfoUploadAPIRequest对象
func NewTaobaoWdkEquipmentWcsWcsinfoUploadRequest() *TaobaoWdkEquipmentWcsWcsinfoUploadAPIRequest {
	return &TaobaoWdkEquipmentWcsWcsinfoUploadAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoWdkEquipmentWcsWcsinfoUploadAPIRequest) Reset() {
	r._param0 = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWdkEquipmentWcsWcsinfoUploadAPIRequest) GetApiMethodName() string {
	return "taobao.wdk.equipment.wcs.wcsinfo.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWdkEquipmentWcsWcsinfoUploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoWdkEquipmentWcsWcsinfoUploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 上传信息
func (r *TaobaoWdkEquipmentWcsWcsinfoUploadAPIRequest) SetParam0(_param0 string) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r TaobaoWdkEquipmentWcsWcsinfoUploadAPIRequest) GetParam0() string {
	return r._param0
}

var poolTaobaoWdkEquipmentWcsWcsinfoUploadAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoWdkEquipmentWcsWcsinfoUploadRequest()
	},
}

// GetTaobaoWdkEquipmentWcsWcsinfoUploadRequest 从 sync.Pool 获取 TaobaoWdkEquipmentWcsWcsinfoUploadAPIRequest
func GetTaobaoWdkEquipmentWcsWcsinfoUploadAPIRequest() *TaobaoWdkEquipmentWcsWcsinfoUploadAPIRequest {
	return poolTaobaoWdkEquipmentWcsWcsinfoUploadAPIRequest.Get().(*TaobaoWdkEquipmentWcsWcsinfoUploadAPIRequest)
}

// ReleaseTaobaoWdkEquipmentWcsWcsinfoUploadAPIRequest 将 TaobaoWdkEquipmentWcsWcsinfoUploadAPIRequest 放入 sync.Pool
func ReleaseTaobaoWdkEquipmentWcsWcsinfoUploadAPIRequest(v *TaobaoWdkEquipmentWcsWcsinfoUploadAPIRequest) {
	v.Reset()
	poolTaobaoWdkEquipmentWcsWcsinfoUploadAPIRequest.Put(v)
}
