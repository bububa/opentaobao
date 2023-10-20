package tmc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLsyMiniappMsgPushAPIRequest 零售云小程序消息推送 API请求
// alibaba.lsy.miniapp.msg.push
//
// 零售云小程序消息推送，推送消息至零售云（喵零等）
type AlibabaLsyMiniappMsgPushAPIRequest struct {
	model.Params
	// 小程序ID
	_appId string
	// 消息模板，miaoling_msg_isv_clue - 线索通知消息
	_templateId string
	// 消息参数
	_params string
	// 消息ID
	_msgId int64
	// 摊位ID
	_storeId int64
}

// NewAlibabaLsyMiniappMsgPushRequest 初始化AlibabaLsyMiniappMsgPushAPIRequest对象
func NewAlibabaLsyMiniappMsgPushRequest() *AlibabaLsyMiniappMsgPushAPIRequest {
	return &AlibabaLsyMiniappMsgPushAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLsyMiniappMsgPushAPIRequest) GetApiMethodName() string {
	return "alibaba.lsy.miniapp.msg.push"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLsyMiniappMsgPushAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaLsyMiniappMsgPushAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAppId is AppId Setter
// 小程序ID
func (r *AlibabaLsyMiniappMsgPushAPIRequest) SetAppId(_appId string) error {
	r._appId = _appId
	r.Set("app_id", _appId)
	return nil
}

// GetAppId AppId Getter
func (r AlibabaLsyMiniappMsgPushAPIRequest) GetAppId() string {
	return r._appId
}

// SetTemplateId is TemplateId Setter
// 消息模板，miaoling_msg_isv_clue - 线索通知消息
func (r *AlibabaLsyMiniappMsgPushAPIRequest) SetTemplateId(_templateId string) error {
	r._templateId = _templateId
	r.Set("template_id", _templateId)
	return nil
}

// GetTemplateId TemplateId Getter
func (r AlibabaLsyMiniappMsgPushAPIRequest) GetTemplateId() string {
	return r._templateId
}

// SetParams is Params Setter
// 消息参数
func (r *AlibabaLsyMiniappMsgPushAPIRequest) SetParams(_params string) error {
	r._params = _params
	r.Set("params", _params)
	return nil
}

// GetParams Params Getter
func (r AlibabaLsyMiniappMsgPushAPIRequest) GetParams() string {
	return r._params
}

// SetMsgId is MsgId Setter
// 消息ID
func (r *AlibabaLsyMiniappMsgPushAPIRequest) SetMsgId(_msgId int64) error {
	r._msgId = _msgId
	r.Set("msg_id", _msgId)
	return nil
}

// GetMsgId MsgId Getter
func (r AlibabaLsyMiniappMsgPushAPIRequest) GetMsgId() int64 {
	return r._msgId
}

// SetStoreId is StoreId Setter
// 摊位ID
func (r *AlibabaLsyMiniappMsgPushAPIRequest) SetStoreId(_storeId int64) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r AlibabaLsyMiniappMsgPushAPIRequest) GetStoreId() int64 {
	return r._storeId
}
