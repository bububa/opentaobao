package baichuan

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBaichuanItemSubscribeDailyLeftQueryAPIRequest 查询当天可添加的余量 API请求
// taobao.baichuan.item.subscribe.daily.left.query
//
// 查询当天可添加的余量
type TaobaoBaichuanItemSubscribeDailyLeftQueryAPIRequest struct {
	model.Params
}

// NewTaobaoBaichuanItemSubscribeDailyLeftQueryRequest 初始化TaobaoBaichuanItemSubscribeDailyLeftQueryAPIRequest对象
func NewTaobaoBaichuanItemSubscribeDailyLeftQueryRequest() *TaobaoBaichuanItemSubscribeDailyLeftQueryAPIRequest {
	return &TaobaoBaichuanItemSubscribeDailyLeftQueryAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoBaichuanItemSubscribeDailyLeftQueryAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBaichuanItemSubscribeDailyLeftQueryAPIRequest) GetApiMethodName() string {
	return "taobao.baichuan.item.subscribe.daily.left.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBaichuanItemSubscribeDailyLeftQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoBaichuanItemSubscribeDailyLeftQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolTaobaoBaichuanItemSubscribeDailyLeftQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoBaichuanItemSubscribeDailyLeftQueryRequest()
	},
}

// GetTaobaoBaichuanItemSubscribeDailyLeftQueryRequest 从 sync.Pool 获取 TaobaoBaichuanItemSubscribeDailyLeftQueryAPIRequest
func GetTaobaoBaichuanItemSubscribeDailyLeftQueryAPIRequest() *TaobaoBaichuanItemSubscribeDailyLeftQueryAPIRequest {
	return poolTaobaoBaichuanItemSubscribeDailyLeftQueryAPIRequest.Get().(*TaobaoBaichuanItemSubscribeDailyLeftQueryAPIRequest)
}

// ReleaseTaobaoBaichuanItemSubscribeDailyLeftQueryAPIRequest 将 TaobaoBaichuanItemSubscribeDailyLeftQueryAPIRequest 放入 sync.Pool
func ReleaseTaobaoBaichuanItemSubscribeDailyLeftQueryAPIRequest(v *TaobaoBaichuanItemSubscribeDailyLeftQueryAPIRequest) {
	v.Reset()
	poolTaobaoBaichuanItemSubscribeDailyLeftQueryAPIRequest.Put(v)
}
