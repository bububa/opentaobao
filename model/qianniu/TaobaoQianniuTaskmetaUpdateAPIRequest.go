package qianniu

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQianniuTaskmetaUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.qianniu.taskmeta.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQianniuTaskmetaUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Meta Setter
// 要更新的任务元数据，JSON格式，例如：<br/>meta= {<br/>		"id" : 1,<br/>		"title" : "xxx",<br/>		"content" : "yyyy",<br/>		"biz_sys_Id" : 12232,<br/>		"biz_sys_task_type" : 1212,<br/>		"start_time" : 1380173565480,<br/>		"end_time" : 1380173565480,<br/> "sender_uid" : 213123213,<br/>		"sender_nick" : "tbtest1063",<br/>		"reminder_flag" : 1,<br/>		"finish_strategy" : 1<br/>	}
func (r *TaobaoQianniuTaskmetaUpdateAPIRequest) SetMeta(_meta string) error {
	r._meta = _meta
	r.Set("meta", _meta)
	return nil
}

// Get Meta Getter
func (r TaobaoQianniuTaskmetaUpdateAPIRequest) GetMeta() string {
	return r._meta
}
