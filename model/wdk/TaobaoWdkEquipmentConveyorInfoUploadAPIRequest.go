package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaowdkequipmentconveyorinfouploadAPIRequest 五道口仓库悬挂链信息上报 API请求
// taobao.wdk.equipment.conveyor.info.upload
//
// 五道口仓库悬挂链信息上传
type TaobaowdkequipmentconveyorinfouploadAPIRequest struct {
	model.Params
	// 上传信息
	_param0 string
}

// NewTaobaowdkequipmentconveyorinfouploadRequest 初始化TaobaowdkequipmentconveyorinfouploadAPIRequest对象
func NewTaobaowdkequipmentconveyorinfouploadRequest() *TaobaowdkequipmentconveyorinfouploadAPIRequest {
	return &TaobaowdkequipmentconveyorinfouploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaowdkequipmentconveyorinfouploadAPIRequest) GetApiMethodName() string {
	return "taobao.wdk.equipment.conveyor.info.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaowdkequipmentconveyorinfouploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaowdkequipmentconveyorinfouploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 上传信息
func (r *TaobaowdkequipmentconveyorinfouploadAPIRequest) SetParam0(_param0 string) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r TaobaowdkequipmentconveyorinfouploadAPIRequest) GetParam0() string {
	return r._param0
}
