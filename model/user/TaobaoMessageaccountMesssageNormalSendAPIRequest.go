package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaomessageaccountmesssagenormalsendAPIRequest 下行普通消息 API请求
// taobao.messageaccount.messsage.normal.send
//
// 消息号下行单个普通消息
type TaobaomessageaccountmesssagenormalsendAPIRequest struct {
	model.Params
	// 下行普通消息
	_param *NormalMessageDto
}

// NewTaobaomessageaccountmesssagenormalsendRequest 初始化TaobaomessageaccountmesssagenormalsendAPIRequest对象
func NewTaobaomessageaccountmesssagenormalsendRequest() *TaobaomessageaccountmesssagenormalsendAPIRequest {
	return &TaobaomessageaccountmesssagenormalsendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaomessageaccountmesssagenormalsendAPIRequest) GetApiMethodName() string {
	return "taobao.messageaccount.messsage.normal.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaomessageaccountmesssagenormalsendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaomessageaccountmesssagenormalsendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 下行普通消息
func (r *TaobaomessageaccountmesssagenormalsendAPIRequest) SetParam(_param *NormalMessageDto) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r TaobaomessageaccountmesssagenormalsendAPIRequest) GetParam() *NormalMessageDto {
	return r._param
}
