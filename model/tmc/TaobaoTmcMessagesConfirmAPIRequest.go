package tmc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTmcMessagesConfirmAPIRequest 确认消费消息的状态 API请求
// taobao.tmc.messages.confirm
//
// 确认消费消息的状态
type TaobaoTmcMessagesConfirmAPIRequest struct {
	model.Params
	// 处理成功的消息ID列表 最大 200个ID
	_sMessageIds []int64
	// 处理失败的消息ID列表--已废弃，无需传此字段
	_fMessageIds []int64
	// 分组名称，不传代表默认分组
	_groupName string
}

// NewTaobaoTmcMessagesConfirmRequest 初始化TaobaoTmcMessagesConfirmAPIRequest对象
func NewTaobaoTmcMessagesConfirmRequest() *TaobaoTmcMessagesConfirmAPIRequest {
	return &TaobaoTmcMessagesConfirmAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTmcMessagesConfirmAPIRequest) GetApiMethodName() string {
	return "taobao.tmc.messages.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTmcMessagesConfirmAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetSMessageIds is SMessageIds Setter
// 处理成功的消息ID列表 最大 200个ID
func (r *TaobaoTmcMessagesConfirmAPIRequest) SetSMessageIds(_sMessageIds []int64) error {
	r._sMessageIds = _sMessageIds
	r.Set("s_message_ids", _sMessageIds)
	return nil
}

// GetSMessageIds SMessageIds Getter
func (r TaobaoTmcMessagesConfirmAPIRequest) GetSMessageIds() []int64 {
	return r._sMessageIds
}

// SetFMessageIds is FMessageIds Setter
// 处理失败的消息ID列表--已废弃，无需传此字段
func (r *TaobaoTmcMessagesConfirmAPIRequest) SetFMessageIds(_fMessageIds []int64) error {
	r._fMessageIds = _fMessageIds
	r.Set("f_message_ids", _fMessageIds)
	return nil
}

// GetFMessageIds FMessageIds Getter
func (r TaobaoTmcMessagesConfirmAPIRequest) GetFMessageIds() []int64 {
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
