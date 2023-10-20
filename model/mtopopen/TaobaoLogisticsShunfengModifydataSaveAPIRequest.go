package mtopopen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaologisticsshunfengmodifydatasaveAPIRequest 顺丰小程序修改配送信息回传接口 API请求
// taobao.logistics.shunfeng.modifydata.save
//
// 顺丰小程序修改配送信息回传接口
type TaobaologisticsshunfengmodifydatasaveAPIRequest struct {
	model.Params
	// 请求对象
	_modifyRequest *ModifyRequest
}

// NewTaobaologisticsshunfengmodifydatasaveRequest 初始化TaobaologisticsshunfengmodifydatasaveAPIRequest对象
func NewTaobaologisticsshunfengmodifydatasaveRequest() *TaobaologisticsshunfengmodifydatasaveAPIRequest {
	return &TaobaologisticsshunfengmodifydatasaveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaologisticsshunfengmodifydatasaveAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.shunfeng.modifydata.save"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaologisticsshunfengmodifydatasaveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaologisticsshunfengmodifydatasaveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetModifyRequest is ModifyRequest Setter
// 请求对象
func (r *TaobaologisticsshunfengmodifydatasaveAPIRequest) SetModifyRequest(_modifyRequest *ModifyRequest) error {
	r._modifyRequest = _modifyRequest
	r.Set("modify_request", _modifyRequest)
	return nil
}

// GetModifyRequest ModifyRequest Getter
func (r TaobaologisticsshunfengmodifydatasaveAPIRequest) GetModifyRequest() *ModifyRequest {
	return r._modifyRequest
}
