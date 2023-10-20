package degoperation

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaodegoperationcheckaddrstatusAPIRequest 地址 API请求
// taobao.degoperation.check.addr.status
//
// 激励
type TaobaodegoperationcheckaddrstatusAPIRequest struct {
	model.Params
	// 奖品唯一标识
	_sequenceNo int64
}

// NewTaobaodegoperationcheckaddrstatusRequest 初始化TaobaodegoperationcheckaddrstatusAPIRequest对象
func NewTaobaodegoperationcheckaddrstatusRequest() *TaobaodegoperationcheckaddrstatusAPIRequest {
	return &TaobaodegoperationcheckaddrstatusAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaodegoperationcheckaddrstatusAPIRequest) GetApiMethodName() string {
	return "taobao.degoperation.check.addr.status"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaodegoperationcheckaddrstatusAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaodegoperationcheckaddrstatusAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSequenceNo is SequenceNo Setter
// 奖品唯一标识
func (r *TaobaodegoperationcheckaddrstatusAPIRequest) SetSequenceNo(_sequenceNo int64) error {
	r._sequenceNo = _sequenceNo
	r.Set("sequence_no", _sequenceNo)
	return nil
}

// GetSequenceNo SequenceNo Getter
func (r TaobaodegoperationcheckaddrstatusAPIRequest) GetSequenceNo() int64 {
	return r._sequenceNo
}
