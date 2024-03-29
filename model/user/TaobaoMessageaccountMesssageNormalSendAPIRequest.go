package user

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMessageaccountMesssageNormalSendAPIRequest 下行普通消息 API请求
// taobao.messageaccount.messsage.normal.send
//
// 消息号下行单个普通消息
type TaobaoMessageaccountMesssageNormalSendAPIRequest struct {
	model.Params
	// 下行普通消息
	_param *NormalMessageDto
}

// NewTaobaoMessageaccountMesssageNormalSendRequest 初始化TaobaoMessageaccountMesssageNormalSendAPIRequest对象
func NewTaobaoMessageaccountMesssageNormalSendRequest() *TaobaoMessageaccountMesssageNormalSendAPIRequest {
	return &TaobaoMessageaccountMesssageNormalSendAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoMessageaccountMesssageNormalSendAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoMessageaccountMesssageNormalSendAPIRequest) GetApiMethodName() string {
	return "taobao.messageaccount.messsage.normal.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoMessageaccountMesssageNormalSendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoMessageaccountMesssageNormalSendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 下行普通消息
func (r *TaobaoMessageaccountMesssageNormalSendAPIRequest) SetParam(_param *NormalMessageDto) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r TaobaoMessageaccountMesssageNormalSendAPIRequest) GetParam() *NormalMessageDto {
	return r._param
}

var poolTaobaoMessageaccountMesssageNormalSendAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoMessageaccountMesssageNormalSendRequest()
	},
}

// GetTaobaoMessageaccountMesssageNormalSendRequest 从 sync.Pool 获取 TaobaoMessageaccountMesssageNormalSendAPIRequest
func GetTaobaoMessageaccountMesssageNormalSendAPIRequest() *TaobaoMessageaccountMesssageNormalSendAPIRequest {
	return poolTaobaoMessageaccountMesssageNormalSendAPIRequest.Get().(*TaobaoMessageaccountMesssageNormalSendAPIRequest)
}

// ReleaseTaobaoMessageaccountMesssageNormalSendAPIRequest 将 TaobaoMessageaccountMesssageNormalSendAPIRequest 放入 sync.Pool
func ReleaseTaobaoMessageaccountMesssageNormalSendAPIRequest(v *TaobaoMessageaccountMesssageNormalSendAPIRequest) {
	v.Reset()
	poolTaobaoMessageaccountMesssageNormalSendAPIRequest.Put(v)
}
