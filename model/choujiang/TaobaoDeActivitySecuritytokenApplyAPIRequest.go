package choujiang

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaodeactivitysecuritytokenapplyAPIRequest 安全token获取 API请求
// taobao.de.activity.securitytoken.apply
//
// 新增接口，这个接口是用于在手机端进行抽奖时候的验证使用
type TaobaodeactivitysecuritytokenapplyAPIRequest struct {
	model.Params
	// 运营和cp约定的事件唯一标示
	_eventKey string
}

// NewTaobaodeactivitysecuritytokenapplyRequest 初始化TaobaodeactivitysecuritytokenapplyAPIRequest对象
func NewTaobaodeactivitysecuritytokenapplyRequest() *TaobaodeactivitysecuritytokenapplyAPIRequest {
	return &TaobaodeactivitysecuritytokenapplyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaodeactivitysecuritytokenapplyAPIRequest) GetApiMethodName() string {
	return "taobao.de.activity.securitytoken.apply"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaodeactivitysecuritytokenapplyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaodeactivitysecuritytokenapplyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEventKey is EventKey Setter
// 运营和cp约定的事件唯一标示
func (r *TaobaodeactivitysecuritytokenapplyAPIRequest) SetEventKey(_eventKey string) error {
	r._eventKey = _eventKey
	r.Set("event_key", _eventKey)
	return nil
}

// GetEventKey EventKey Getter
func (r TaobaodeactivitysecuritytokenapplyAPIRequest) GetEventKey() string {
	return r._eventKey
}
