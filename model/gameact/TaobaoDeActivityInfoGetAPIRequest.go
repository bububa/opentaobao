package gameact

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoDeActivityInfoGetAPIRequest 获取活动信息 API请求
// taobao.de.activity.info.get
//
// 根据appKey和活动id获取活动
type TaobaoDeActivityInfoGetAPIRequest struct {
	model.Params
	// 事件唯一标识
	_eventKey string
}

// NewTaobaoDeActivityInfoGetRequest 初始化TaobaoDeActivityInfoGetAPIRequest对象
func NewTaobaoDeActivityInfoGetRequest() *TaobaoDeActivityInfoGetAPIRequest {
	return &TaobaoDeActivityInfoGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoDeActivityInfoGetAPIRequest) Reset() {
	r._eventKey = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoDeActivityInfoGetAPIRequest) GetApiMethodName() string {
	return "taobao.de.activity.info.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoDeActivityInfoGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoDeActivityInfoGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEventKey is EventKey Setter
// 事件唯一标识
func (r *TaobaoDeActivityInfoGetAPIRequest) SetEventKey(_eventKey string) error {
	r._eventKey = _eventKey
	r.Set("event_key", _eventKey)
	return nil
}

// GetEventKey EventKey Getter
func (r TaobaoDeActivityInfoGetAPIRequest) GetEventKey() string {
	return r._eventKey
}

var poolTaobaoDeActivityInfoGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoDeActivityInfoGetRequest()
	},
}

// GetTaobaoDeActivityInfoGetRequest 从 sync.Pool 获取 TaobaoDeActivityInfoGetAPIRequest
func GetTaobaoDeActivityInfoGetAPIRequest() *TaobaoDeActivityInfoGetAPIRequest {
	return poolTaobaoDeActivityInfoGetAPIRequest.Get().(*TaobaoDeActivityInfoGetAPIRequest)
}

// ReleaseTaobaoDeActivityInfoGetAPIRequest 将 TaobaoDeActivityInfoGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoDeActivityInfoGetAPIRequest(v *TaobaoDeActivityInfoGetAPIRequest) {
	v.Reset()
	poolTaobaoDeActivityInfoGetAPIRequest.Put(v)
}
