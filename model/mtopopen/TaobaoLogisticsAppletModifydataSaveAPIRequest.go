package mtopopen

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsAppletModifydataSaveAPIRequest 物流小程序修改物流信息回传接口 API请求
// taobao.logistics.applet.modifydata.save
//
// 物流小程序修改物流信息回传接口
type TaobaoLogisticsAppletModifydataSaveAPIRequest struct {
	model.Params
	// 请求对象
	_modifyRequest *ModifyDeliveryRequest
}

// NewTaobaoLogisticsAppletModifydataSaveRequest 初始化TaobaoLogisticsAppletModifydataSaveAPIRequest对象
func NewTaobaoLogisticsAppletModifydataSaveRequest() *TaobaoLogisticsAppletModifydataSaveAPIRequest {
	return &TaobaoLogisticsAppletModifydataSaveAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoLogisticsAppletModifydataSaveAPIRequest) Reset() {
	r._modifyRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoLogisticsAppletModifydataSaveAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.applet.modifydata.save"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoLogisticsAppletModifydataSaveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoLogisticsAppletModifydataSaveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetModifyRequest is ModifyRequest Setter
// 请求对象
func (r *TaobaoLogisticsAppletModifydataSaveAPIRequest) SetModifyRequest(_modifyRequest *ModifyDeliveryRequest) error {
	r._modifyRequest = _modifyRequest
	r.Set("modify_request", _modifyRequest)
	return nil
}

// GetModifyRequest ModifyRequest Getter
func (r TaobaoLogisticsAppletModifydataSaveAPIRequest) GetModifyRequest() *ModifyDeliveryRequest {
	return r._modifyRequest
}

var poolTaobaoLogisticsAppletModifydataSaveAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoLogisticsAppletModifydataSaveRequest()
	},
}

// GetTaobaoLogisticsAppletModifydataSaveRequest 从 sync.Pool 获取 TaobaoLogisticsAppletModifydataSaveAPIRequest
func GetTaobaoLogisticsAppletModifydataSaveAPIRequest() *TaobaoLogisticsAppletModifydataSaveAPIRequest {
	return poolTaobaoLogisticsAppletModifydataSaveAPIRequest.Get().(*TaobaoLogisticsAppletModifydataSaveAPIRequest)
}

// ReleaseTaobaoLogisticsAppletModifydataSaveAPIRequest 将 TaobaoLogisticsAppletModifydataSaveAPIRequest 放入 sync.Pool
func ReleaseTaobaoLogisticsAppletModifydataSaveAPIRequest(v *TaobaoLogisticsAppletModifydataSaveAPIRequest) {
	v.Reset()
	poolTaobaoLogisticsAppletModifydataSaveAPIRequest.Put(v)
}
