package cmns

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunosservicecmnscoamessagecancelAPIRequest CMNS消息撤回 API请求
// yunos.service.cmns.coa.message.cancel
//
// 此接口用户撤回之前已经发出去的消息，根据消息ID撤回，只能撤回此appKey创建的消息。
type YunosservicecmnscoamessagecancelAPIRequest struct {
	model.Params
	// 消息ID
	_mid int64
}

// NewYunosservicecmnscoamessagecancelRequest 初始化YunosservicecmnscoamessagecancelAPIRequest对象
func NewYunosservicecmnscoamessagecancelRequest() *YunosservicecmnscoamessagecancelAPIRequest {
	return &YunosservicecmnscoamessagecancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosservicecmnscoamessagecancelAPIRequest) GetApiMethodName() string {
	return "yunos.service.cmns.coa.message.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosservicecmnscoamessagecancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosservicecmnscoamessagecancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMid is Mid Setter
// 消息ID
func (r *YunosservicecmnscoamessagecancelAPIRequest) SetMid(_mid int64) error {
	r._mid = _mid
	r.Set("mid", _mid)
	return nil
}

// GetMid Mid Getter
func (r YunosservicecmnscoamessagecancelAPIRequest) GetMid() int64 {
	return r._mid
}
