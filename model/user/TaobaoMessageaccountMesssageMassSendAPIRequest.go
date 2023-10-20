package user

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoMessageaccountMesssageMassSendAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoMessageaccountMesssageMassSendAPIRequest) GetApiMethodName() string {
	return "taobao.messageaccount.messsage.mass.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoMessageaccountMesssageMassSendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoMessageaccountMesssageMassSendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 参数
func (r *TaobaoMessageaccountMesssageMassSendAPIRequest) SetParam(_param *MassMessageDto) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r TaobaoMessageaccountMesssageMassSendAPIRequest) GetParam() *MassMessageDto {
	return r._param
}

var poolTaobaoMessageaccountMesssageMassSendAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoMessageaccountMesssageMassSendRequest()
	},
}

// GetTaobaoMessageaccountMesssageMassSendRequest 从 sync.Pool 获取 TaobaoMessageaccountMesssageMassSendAPIRequest
func GetTaobaoMessageaccountMesssageMassSendAPIRequest() *TaobaoMessageaccountMesssageMassSendAPIRequest {
	return poolTaobaoMessageaccountMesssageMassSendAPIRequest.Get().(*TaobaoMessageaccountMesssageMassSendAPIRequest)
}

// ReleaseTaobaoMessageaccountMesssageMassSendAPIRequest 将 TaobaoMessageaccountMesssageMassSendAPIRequest 放入 sync.Pool
func ReleaseTaobaoMessageaccountMesssageMassSendAPIRequest(v *TaobaoMessageaccountMesssageMassSendAPIRequest) {
	v.Reset()
	poolTaobaoMessageaccountMesssageMassSendAPIRequest.Put(v)
}
