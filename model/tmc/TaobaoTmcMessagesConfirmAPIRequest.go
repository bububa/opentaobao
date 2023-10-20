package tmc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotmcmessagesconfirmAPIRequest 确认消费消息的状态 API请求
// taobao.tmc.messages.confirm
//
// 确认消费消息的状态
type TaobaotmcmessagesconfirmAPIRequest struct {
	model.Params
	// 处理成功的消息ID列表 最大 200个ID
	_sMessageIds []string
	// 处理失败的消息ID列表--已废弃，无需传此字段
	_fMessageIds []string
	// 分组名称，不传代表默认分组
	_groupName string
}

// NewTaobaotmcmessagesconfirmRequest 初始化TaobaotmcmessagesconfirmAPIRequest对象
func NewTaobaotmcmessagesconfirmRequest() *TaobaotmcmessagesconfirmAPIRequest {
	return &TaobaotmcmessagesconfirmAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotmcmessagesconfirmAPIRequest) GetApiMethodName() string {
	return "taobao.tmc.messages.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotmcmessagesconfirmAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotmcmessagesconfirmAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSMessageIds is SMessageIds Setter
// 处理成功的消息ID列表 最大 200个ID
func (r *TaobaotmcmessagesconfirmAPIRequest) SetSMessageIds(_sMessageIds []string) error {
	r._sMessageIds = _sMessageIds
	r.Set("s_message_ids", _sMessageIds)
	return nil
}

// GetSMessageIds SMessageIds Getter
func (r TaobaotmcmessagesconfirmAPIRequest) GetSMessageIds() []string {
	return r._sMessageIds
}

// SetFMessageIds is FMessageIds Setter
// 处理失败的消息ID列表--已废弃，无需传此字段
func (r *TaobaotmcmessagesconfirmAPIRequest) SetFMessageIds(_fMessageIds []string) error {
	r._fMessageIds = _fMessageIds
	r.Set("f_message_ids", _fMessageIds)
	return nil
}

// GetFMessageIds FMessageIds Getter
func (r TaobaotmcmessagesconfirmAPIRequest) GetFMessageIds() []string {
	return r._fMessageIds
}

// SetGroupName is GroupName Setter
// 分组名称，不传代表默认分组
func (r *TaobaotmcmessagesconfirmAPIRequest) SetGroupName(_groupName string) error {
	r._groupName = _groupName
	r.Set("group_name", _groupName)
	return nil
}

// GetGroupName GroupName Getter
func (r TaobaotmcmessagesconfirmAPIRequest) GetGroupName() string {
	return r._groupName
}
