package cmns

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunosservicecmnscoamessageresultgetAPIRequest CMNS消息发送到达查询 API请求
// yunos.service.cmns.coa.messageresult.get
//
// CMNS消息发送到达查询,根据消息ID查询，仅能查询该appKey所发送的消息
type YunosservicecmnscoamessageresultgetAPIRequest struct {
	model.Params
	// 消息ID
	_mid int64
}

// NewYunosservicecmnscoamessageresultgetRequest 初始化YunosservicecmnscoamessageresultgetAPIRequest对象
func NewYunosservicecmnscoamessageresultgetRequest() *YunosservicecmnscoamessageresultgetAPIRequest {
	return &YunosservicecmnscoamessageresultgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosservicecmnscoamessageresultgetAPIRequest) GetApiMethodName() string {
	return "yunos.service.cmns.coa.messageresult.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosservicecmnscoamessageresultgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosservicecmnscoamessageresultgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMid is Mid Setter
// 消息ID
func (r *YunosservicecmnscoamessageresultgetAPIRequest) SetMid(_mid int64) error {
	r._mid = _mid
	r.Set("mid", _mid)
	return nil
}

// GetMid Mid Getter
func (r YunosservicecmnscoamessageresultgetAPIRequest) GetMid() int64 {
	return r._mid
}
