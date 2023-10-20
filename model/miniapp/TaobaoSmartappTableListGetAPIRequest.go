package miniapp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSmartappTableListGetAPIRequest 智能应用服务登记工作表列表查询 API请求
// taobao.smartapp.table.list.get
//
// 智能应用服务登记工作表列表查询
type TaobaoSmartappTableListGetAPIRequest struct {
	model.Params
}

// NewTaobaoSmartappTableListGetRequest 初始化TaobaoSmartappTableListGetAPIRequest对象
func NewTaobaoSmartappTableListGetRequest() *TaobaoSmartappTableListGetAPIRequest {
	return &TaobaoSmartappTableListGetAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSmartappTableListGetAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSmartappTableListGetAPIRequest) GetApiMethodName() string {
	return "taobao.smartapp.table.list.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSmartappTableListGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSmartappTableListGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolTaobaoSmartappTableListGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSmartappTableListGetRequest()
	},
}

// GetTaobaoSmartappTableListGetRequest 从 sync.Pool 获取 TaobaoSmartappTableListGetAPIRequest
func GetTaobaoSmartappTableListGetAPIRequest() *TaobaoSmartappTableListGetAPIRequest {
	return poolTaobaoSmartappTableListGetAPIRequest.Get().(*TaobaoSmartappTableListGetAPIRequest)
}

// ReleaseTaobaoSmartappTableListGetAPIRequest 将 TaobaoSmartappTableListGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoSmartappTableListGetAPIRequest(v *TaobaoSmartappTableListGetAPIRequest) {
	v.Reset()
	poolTaobaoSmartappTableListGetAPIRequest.Put(v)
}
