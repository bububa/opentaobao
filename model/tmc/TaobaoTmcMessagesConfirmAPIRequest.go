package tmc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTmcMessagesConfirmAPIRequest 确认消费消息的状态 API请求
// taobao.tmc.messages.confirm
//
// 确认消费消息的状态
type TaobaoTmcMessagesConfirmAPIRequest struct {
	model.Params
	// 处理成功的消息ID列表 最大 200个ID
	_sMessageIds []string
	// 处理失败的消息ID列表--已废弃，无需传此字段
	_fMessageIds []string
	// 分组名称，不传代表默认分组
	_groupName string
}

// NewTaobaoTmcMessagesConfirmRequest 初始化TaobaoTmcMessagesConfirmAPIRequest对象
func NewTaobaoTmcMessagesConfirmRequest() *TaobaoTmcMessagesConfirmAPIRequest {
	return &TaobaoTmcMessagesConfirmAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTmcMessagesConfirmAPIRequest) Reset() {
	r._sMessageIds = r._sMessageIds[:0]
	r._fMessageIds = r._fMessageIds[:0]
	r._groupName = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTmcMessagesConfirmAPIRequest) GetApiMethodName() string {
	return "taobao.tmc.messages.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTmcMessagesConfirmAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTmcMessagesConfirmAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSMessageIds is SMessageIds Setter
// 处理成功的消息ID列表 最大 200个ID
func (r *TaobaoTmcMessagesConfirmAPIRequest) SetSMessageIds(_sMessageIds []string) error {
	r._sMessageIds = _sMessageIds
	r.Set("s_message_ids", _sMessageIds)
	return nil
}

// GetSMessageIds SMessageIds Getter
func (r TaobaoTmcMessagesConfirmAPIRequest) GetSMessageIds() []string {
	return r._sMessageIds
}

// SetFMessageIds is FMessageIds Setter
// 处理失败的消息ID列表--已废弃，无需传此字段
func (r *TaobaoTmcMessagesConfirmAPIRequest) SetFMessageIds(_fMessageIds []string) error {
	r._fMessageIds = _fMessageIds
	r.Set("f_message_ids", _fMessageIds)
	return nil
}

// GetFMessageIds FMessageIds Getter
func (r TaobaoTmcMessagesConfirmAPIRequest) GetFMessageIds() []string {
	return r._fMessageIds
}

// SetGroupName is GroupName Setter
// 分组名称，不传代表默认分组
func (r *TaobaoTmcMessagesConfirmAPIRequest) SetGroupName(_groupName string) error {
	r._groupName = _groupName
	r.Set("group_name", _groupName)
	return nil
}

// GetGroupName GroupName Getter
func (r TaobaoTmcMessagesConfirmAPIRequest) GetGroupName() string {
	return r._groupName
}

var poolTaobaoTmcMessagesConfirmAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTmcMessagesConfirmRequest()
	},
}

// GetTaobaoTmcMessagesConfirmRequest 从 sync.Pool 获取 TaobaoTmcMessagesConfirmAPIRequest
func GetTaobaoTmcMessagesConfirmAPIRequest() *TaobaoTmcMessagesConfirmAPIRequest {
	return poolTaobaoTmcMessagesConfirmAPIRequest.Get().(*TaobaoTmcMessagesConfirmAPIRequest)
}

// ReleaseTaobaoTmcMessagesConfirmAPIRequest 将 TaobaoTmcMessagesConfirmAPIRequest 放入 sync.Pool
func ReleaseTaobaoTmcMessagesConfirmAPIRequest(v *TaobaoTmcMessagesConfirmAPIRequest) {
	v.Reset()
	poolTaobaoTmcMessagesConfirmAPIRequest.Put(v)
}
