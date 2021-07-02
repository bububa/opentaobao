package qianniu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQianniuTaskCreateAPIRequest 创建轻任务 API请求
// taobao.qianniu.task.create
//
// 发起一个轻任务，分配给多个执行者，并发送消息提醒，由任务发起者调用
type TaobaoQianniuTaskCreateAPIRequest struct {
	model.Params
	// 任务元数据，JSON格式，例如：<br/>meta = {<br/>            title : "可自定义",<br/>            content : “任务正文”,<br/>            sender_uid : user_id,<br/>            sender_nick : user_nick,<br/>            reminder_flag : 1,<br/>            finish_strategy : 0,<br/>         biz_type : "memo",<br/>         priority : 0<br/>        };<br/>说明：reminder_flag:1表示需要发送任务提醒消息,0表示不需要消息提醒。建议写1;<br/>finish_strategy : 0表示只要一个人完成任务就可以，1表示所有人都需要完成任务。根据场景设置，建议选0;<br/>biz_type : 任务类型，请咨询千牛官方获取正确的任务类型;<br/>priority : 1表示高优先级，0表示普通;<br/>这里的举例为必填字段，一些选填字段没有列出，如有其它需求请联系千牛官方。
	_meta string
	// 任务列表，JSON格式的数组，即支持多个接收人，例如：<br/>task = [{<br/>            receiver_uid : user_id,<br/>            receiver_nick : user_nick,<br/>            biz_type : "memo",<br/>            sub_biz_type : "memo",<br/>            biz_id : user_nick,<br/>            biz_nick : user_nick<br/>         }];<br/>上述为必填字段，其它字段请咨询千牛官方。
	_tasks string
}

// NewTaobaoQianniuTaskCreateRequest 初始化TaobaoQianniuTaskCreateAPIRequest对象
func NewTaobaoQianniuTaskCreateRequest() *TaobaoQianniuTaskCreateAPIRequest {
	return &TaobaoQianniuTaskCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQianniuTaskCreateAPIRequest) GetApiMethodName() string {
	return "taobao.qianniu.task.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQianniuTaskCreateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Meta Setter
// 任务元数据，JSON格式，例如：<br/>meta = {<br/>            title : "可自定义",<br/>            content : “任务正文”,<br/>            sender_uid : user_id,<br/>            sender_nick : user_nick,<br/>            reminder_flag : 1,<br/>            finish_strategy : 0,<br/>         biz_type : "memo",<br/>         priority : 0<br/>        };<br/>说明：reminder_flag:1表示需要发送任务提醒消息,0表示不需要消息提醒。建议写1;<br/>finish_strategy : 0表示只要一个人完成任务就可以，1表示所有人都需要完成任务。根据场景设置，建议选0;<br/>biz_type : 任务类型，请咨询千牛官方获取正确的任务类型;<br/>priority : 1表示高优先级，0表示普通;<br/>这里的举例为必填字段，一些选填字段没有列出，如有其它需求请联系千牛官方。
func (r *TaobaoQianniuTaskCreateAPIRequest) SetMeta(_meta string) error {
	r._meta = _meta
	r.Set("meta", _meta)
	return nil
}

// Get Meta Getter
func (r TaobaoQianniuTaskCreateAPIRequest) GetMeta() string {
	return r._meta
}

// Set is Tasks Setter
// 任务列表，JSON格式的数组，即支持多个接收人，例如：<br/>task = [{<br/>            receiver_uid : user_id,<br/>            receiver_nick : user_nick,<br/>            biz_type : "memo",<br/>            sub_biz_type : "memo",<br/>            biz_id : user_nick,<br/>            biz_nick : user_nick<br/>         }];<br/>上述为必填字段，其它字段请咨询千牛官方。
func (r *TaobaoQianniuTaskCreateAPIRequest) SetTasks(_tasks string) error {
	r._tasks = _tasks
	r.Set("tasks", _tasks)
	return nil
}

// Get Tasks Getter
func (r TaobaoQianniuTaskCreateAPIRequest) GetTasks() string {
	return r._tasks
}
