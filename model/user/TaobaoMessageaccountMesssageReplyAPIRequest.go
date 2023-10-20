package user

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMessageaccountMesssageReplyAPIRequest 消息号下行回复接口 API请求
// taobao.messageaccount.messsage.reply
//
// 外部 isv 调用该进口来进行消息号消息的回复
type TaobaoMessageaccountMesssageReplyAPIRequest struct {
	model.Params
	// 入参
	_param0 *ReplyMessageDto
}

// NewTaobaoMessageaccountMesssageReplyRequest 初始化TaobaoMessageaccountMesssageReplyAPIRequest对象
func NewTaobaoMessageaccountMesssageReplyRequest() *TaobaoMessageaccountMesssageReplyAPIRequest {
	return &TaobaoMessageaccountMesssageReplyAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoMessageaccountMesssageReplyAPIRequest) Reset() {
	r._param0 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoMessageaccountMesssageReplyAPIRequest) GetApiMethodName() string {
	return "taobao.messageaccount.messsage.reply"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoMessageaccountMesssageReplyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoMessageaccountMesssageReplyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 入参
func (r *TaobaoMessageaccountMesssageReplyAPIRequest) SetParam0(_param0 *ReplyMessageDto) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r TaobaoMessageaccountMesssageReplyAPIRequest) GetParam0() *ReplyMessageDto {
	return r._param0
}

var poolTaobaoMessageaccountMesssageReplyAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoMessageaccountMesssageReplyRequest()
	},
}

// GetTaobaoMessageaccountMesssageReplyRequest 从 sync.Pool 获取 TaobaoMessageaccountMesssageReplyAPIRequest
func GetTaobaoMessageaccountMesssageReplyAPIRequest() *TaobaoMessageaccountMesssageReplyAPIRequest {
	return poolTaobaoMessageaccountMesssageReplyAPIRequest.Get().(*TaobaoMessageaccountMesssageReplyAPIRequest)
}

// ReleaseTaobaoMessageaccountMesssageReplyAPIRequest 将 TaobaoMessageaccountMesssageReplyAPIRequest 放入 sync.Pool
func ReleaseTaobaoMessageaccountMesssageReplyAPIRequest(v *TaobaoMessageaccountMesssageReplyAPIRequest) {
	v.Reset()
	poolTaobaoMessageaccountMesssageReplyAPIRequest.Put(v)
}
