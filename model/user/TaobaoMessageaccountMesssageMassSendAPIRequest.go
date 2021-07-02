package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMessageaccountMesssageMassSendAPIRequest 消息号开放-消息群发 API请求
// taobao.messageaccount.messsage.mass.send
//
// 外部 isv 调用该进口来进行消息号消息的群发
type TaobaoMessageaccountMesssageMassSendAPIRequest struct {
	model.Params
	// 参数
	_param *MassMessageDto
}

// NewTaobaoMessageaccountMesssageMassSendRequest 初始化TaobaoMessageaccountMesssageMassSendAPIRequest对象
func NewTaobaoMessageaccountMesssageMassSendRequest() *TaobaoMessageaccountMesssageMassSendAPIRequest {
	return &TaobaoMessageaccountMesssageMassSendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoMessageaccountMesssageMassSendAPIRequest) GetApiMethodName() string {
	return "taobao.messageaccount.messsage.mass.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoMessageaccountMesssageMassSendAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Param Setter
// 参数
func (r *TaobaoMessageaccountMesssageMassSendAPIRequest) SetParam(_param *MassMessageDto) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// Get Param Getter
func (r TaobaoMessageaccountMesssageMassSendAPIRequest) GetParam() *MassMessageDto {
	return r._param
}
