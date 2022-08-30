package aliqin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAliqinFcVoiceNumDoublecallAPIRequest 多方通话 API请求
// alibaba.aliqin.fc.voice.num.doublecall
//
// 根据传入的主叫号码与被叫号码，由系统依次向主叫号码与被叫号码发起呼叫，双方都应答后，开始一对一通话并开始计费。使用前需要在阿里大于管理中心添加呼叫双方的显示号码。
type AlibabaAliqinFcVoiceNumDoublecallAPIRequest struct {
	model.Params
	// 主叫号码，支持国内手机号与固话号码,格式如下057188773344,13911112222,4001112222,95500
	_callerNum string
	// 主叫号码侧的号码显示，传入的显示号码必须是阿里大于“管理中心-号码管理”中申请通过的号码。显示号码格式如下057188773344，4001112222，95500
	_callerShowNum string
	// 被叫号码，支持国内手机号与固话号码,格式如下057188773344,13911112222,4001112222,95500
	_calledNum string
	// 被叫号码侧的号码显示，传入的显示号码可以是阿里大于“管理中心-号码管理”中申请通过的号码。显示号码格式如下057188773344，4001112222，95500。显示号码也可以为主叫号码。
	_calledShowNum string
	// 公共回传参数，在“消息返回”中会透传回该参数；举例：用户可以传入自己下级的会员ID，在消息返回时，该会员ID会包含在内，用户可以根据该会员ID识别是哪位会员使用了你的应用
	_extend string
	// 通话超时时长，如接通后到达120秒时，通话会因为超时自动挂断。若无需设置超时时长，可不传。
	_sessionTimeOut string
}

// NewAlibabaAliqinFcVoiceNumDoublecallRequest 初始化AlibabaAliqinFcVoiceNumDoublecallAPIRequest对象
func NewAlibabaAliqinFcVoiceNumDoublecallRequest() *AlibabaAliqinFcVoiceNumDoublecallAPIRequest {
	return &AlibabaAliqinFcVoiceNumDoublecallAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAliqinFcVoiceNumDoublecallAPIRequest) GetApiMethodName() string {
	return "alibaba.aliqin.fc.voice.num.doublecall"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAliqinFcVoiceNumDoublecallAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetCallerNum is CallerNum Setter
// 主叫号码，支持国内手机号与固话号码,格式如下057188773344,13911112222,4001112222,95500
func (r *AlibabaAliqinFcVoiceNumDoublecallAPIRequest) SetCallerNum(_callerNum string) error {
	r._callerNum = _callerNum
	r.Set("caller_num", _callerNum)
	return nil
}

// GetCallerNum CallerNum Getter
func (r AlibabaAliqinFcVoiceNumDoublecallAPIRequest) GetCallerNum() string {
	return r._callerNum
}

// SetCallerShowNum is CallerShowNum Setter
// 主叫号码侧的号码显示，传入的显示号码必须是阿里大于“管理中心-号码管理”中申请通过的号码。显示号码格式如下057188773344，4001112222，95500
func (r *AlibabaAliqinFcVoiceNumDoublecallAPIRequest) SetCallerShowNum(_callerShowNum string) error {
	r._callerShowNum = _callerShowNum
	r.Set("caller_show_num", _callerShowNum)
	return nil
}

// GetCallerShowNum CallerShowNum Getter
func (r AlibabaAliqinFcVoiceNumDoublecallAPIRequest) GetCallerShowNum() string {
	return r._callerShowNum
}

// SetCalledNum is CalledNum Setter
// 被叫号码，支持国内手机号与固话号码,格式如下057188773344,13911112222,4001112222,95500
func (r *AlibabaAliqinFcVoiceNumDoublecallAPIRequest) SetCalledNum(_calledNum string) error {
	r._calledNum = _calledNum
	r.Set("called_num", _calledNum)
	return nil
}

// GetCalledNum CalledNum Getter
func (r AlibabaAliqinFcVoiceNumDoublecallAPIRequest) GetCalledNum() string {
	return r._calledNum
}

// SetCalledShowNum is CalledShowNum Setter
// 被叫号码侧的号码显示，传入的显示号码可以是阿里大于“管理中心-号码管理”中申请通过的号码。显示号码格式如下057188773344，4001112222，95500。显示号码也可以为主叫号码。
func (r *AlibabaAliqinFcVoiceNumDoublecallAPIRequest) SetCalledShowNum(_calledShowNum string) error {
	r._calledShowNum = _calledShowNum
	r.Set("called_show_num", _calledShowNum)
	return nil
}

// GetCalledShowNum CalledShowNum Getter
func (r AlibabaAliqinFcVoiceNumDoublecallAPIRequest) GetCalledShowNum() string {
	return r._calledShowNum
}

// SetExtend is Extend Setter
// 公共回传参数，在“消息返回”中会透传回该参数；举例：用户可以传入自己下级的会员ID，在消息返回时，该会员ID会包含在内，用户可以根据该会员ID识别是哪位会员使用了你的应用
func (r *AlibabaAliqinFcVoiceNumDoublecallAPIRequest) SetExtend(_extend string) error {
	r._extend = _extend
	r.Set("extend", _extend)
	return nil
}

// GetExtend Extend Getter
func (r AlibabaAliqinFcVoiceNumDoublecallAPIRequest) GetExtend() string {
	return r._extend
}

// SetSessionTimeOut is SessionTimeOut Setter
// 通话超时时长，如接通后到达120秒时，通话会因为超时自动挂断。若无需设置超时时长，可不传。
func (r *AlibabaAliqinFcVoiceNumDoublecallAPIRequest) SetSessionTimeOut(_sessionTimeOut string) error {
	r._sessionTimeOut = _sessionTimeOut
	r.Set("session_time_out", _sessionTimeOut)
	return nil
}

// GetSessionTimeOut SessionTimeOut Getter
func (r AlibabaAliqinFcVoiceNumDoublecallAPIRequest) GetSessionTimeOut() string {
	return r._sessionTimeOut
}
