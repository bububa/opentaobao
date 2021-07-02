package cmns

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunosServiceCmnsCoaMessageresultGetAPIRequest CMNS消息发送到达查询 API请求
// yunos.service.cmns.coa.messageresult.get
//
// CMNS消息发送到达查询,根据消息ID查询，仅能查询该appKey所发送的消息
type YunosServiceCmnsCoaMessageresultGetAPIRequest struct {
	model.Params
	// 消息ID
	_mid int64
}

// NewYunosServiceCmnsCoaMessageresultGetRequest 初始化YunosServiceCmnsCoaMessageresultGetAPIRequest对象
func NewYunosServiceCmnsCoaMessageresultGetRequest() *YunosServiceCmnsCoaMessageresultGetAPIRequest {
	return &YunosServiceCmnsCoaMessageresultGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosServiceCmnsCoaMessageresultGetAPIRequest) GetApiMethodName() string {
	return "yunos.service.cmns.coa.messageresult.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosServiceCmnsCoaMessageresultGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetMid is Mid Setter
// 消息ID
func (r *YunosServiceCmnsCoaMessageresultGetAPIRequest) SetMid(_mid int64) error {
	r._mid = _mid
	r.Set("mid", _mid)
	return nil
}

// GetMid Mid Getter
func (r YunosServiceCmnsCoaMessageresultGetAPIRequest) GetMid() int64 {
	return r._mid
}
