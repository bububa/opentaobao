package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWdkEquipmentConveyorInfoUploadAPIRequest 五道口仓库悬挂链信息上报 API请求
// taobao.wdk.equipment.conveyor.info.upload
//
// 五道口仓库悬挂链信息上传
type TaobaoWdkEquipmentConveyorInfoUploadAPIRequest struct {
	model.Params
	// 上传信息
	_param0 string
}

// NewTaobaoWdkEquipmentConveyorInfoUploadRequest 初始化TaobaoWdkEquipmentConveyorInfoUploadAPIRequest对象
func NewTaobaoWdkEquipmentConveyorInfoUploadRequest() *TaobaoWdkEquipmentConveyorInfoUploadAPIRequest {
	return &TaobaoWdkEquipmentConveyorInfoUploadAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoWdkEquipmentConveyorInfoUploadAPIRequest) Reset() {
	r._param0 = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWdkEquipmentConveyorInfoUploadAPIRequest) GetApiMethodName() string {
	return "taobao.wdk.equipment.conveyor.info.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWdkEquipmentConveyorInfoUploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoWdkEquipmentConveyorInfoUploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 上传信息
func (r *TaobaoWdkEquipmentConveyorInfoUploadAPIRequest) SetParam0(_param0 string) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r TaobaoWdkEquipmentConveyorInfoUploadAPIRequest) GetParam0() string {
	return r._param0
}

var poolTaobaoWdkEquipmentConveyorInfoUploadAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoWdkEquipmentConveyorInfoUploadRequest()
	},
}

// GetTaobaoWdkEquipmentConveyorInfoUploadRequest 从 sync.Pool 获取 TaobaoWdkEquipmentConveyorInfoUploadAPIRequest
func GetTaobaoWdkEquipmentConveyorInfoUploadAPIRequest() *TaobaoWdkEquipmentConveyorInfoUploadAPIRequest {
	return poolTaobaoWdkEquipmentConveyorInfoUploadAPIRequest.Get().(*TaobaoWdkEquipmentConveyorInfoUploadAPIRequest)
}

// ReleaseTaobaoWdkEquipmentConveyorInfoUploadAPIRequest 将 TaobaoWdkEquipmentConveyorInfoUploadAPIRequest 放入 sync.Pool
func ReleaseTaobaoWdkEquipmentConveyorInfoUploadAPIRequest(v *TaobaoWdkEquipmentConveyorInfoUploadAPIRequest) {
	v.Reset()
	poolTaobaoWdkEquipmentConveyorInfoUploadAPIRequest.Put(v)
}
