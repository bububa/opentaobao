package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMiniappMesssageNormalSendAPIRequest 轻店铺下行普通消息给用户 API请求
// taobao.miniapp.messsage.normal.send
//
// 小程序下行单个普通消息
type TaobaoMiniappMesssageNormalSendAPIRequest struct {
	model.Params
	// 普通消息结构
	_param *DownNormalMessageDto
}

// NewTaobaoMiniappMesssageNormalSendRequest 初始化TaobaoMiniappMesssageNormalSendAPIRequest对象
func NewTaobaoMiniappMesssageNormalSendRequest() *TaobaoMiniappMesssageNormalSendAPIRequest {
	return &TaobaoMiniappMesssageNormalSendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoMiniappMesssageNormalSendAPIRequest) GetApiMethodName() string {
	return "taobao.miniapp.messsage.normal.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoMiniappMesssageNormalSendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoMiniappMesssageNormalSendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 普通消息结构
func (r *TaobaoMiniappMesssageNormalSendAPIRequest) SetParam(_param *DownNormalMessageDto) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r TaobaoMiniappMesssageNormalSendAPIRequest) GetParam() *DownNormalMessageDto {
	return r._param
}
