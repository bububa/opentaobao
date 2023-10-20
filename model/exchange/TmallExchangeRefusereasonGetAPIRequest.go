package exchange

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallExchangeRefusereasonGetAPIRequest 获取拒绝换货原因列表 API请求
// tmall.exchange.refusereason.get
//
// 获取拒绝换货原因列表
type TmallExchangeRefusereasonGetAPIRequest struct {
	model.Params
	// 返回字段
	_fields []string
	// 换货单号ID
	_disputeId int64
	// 换货申请类型：0-任意类型；1-售中；2-售后
	_disputeType int64
}

// NewTmallExchangeRefusereasonGetRequest 初始化TmallExchangeRefusereasonGetAPIRequest对象
func NewTmallExchangeRefusereasonGetRequest() *TmallExchangeRefusereasonGetAPIRequest {
	return &TmallExchangeRefusereasonGetAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallExchangeRefusereasonGetAPIRequest) Reset() {
	r._fields = r._fields[:0]
	r._disputeId = 0
	r._disputeType = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallExchangeRefusereasonGetAPIRequest) GetApiMethodName() string {
	return "tmall.exchange.refusereason.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallExchangeRefusereasonGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallExchangeRefusereasonGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFields is Fields Setter
// 返回字段
func (r *TmallExchangeRefusereasonGetAPIRequest) SetFields(_fields []string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TmallExchangeRefusereasonGetAPIRequest) GetFields() []string {
	return r._fields
}

// SetDisputeId is DisputeId Setter
// 换货单号ID
func (r *TmallExchangeRefusereasonGetAPIRequest) SetDisputeId(_disputeId int64) error {
	r._disputeId = _disputeId
	r.Set("dispute_id", _disputeId)
	return nil
}

// GetDisputeId DisputeId Getter
func (r TmallExchangeRefusereasonGetAPIRequest) GetDisputeId() int64 {
	return r._disputeId
}

// SetDisputeType is DisputeType Setter
// 换货申请类型：0-任意类型；1-售中；2-售后
func (r *TmallExchangeRefusereasonGetAPIRequest) SetDisputeType(_disputeType int64) error {
	r._disputeType = _disputeType
	r.Set("dispute_type", _disputeType)
	return nil
}

// GetDisputeType DisputeType Getter
func (r TmallExchangeRefusereasonGetAPIRequest) GetDisputeType() int64 {
	return r._disputeType
}

var poolTmallExchangeRefusereasonGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallExchangeRefusereasonGetRequest()
	},
}

// GetTmallExchangeRefusereasonGetRequest 从 sync.Pool 获取 TmallExchangeRefusereasonGetAPIRequest
func GetTmallExchangeRefusereasonGetAPIRequest() *TmallExchangeRefusereasonGetAPIRequest {
	return poolTmallExchangeRefusereasonGetAPIRequest.Get().(*TmallExchangeRefusereasonGetAPIRequest)
}

// ReleaseTmallExchangeRefusereasonGetAPIRequest 将 TmallExchangeRefusereasonGetAPIRequest 放入 sync.Pool
func ReleaseTmallExchangeRefusereasonGetAPIRequest(v *TmallExchangeRefusereasonGetAPIRequest) {
	v.Reset()
	poolTmallExchangeRefusereasonGetAPIRequest.Put(v)
}
