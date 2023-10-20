package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaomessageaccountmesssagereplyAPIRequest 消息号下行回复接口 API请求
// taobao.messageaccount.messsage.reply
//
// 外部 isv 调用该进口来进行消息号消息的回复
type TaobaomessageaccountmesssagereplyAPIRequest struct {
	model.Params
	// 入参
	_param0 *ReplyMessageDto
}

// NewTaobaomessageaccountmesssagereplyRequest 初始化TaobaomessageaccountmesssagereplyAPIRequest对象
func NewTaobaomessageaccountmesssagereplyRequest() *TaobaomessageaccountmesssagereplyAPIRequest {
	return &TaobaomessageaccountmesssagereplyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaomessageaccountmesssagereplyAPIRequest) GetApiMethodName() string {
	return "taobao.messageaccount.messsage.reply"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaomessageaccountmesssagereplyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaomessageaccountmesssagereplyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 入参
func (r *TaobaomessageaccountmesssagereplyAPIRequest) SetParam0(_param0 *ReplyMessageDto) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r TaobaomessageaccountmesssagereplyAPIRequest) GetParam0() *ReplyMessageDto {
	return r._param0
}
