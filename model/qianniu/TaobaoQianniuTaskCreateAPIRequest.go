package qianniu

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoQianniuTaskCreateAPIRequest) Reset() {
	r._meta = ""
	r._tasks = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQianniuTaskCreateAPIRequest) GetApiMethodName() string {
	return "taobao.qianniu.task.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQianniuTaskCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoQianniuTaskCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMeta is Meta Setter
// 任务元数据，JSON格式，例如：&lt;br/&gt;meta = {&lt;br/&gt;            title : &#34;可自定义&#34;,&lt;br/&gt;            content : “任务正文”,&lt;br/&gt;            sender_uid : user_id,&lt;br/&gt;            sender_nick : user_nick,&lt;br/&gt;            reminder_flag : 1,&lt;br/&gt;            finish_strategy : 0,&lt;br/&gt;         biz_type : &#34;memo&#34;,&lt;br/&gt;         priority : 0&lt;br/&gt;        };&lt;br/&gt;说明：reminder_flag:1表示需要发送任务提醒消息,0表示不需要消息提醒。建议写1;&lt;br/&gt;finish_strategy : 0表示只要一个人完成任务就可以，1表示所有人都需要完成任务。根据场景设置，建议选0;&lt;br/&gt;biz_type : 任务类型，请咨询千牛官方获取正确的任务类型;&lt;br/&gt;priority : 1表示高优先级，0表示普通;&lt;br/&gt;这里的举例为必填字段，一些选填字段没有列出，如有其它需求请联系千牛官方。
func (r *TaobaoQianniuTaskCreateAPIRequest) SetMeta(_meta string) error {
	r._meta = _meta
	r.Set("meta", _meta)
	return nil
}

// GetMeta Meta Getter
func (r TaobaoQianniuTaskCreateAPIRequest) GetMeta() string {
	return r._meta
}

// SetTasks is Tasks Setter
// 任务列表，JSON格式的数组，即支持多个接收人，例如：&lt;br/&gt;task = [{&lt;br/&gt;            receiver_uid : user_id,&lt;br/&gt;            receiver_nick : user_nick,&lt;br/&gt;            biz_type : &#34;memo&#34;,&lt;br/&gt;            sub_biz_type : &#34;memo&#34;,&lt;br/&gt;            biz_id : user_nick,&lt;br/&gt;            biz_nick : user_nick&lt;br/&gt;         }];&lt;br/&gt;上述为必填字段，其它字段请咨询千牛官方。
func (r *TaobaoQianniuTaskCreateAPIRequest) SetTasks(_tasks string) error {
	r._tasks = _tasks
	r.Set("tasks", _tasks)
	return nil
}

// GetTasks Tasks Getter
func (r TaobaoQianniuTaskCreateAPIRequest) GetTasks() string {
	return r._tasks
}

var poolTaobaoQianniuTaskCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoQianniuTaskCreateRequest()
	},
}

// GetTaobaoQianniuTaskCreateRequest 从 sync.Pool 获取 TaobaoQianniuTaskCreateAPIRequest
func GetTaobaoQianniuTaskCreateAPIRequest() *TaobaoQianniuTaskCreateAPIRequest {
	return poolTaobaoQianniuTaskCreateAPIRequest.Get().(*TaobaoQianniuTaskCreateAPIRequest)
}

// ReleaseTaobaoQianniuTaskCreateAPIRequest 将 TaobaoQianniuTaskCreateAPIRequest 放入 sync.Pool
func ReleaseTaobaoQianniuTaskCreateAPIRequest(v *TaobaoQianniuTaskCreateAPIRequest) {
	v.Reset()
	poolTaobaoQianniuTaskCreateAPIRequest.Put(v)
}
