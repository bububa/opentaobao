package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaomessageaccountmesssagemasssendAPIRequest 消息号开放-消息群发 API请求
// taobao.messageaccount.messsage.mass.send
//
// 外部 isv 调用该进口来进行消息号消息的群发
type TaobaomessageaccountmesssagemasssendAPIRequest struct {
	model.Params
	// 参数
	_param *MassMessageDto
}

// NewTaobaomessageaccountmesssagemasssendRequest 初始化TaobaomessageaccountmesssagemasssendAPIRequest对象
func NewTaobaomessageaccountmesssagemasssendRequest() *TaobaomessageaccountmesssagemasssendAPIRequest {
	return &TaobaomessageaccountmesssagemasssendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaomessageaccountmesssagemasssendAPIRequest) GetApiMethodName() string {
	return "taobao.messageaccount.messsage.mass.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaomessageaccountmesssagemasssendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaomessageaccountmesssagemasssendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 参数
func (r *TaobaomessageaccountmesssagemasssendAPIRequest) SetParam(_param *MassMessageDto) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r TaobaomessageaccountmesssagemasssendAPIRequest) GetParam() *MassMessageDto {
	return r._param
}
