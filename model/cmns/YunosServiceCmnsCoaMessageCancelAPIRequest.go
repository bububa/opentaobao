package cmns

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosServiceCmnsCoaMessageCancelAPIRequest CMNS消息撤回 API请求
// yunos.service.cmns.coa.message.cancel
//
// 此接口用户撤回之前已经发出去的消息，根据消息ID撤回，只能撤回此appKey创建的消息。
type YunosServiceCmnsCoaMessageCancelAPIRequest struct {
	model.Params
	// 消息ID
	_mid int64
}

// NewYunosServiceCmnsCoaMessageCancelRequest 初始化YunosServiceCmnsCoaMessageCancelAPIRequest对象
func NewYunosServiceCmnsCoaMessageCancelRequest() *YunosServiceCmnsCoaMessageCancelAPIRequest {
	return &YunosServiceCmnsCoaMessageCancelAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YunosServiceCmnsCoaMessageCancelAPIRequest) Reset() {
	r._mid = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosServiceCmnsCoaMessageCancelAPIRequest) GetApiMethodName() string {
	return "yunos.service.cmns.coa.message.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosServiceCmnsCoaMessageCancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosServiceCmnsCoaMessageCancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMid is Mid Setter
// 消息ID
func (r *YunosServiceCmnsCoaMessageCancelAPIRequest) SetMid(_mid int64) error {
	r._mid = _mid
	r.Set("mid", _mid)
	return nil
}

// GetMid Mid Getter
func (r YunosServiceCmnsCoaMessageCancelAPIRequest) GetMid() int64 {
	return r._mid
}

var poolYunosServiceCmnsCoaMessageCancelAPIRequest = sync.Pool{
	New: func() any {
		return NewYunosServiceCmnsCoaMessageCancelRequest()
	},
}

// GetYunosServiceCmnsCoaMessageCancelRequest 从 sync.Pool 获取 YunosServiceCmnsCoaMessageCancelAPIRequest
func GetYunosServiceCmnsCoaMessageCancelAPIRequest() *YunosServiceCmnsCoaMessageCancelAPIRequest {
	return poolYunosServiceCmnsCoaMessageCancelAPIRequest.Get().(*YunosServiceCmnsCoaMessageCancelAPIRequest)
}

// ReleaseYunosServiceCmnsCoaMessageCancelAPIRequest 将 YunosServiceCmnsCoaMessageCancelAPIRequest 放入 sync.Pool
func ReleaseYunosServiceCmnsCoaMessageCancelAPIRequest(v *YunosServiceCmnsCoaMessageCancelAPIRequest) {
	v.Reset()
	poolYunosServiceCmnsCoaMessageCancelAPIRequest.Put(v)
}
