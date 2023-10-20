package degoperation

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoDegoperationCheckAddrStatusAPIRequest 地址 API请求
// taobao.degoperation.check.addr.status
//
// 激励
type TaobaoDegoperationCheckAddrStatusAPIRequest struct {
	model.Params
	// 奖品唯一标识
	_sequenceNo int64
}

// NewTaobaoDegoperationCheckAddrStatusRequest 初始化TaobaoDegoperationCheckAddrStatusAPIRequest对象
func NewTaobaoDegoperationCheckAddrStatusRequest() *TaobaoDegoperationCheckAddrStatusAPIRequest {
	return &TaobaoDegoperationCheckAddrStatusAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoDegoperationCheckAddrStatusAPIRequest) Reset() {
	r._sequenceNo = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoDegoperationCheckAddrStatusAPIRequest) GetApiMethodName() string {
	return "taobao.degoperation.check.addr.status"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoDegoperationCheckAddrStatusAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoDegoperationCheckAddrStatusAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSequenceNo is SequenceNo Setter
// 奖品唯一标识
func (r *TaobaoDegoperationCheckAddrStatusAPIRequest) SetSequenceNo(_sequenceNo int64) error {
	r._sequenceNo = _sequenceNo
	r.Set("sequence_no", _sequenceNo)
	return nil
}

// GetSequenceNo SequenceNo Getter
func (r TaobaoDegoperationCheckAddrStatusAPIRequest) GetSequenceNo() int64 {
	return r._sequenceNo
}

var poolTaobaoDegoperationCheckAddrStatusAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoDegoperationCheckAddrStatusRequest()
	},
}

// GetTaobaoDegoperationCheckAddrStatusRequest 从 sync.Pool 获取 TaobaoDegoperationCheckAddrStatusAPIRequest
func GetTaobaoDegoperationCheckAddrStatusAPIRequest() *TaobaoDegoperationCheckAddrStatusAPIRequest {
	return poolTaobaoDegoperationCheckAddrStatusAPIRequest.Get().(*TaobaoDegoperationCheckAddrStatusAPIRequest)
}

// ReleaseTaobaoDegoperationCheckAddrStatusAPIRequest 将 TaobaoDegoperationCheckAddrStatusAPIRequest 放入 sync.Pool
func ReleaseTaobaoDegoperationCheckAddrStatusAPIRequest(v *TaobaoDegoperationCheckAddrStatusAPIRequest) {
	v.Reset()
	poolTaobaoDegoperationCheckAddrStatusAPIRequest.Put(v)
}
