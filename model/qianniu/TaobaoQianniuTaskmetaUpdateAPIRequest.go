package qianniu

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQianniuTaskmetaUpdateAPIRequest 更新任务元数据 API请求
// taobao.qianniu.taskmeta.update
//
// 由任务发起者调用
type TaobaoQianniuTaskmetaUpdateAPIRequest struct {
	model.Params
	// 要更新的任务元数据，JSON格式，例如：<br/>meta= {<br/>		"id" : 1,<br/>		"title" : "xxx",<br/>		"content" : "yyyy",<br/>		"biz_sys_Id" : 12232,<br/>		"biz_sys_task_type" : 1212,<br/>		"start_time" : 1380173565480,<br/>		"end_time" : 1380173565480,<br/> "sender_uid" : 213123213,<br/>		"sender_nick" : "tbtest1063",<br/>		"reminder_flag" : 1,<br/>		"finish_strategy" : 1<br/>	}
	_meta string
}

// NewTaobaoQianniuTaskmetaUpdateRequest 初始化TaobaoQianniuTaskmetaUpdateAPIRequest对象
func NewTaobaoQianniuTaskmetaUpdateRequest() *TaobaoQianniuTaskmetaUpdateAPIRequest {
	return &TaobaoQianniuTaskmetaUpdateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoQianniuTaskmetaUpdateAPIRequest) Reset() {
	r._meta = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQianniuTaskmetaUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.qianniu.taskmeta.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQianniuTaskmetaUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoQianniuTaskmetaUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMeta is Meta Setter
// 要更新的任务元数据，JSON格式，例如：&lt;br/&gt;meta= {&lt;br/&gt;		&#34;id&#34; : 1,&lt;br/&gt;		&#34;title&#34; : &#34;xxx&#34;,&lt;br/&gt;		&#34;content&#34; : &#34;yyyy&#34;,&lt;br/&gt;		&#34;biz_sys_Id&#34; : 12232,&lt;br/&gt;		&#34;biz_sys_task_type&#34; : 1212,&lt;br/&gt;		&#34;start_time&#34; : 1380173565480,&lt;br/&gt;		&#34;end_time&#34; : 1380173565480,&lt;br/&gt; &#34;sender_uid&#34; : 213123213,&lt;br/&gt;		&#34;sender_nick&#34; : &#34;tbtest1063&#34;,&lt;br/&gt;		&#34;reminder_flag&#34; : 1,&lt;br/&gt;		&#34;finish_strategy&#34; : 1&lt;br/&gt;	}
func (r *TaobaoQianniuTaskmetaUpdateAPIRequest) SetMeta(_meta string) error {
	r._meta = _meta
	r.Set("meta", _meta)
	return nil
}

// GetMeta Meta Getter
func (r TaobaoQianniuTaskmetaUpdateAPIRequest) GetMeta() string {
	return r._meta
}

var poolTaobaoQianniuTaskmetaUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoQianniuTaskmetaUpdateRequest()
	},
}

// GetTaobaoQianniuTaskmetaUpdateRequest 从 sync.Pool 获取 TaobaoQianniuTaskmetaUpdateAPIRequest
func GetTaobaoQianniuTaskmetaUpdateAPIRequest() *TaobaoQianniuTaskmetaUpdateAPIRequest {
	return poolTaobaoQianniuTaskmetaUpdateAPIRequest.Get().(*TaobaoQianniuTaskmetaUpdateAPIRequest)
}

// ReleaseTaobaoQianniuTaskmetaUpdateAPIRequest 将 TaobaoQianniuTaskmetaUpdateAPIRequest 放入 sync.Pool
func ReleaseTaobaoQianniuTaskmetaUpdateAPIRequest(v *TaobaoQianniuTaskmetaUpdateAPIRequest) {
	v.Reset()
	poolTaobaoQianniuTaskmetaUpdateAPIRequest.Put(v)
}
