package choujiang

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoDeActivitySecuritytokenApplyAPIRequest 安全token获取 API请求
// taobao.de.activity.securitytoken.apply
//
// 新增接口，这个接口是用于在手机端进行抽奖时候的验证使用
type TaobaoDeActivitySecuritytokenApplyAPIRequest struct {
	model.Params
	// 运营和cp约定的事件唯一标示
	_eventKey string
}

// NewTaobaoDeActivitySecuritytokenApplyRequest 初始化TaobaoDeActivitySecuritytokenApplyAPIRequest对象
func NewTaobaoDeActivitySecuritytokenApplyRequest() *TaobaoDeActivitySecuritytokenApplyAPIRequest {
	return &TaobaoDeActivitySecuritytokenApplyAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoDeActivitySecuritytokenApplyAPIRequest) Reset() {
	r._eventKey = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoDeActivitySecuritytokenApplyAPIRequest) GetApiMethodName() string {
	return "taobao.de.activity.securitytoken.apply"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoDeActivitySecuritytokenApplyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoDeActivitySecuritytokenApplyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEventKey is EventKey Setter
// 运营和cp约定的事件唯一标示
func (r *TaobaoDeActivitySecuritytokenApplyAPIRequest) SetEventKey(_eventKey string) error {
	r._eventKey = _eventKey
	r.Set("event_key", _eventKey)
	return nil
}

// GetEventKey EventKey Getter
func (r TaobaoDeActivitySecuritytokenApplyAPIRequest) GetEventKey() string {
	return r._eventKey
}

var poolTaobaoDeActivitySecuritytokenApplyAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoDeActivitySecuritytokenApplyRequest()
	},
}

// GetTaobaoDeActivitySecuritytokenApplyRequest 从 sync.Pool 获取 TaobaoDeActivitySecuritytokenApplyAPIRequest
func GetTaobaoDeActivitySecuritytokenApplyAPIRequest() *TaobaoDeActivitySecuritytokenApplyAPIRequest {
	return poolTaobaoDeActivitySecuritytokenApplyAPIRequest.Get().(*TaobaoDeActivitySecuritytokenApplyAPIRequest)
}

// ReleaseTaobaoDeActivitySecuritytokenApplyAPIRequest 将 TaobaoDeActivitySecuritytokenApplyAPIRequest 放入 sync.Pool
func ReleaseTaobaoDeActivitySecuritytokenApplyAPIRequest(v *TaobaoDeActivitySecuritytokenApplyAPIRequest) {
	v.Reset()
	poolTaobaoDeActivitySecuritytokenApplyAPIRequest.Put(v)
}
