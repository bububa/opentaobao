package sungari

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosungaridisposesubmitAPIRequest 商品商家处置提交任务 API请求
// taobao.sungari.dispose.submit
//
// 商品商家处置信息接口，提供政府部门发送处置信息给阿里
type TaobaosungaridisposesubmitAPIRequest struct {
	model.Params
	// 平台处置信息入参
	_info *DisposeInfoDo
}

// NewTaobaosungaridisposesubmitRequest 初始化TaobaosungaridisposesubmitAPIRequest对象
func NewTaobaosungaridisposesubmitRequest() *TaobaosungaridisposesubmitAPIRequest {
	return &TaobaosungaridisposesubmitAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosungaridisposesubmitAPIRequest) GetApiMethodName() string {
	return "taobao.sungari.dispose.submit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosungaridisposesubmitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosungaridisposesubmitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInfo is Info Setter
// 平台处置信息入参
func (r *TaobaosungaridisposesubmitAPIRequest) SetInfo(_info *DisposeInfoDo) error {
	r._info = _info
	r.Set("info", _info)
	return nil
}

// GetInfo Info Getter
func (r TaobaosungaridisposesubmitAPIRequest) GetInfo() *DisposeInfoDo {
	return r._info
}
