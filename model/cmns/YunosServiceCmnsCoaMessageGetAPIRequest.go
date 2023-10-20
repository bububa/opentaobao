package cmns

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunosservicecmnscoamessagegetAPIRequest 消息详情查询 API请求
// yunos.service.cmns.coa.message.get
//
// 第三方应用开发者调用此接口查询消息详情，只能查询此appKey发的消息
type YunosservicecmnscoamessagegetAPIRequest struct {
	model.Params
	// 消息id
	_mid int64
}

// NewYunosservicecmnscoamessagegetRequest 初始化YunosservicecmnscoamessagegetAPIRequest对象
func NewYunosservicecmnscoamessagegetRequest() *YunosservicecmnscoamessagegetAPIRequest {
	return &YunosservicecmnscoamessagegetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosservicecmnscoamessagegetAPIRequest) GetApiMethodName() string {
	return "yunos.service.cmns.coa.message.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosservicecmnscoamessagegetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosservicecmnscoamessagegetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMid is Mid Setter
// 消息id
func (r *YunosservicecmnscoamessagegetAPIRequest) SetMid(_mid int64) error {
	r._mid = _mid
	r.Set("mid", _mid)
	return nil
}

// GetMid Mid Getter
func (r YunosservicecmnscoamessagegetAPIRequest) GetMid() int64 {
	return r._mid
}
