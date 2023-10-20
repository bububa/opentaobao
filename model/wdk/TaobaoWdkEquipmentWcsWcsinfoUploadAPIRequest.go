package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaowdkequipmentwcswcsinfouploadAPIRequest 悬挂链业务信息上传 API请求
// taobao.wdk.equipment.wcs.wcsinfo.upload
//
// 五道口仓库悬挂链信息上传
type TaobaowdkequipmentwcswcsinfouploadAPIRequest struct {
	model.Params
	// 上传信息
	_param0 string
}

// NewTaobaowdkequipmentwcswcsinfouploadRequest 初始化TaobaowdkequipmentwcswcsinfouploadAPIRequest对象
func NewTaobaowdkequipmentwcswcsinfouploadRequest() *TaobaowdkequipmentwcswcsinfouploadAPIRequest {
	return &TaobaowdkequipmentwcswcsinfouploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaowdkequipmentwcswcsinfouploadAPIRequest) GetApiMethodName() string {
	return "taobao.wdk.equipment.wcs.wcsinfo.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaowdkequipmentwcswcsinfouploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaowdkequipmentwcswcsinfouploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 上传信息
func (r *TaobaowdkequipmentwcswcsinfouploadAPIRequest) SetParam0(_param0 string) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r TaobaowdkequipmentwcswcsinfouploadAPIRequest) GetParam0() string {
	return r._param0
}
