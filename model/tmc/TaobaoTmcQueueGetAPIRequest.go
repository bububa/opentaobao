package tmc

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTmcQueueGetAPIRequest) GetApiMethodName() string {
	return "taobao.tmc.queue.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTmcQueueGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is GroupName Setter
// TMC组名
func (r *TaobaoTmcQueueGetAPIRequest) SetGroupName(_groupName string) error {
	r._groupName = _groupName
	r.Set("group_name", _groupName)
	return nil
}

// Get GroupName Getter
func (r TaobaoTmcQueueGetAPIRequest) GetGroupName() string {
	return r._groupName
}
