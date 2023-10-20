package user

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMiniappMesssageReplyAPIRequest 轻店铺下行回复接口 API请求
// taobao.miniapp.messsage.reply
//
// 外部 isv 调用该进口来进行轻店铺消息的回复
type TaobaoMiniappMesssageReplyAPIRequest struct {
	model.Params
	// 入参
	_param *ReplyMessageDto
}

// NewTaobaoMiniappMesssageReplyRequest 初始化TaobaoMiniappMesssageReplyAPIRequest对象
func NewTaobaoMiniappMesssageReplyRequest() *TaobaoMiniappMesssageReplyAPIRequest {
	return &TaobaoMiniappMesssageReplyAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoMiniappMesssageReplyAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoMiniappMesssageReplyAPIRequest) GetApiMethodName() string {
	return "taobao.miniapp.messsage.reply"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoMiniappMesssageReplyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoMiniappMesssageReplyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 入参
func (r *TaobaoMiniappMesssageReplyAPIRequest) SetParam(_param *ReplyMessageDto) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r TaobaoMiniappMesssageReplyAPIRequest) GetParam() *ReplyMessageDto {
	return r._param
}

var poolTaobaoMiniappMesssageReplyAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoMiniappMesssageReplyRequest()
	},
}

// GetTaobaoMiniappMesssageReplyRequest 从 sync.Pool 获取 TaobaoMiniappMesssageReplyAPIRequest
func GetTaobaoMiniappMesssageReplyAPIRequest() *TaobaoMiniappMesssageReplyAPIRequest {
	return poolTaobaoMiniappMesssageReplyAPIRequest.Get().(*TaobaoMiniappMesssageReplyAPIRequest)
}

// ReleaseTaobaoMiniappMesssageReplyAPIRequest 将 TaobaoMiniappMesssageReplyAPIRequest 放入 sync.Pool
func ReleaseTaobaoMiniappMesssageReplyAPIRequest(v *TaobaoMiniappMesssageReplyAPIRequest) {
	v.Reset()
	poolTaobaoMiniappMesssageReplyAPIRequest.Put(v)
}
