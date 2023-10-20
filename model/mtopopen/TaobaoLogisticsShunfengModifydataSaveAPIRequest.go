package mtopopen

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsShunfengModifydataSaveAPIRequest 顺丰小程序修改配送信息回传接口 API请求
// taobao.logistics.shunfeng.modifydata.save
//
// 顺丰小程序修改配送信息回传接口
type TaobaoLogisticsShunfengModifydataSaveAPIRequest struct {
	model.Params
	// 请求对象
	_modifyRequest *ModifyRequest
}

// NewTaobaoLogisticsShunfengModifydataSaveRequest 初始化TaobaoLogisticsShunfengModifydataSaveAPIRequest对象
func NewTaobaoLogisticsShunfengModifydataSaveRequest() *TaobaoLogisticsShunfengModifydataSaveAPIRequest {
	return &TaobaoLogisticsShunfengModifydataSaveAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoLogisticsShunfengModifydataSaveAPIRequest) Reset() {
	r._modifyRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoLogisticsShunfengModifydataSaveAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.shunfeng.modifydata.save"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoLogisticsShunfengModifydataSaveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoLogisticsShunfengModifydataSaveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetModifyRequest is ModifyRequest Setter
// 请求对象
func (r *TaobaoLogisticsShunfengModifydataSaveAPIRequest) SetModifyRequest(_modifyRequest *ModifyRequest) error {
	r._modifyRequest = _modifyRequest
	r.Set("modify_request", _modifyRequest)
	return nil
}

// GetModifyRequest ModifyRequest Getter
func (r TaobaoLogisticsShunfengModifydataSaveAPIRequest) GetModifyRequest() *ModifyRequest {
	return r._modifyRequest
}

var poolTaobaoLogisticsShunfengModifydataSaveAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoLogisticsShunfengModifydataSaveRequest()
	},
}

// GetTaobaoLogisticsShunfengModifydataSaveRequest 从 sync.Pool 获取 TaobaoLogisticsShunfengModifydataSaveAPIRequest
func GetTaobaoLogisticsShunfengModifydataSaveAPIRequest() *TaobaoLogisticsShunfengModifydataSaveAPIRequest {
	return poolTaobaoLogisticsShunfengModifydataSaveAPIRequest.Get().(*TaobaoLogisticsShunfengModifydataSaveAPIRequest)
}

// ReleaseTaobaoLogisticsShunfengModifydataSaveAPIRequest 将 TaobaoLogisticsShunfengModifydataSaveAPIRequest 放入 sync.Pool
func ReleaseTaobaoLogisticsShunfengModifydataSaveAPIRequest(v *TaobaoLogisticsShunfengModifydataSaveAPIRequest) {
	v.Reset()
	poolTaobaoLogisticsShunfengModifydataSaveAPIRequest.Put(v)
}
