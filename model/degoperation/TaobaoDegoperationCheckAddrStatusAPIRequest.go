package degoperation

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoDegoperationCheckAddrStatusAPIRequest
地址 API请求
taobao.degoperation.check.addr.status

激励 */
type TaobaoDegoperationCheckAddrStatusAPIRequest struct {
	model.Params
	// 奖品唯一标识
	_sequenceNo int64
}

// NewTaobaoDegoperationCheckAddrStatusRequest 初始化TaobaoDegoperationCheckAddrStatusAPIRequest对象
func NewTaobaoDegoperationCheckAddrStatusRequest() *TaobaoDegoperationCheckAddrStatusAPIRequest {
	return &TaobaoDegoperationCheckAddrStatusAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoDegoperationCheckAddrStatusAPIRequest) GetApiMethodName() string {
	return "taobao.degoperation.check.addr.status"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoDegoperationCheckAddrStatusAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is SequenceNo Setter
// 奖品唯一标识
func (r *TaobaoDegoperationCheckAddrStatusAPIRequest) SetSequenceNo(_sequenceNo int64) error {
	r._sequenceNo = _sequenceNo
	r.Set("sequence_no", _sequenceNo)
	return nil
}

// Get SequenceNo Getter
func (r TaobaoDegoperationCheckAddrStatusAPIRequest) GetSequenceNo() int64 {
	return r._sequenceNo
}
