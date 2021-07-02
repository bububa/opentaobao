package wdk

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWdkEquipmentConveyorInfoUploadAPIRequest) GetApiMethodName() string {
	return "taobao.wdk.equipment.conveyor.info.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWdkEquipmentConveyorInfoUploadAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
