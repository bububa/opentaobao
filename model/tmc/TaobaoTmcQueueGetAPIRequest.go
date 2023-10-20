package tmc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTmcQueueGetAPIRequest 获取消息队列积压情况 API请求
// taobao.tmc.queue.get
//
// 根据appkey和groupName获取消息队列积压情况
type TaobaoTmcQueueGetAPIRequest struct {
	model.Params
	// TMC组名
	_groupName string
}

// NewTaobaoTmcQueueGetRequest 初始化TaobaoTmcQueueGetAPIRequest对象
func NewTaobaoTmcQueueGetRequest() *TaobaoTmcQueueGetAPIRequest {
	return &TaobaoTmcQueueGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTmcQueueGetAPIRequest) Reset() {
	r._groupName = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTmcQueueGetAPIRequest) GetApiMethodName() string {
	return "taobao.tmc.queue.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTmcQueueGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTmcQueueGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetGroupName is GroupName Setter
// TMC组名
func (r *TaobaoTmcQueueGetAPIRequest) SetGroupName(_groupName string) error {
	r._groupName = _groupName
	r.Set("group_name", _groupName)
	return nil
}

// GetGroupName GroupName Getter
func (r TaobaoTmcQueueGetAPIRequest) GetGroupName() string {
	return r._groupName
}

var poolTaobaoTmcQueueGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTmcQueueGetRequest()
	},
}

// GetTaobaoTmcQueueGetRequest 从 sync.Pool 获取 TaobaoTmcQueueGetAPIRequest
func GetTaobaoTmcQueueGetAPIRequest() *TaobaoTmcQueueGetAPIRequest {
	return poolTaobaoTmcQueueGetAPIRequest.Get().(*TaobaoTmcQueueGetAPIRequest)
}

// ReleaseTaobaoTmcQueueGetAPIRequest 将 TaobaoTmcQueueGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoTmcQueueGetAPIRequest(v *TaobaoTmcQueueGetAPIRequest) {
	v.Reset()
	poolTaobaoTmcQueueGetAPIRequest.Put(v)
}
