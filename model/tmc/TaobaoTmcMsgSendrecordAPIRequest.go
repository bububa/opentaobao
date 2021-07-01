package tmc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTmcMsgSendrecordAPIRequest
消息发送记录查询 API请求
taobao.tmc.msg.sendrecord

查询单条消息发送记录，只返回返回条数和时间。 */
type TaobaoTmcMsgSendrecordAPIRequest struct {
	model.Params
	// 消息分组名
	_groupName string
	// TOPIC名称
	_topicName string
	// 消息主键ID
	_dataId string
}

// NewTaobaoTmcMsgSendrecordRequest 初始化TaobaoTmcMsgSendrecordAPIRequest对象
func NewTaobaoTmcMsgSendrecordRequest() *TaobaoTmcMsgSendrecordAPIRequest {
	return &TaobaoTmcMsgSendrecordAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTmcMsgSendrecordAPIRequest) GetApiMethodName() string {
	return "taobao.tmc.msg.sendrecord"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTmcMsgSendrecordAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is GroupName Setter
// 消息分组名
func (r *TaobaoTmcMsgSendrecordAPIRequest) SetGroupName(_groupName string) error {
	r._groupName = _groupName
	r.Set("group_name", _groupName)
	return nil
}

// Get GroupName Getter
func (r TaobaoTmcMsgSendrecordAPIRequest) GetGroupName() string {
	return r._groupName
}

// Set is TopicName Setter
// TOPIC名称
func (r *TaobaoTmcMsgSendrecordAPIRequest) SetTopicName(_topicName string) error {
	r._topicName = _topicName
	r.Set("topic_name", _topicName)
	return nil
}

// Get TopicName Getter
func (r TaobaoTmcMsgSendrecordAPIRequest) GetTopicName() string {
	return r._topicName
}

// Set is DataId Setter
// 消息主键ID
func (r *TaobaoTmcMsgSendrecordAPIRequest) SetDataId(_dataId string) error {
	r._dataId = _dataId
	r.Set("data_id", _dataId)
	return nil
}

// Get DataId Getter
func (r TaobaoTmcMsgSendrecordAPIRequest) GetDataId() string {
	return r._dataId
}
