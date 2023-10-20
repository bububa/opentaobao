package tmc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotmcqueuegetAPIRequest 获取消息队列积压情况 API请求
// taobao.tmc.queue.get
//
// 根据appkey和groupName获取消息队列积压情况
type TaobaotmcqueuegetAPIRequest struct {
	model.Params
	// TMC组名
	_groupName string
}

// NewTaobaotmcqueuegetRequest 初始化TaobaotmcqueuegetAPIRequest对象
func NewTaobaotmcqueuegetRequest() *TaobaotmcqueuegetAPIRequest {
	return &TaobaotmcqueuegetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotmcqueuegetAPIRequest) GetApiMethodName() string {
	return "taobao.tmc.queue.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotmcqueuegetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotmcqueuegetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetGroupName is GroupName Setter
// TMC组名
func (r *TaobaotmcqueuegetAPIRequest) SetGroupName(_groupName string) error {
	r._groupName = _groupName
	r.Set("group_name", _groupName)
	return nil
}

// GetGroupName GroupName Getter
func (r TaobaotmcqueuegetAPIRequest) GetGroupName() string {
	return r._groupName
}
