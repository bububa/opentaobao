package gameact

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoDeActivityMachineidGetAPIRequest 获取设备号 API请求
// taobao.de.activity.machineid.get
//
// 获取机器设备id
type TaobaoDeActivityMachineidGetAPIRequest struct {
	model.Params
}

// NewTaobaoDeActivityMachineidGetRequest 初始化TaobaoDeActivityMachineidGetAPIRequest对象
func NewTaobaoDeActivityMachineidGetRequest() *TaobaoDeActivityMachineidGetAPIRequest {
	return &TaobaoDeActivityMachineidGetAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoDeActivityMachineidGetAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoDeActivityMachineidGetAPIRequest) GetApiMethodName() string {
	return "taobao.de.activity.machineid.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoDeActivityMachineidGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoDeActivityMachineidGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolTaobaoDeActivityMachineidGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoDeActivityMachineidGetRequest()
	},
}

// GetTaobaoDeActivityMachineidGetRequest 从 sync.Pool 获取 TaobaoDeActivityMachineidGetAPIRequest
func GetTaobaoDeActivityMachineidGetAPIRequest() *TaobaoDeActivityMachineidGetAPIRequest {
	return poolTaobaoDeActivityMachineidGetAPIRequest.Get().(*TaobaoDeActivityMachineidGetAPIRequest)
}

// ReleaseTaobaoDeActivityMachineidGetAPIRequest 将 TaobaoDeActivityMachineidGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoDeActivityMachineidGetAPIRequest(v *TaobaoDeActivityMachineidGetAPIRequest) {
	v.Reset()
	poolTaobaoDeActivityMachineidGetAPIRequest.Put(v)
}
