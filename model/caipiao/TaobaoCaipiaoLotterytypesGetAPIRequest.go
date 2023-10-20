package caipiao

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCaipiaoLotterytypesGetAPIRequest 获取可用的彩种列表 API请求
// taobao.caipiao.lotterytypes.get
//
// 获取彩票系统支持的可用于赠送的彩种列表
type TaobaoCaipiaoLotterytypesGetAPIRequest struct {
	model.Params
}

// NewTaobaoCaipiaoLotterytypesGetRequest 初始化TaobaoCaipiaoLotterytypesGetAPIRequest对象
func NewTaobaoCaipiaoLotterytypesGetRequest() *TaobaoCaipiaoLotterytypesGetAPIRequest {
	return &TaobaoCaipiaoLotterytypesGetAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoCaipiaoLotterytypesGetAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoCaipiaoLotterytypesGetAPIRequest) GetApiMethodName() string {
	return "taobao.caipiao.lotterytypes.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoCaipiaoLotterytypesGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoCaipiaoLotterytypesGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolTaobaoCaipiaoLotterytypesGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoCaipiaoLotterytypesGetRequest()
	},
}

// GetTaobaoCaipiaoLotterytypesGetRequest 从 sync.Pool 获取 TaobaoCaipiaoLotterytypesGetAPIRequest
func GetTaobaoCaipiaoLotterytypesGetAPIRequest() *TaobaoCaipiaoLotterytypesGetAPIRequest {
	return poolTaobaoCaipiaoLotterytypesGetAPIRequest.Get().(*TaobaoCaipiaoLotterytypesGetAPIRequest)
}

// ReleaseTaobaoCaipiaoLotterytypesGetAPIRequest 将 TaobaoCaipiaoLotterytypesGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoCaipiaoLotterytypesGetAPIRequest(v *TaobaoCaipiaoLotterytypesGetAPIRequest) {
	v.Reset()
	poolTaobaoCaipiaoLotterytypesGetAPIRequest.Put(v)
}
