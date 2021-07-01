package cmns

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosServiceCmnsCoaMessageGetAPIRequest
消息详情查询 API请求
yunos.service.cmns.coa.message.get

第三方应用开发者调用此接口查询消息详情，只能查询此appKey发的消息 */
type YunosServiceCmnsCoaMessageGetAPIRequest struct {
	model.Params
	// 消息id
	_mid int64
}

// NewYunosServiceCmnsCoaMessageGetRequest 初始化YunosServiceCmnsCoaMessageGetAPIRequest对象
func NewYunosServiceCmnsCoaMessageGetRequest() *YunosServiceCmnsCoaMessageGetAPIRequest {
	return &YunosServiceCmnsCoaMessageGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosServiceCmnsCoaMessageGetAPIRequest) GetApiMethodName() string {
	return "yunos.service.cmns.coa.message.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosServiceCmnsCoaMessageGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Mid Setter
// 消息id
func (r *YunosServiceCmnsCoaMessageGetAPIRequest) SetMid(_mid int64) error {
	r._mid = _mid
	r.Set("mid", _mid)
	return nil
}

// Get Mid Getter
func (r YunosServiceCmnsCoaMessageGetAPIRequest) GetMid() int64 {
	return r._mid
}
