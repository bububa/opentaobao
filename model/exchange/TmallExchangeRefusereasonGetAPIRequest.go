package exchange

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallExchangeRefusereasonGetAPIRequest
获取拒绝换货原因列表 API请求
tmall.exchange.refusereason.get

获取拒绝换货原因列表 */
type TmallExchangeRefusereasonGetAPIRequest struct {
	model.Params
	// 换货单号ID
	_disputeId int64
	// 返回字段
	_fields []string
	// 换货申请类型：0-任意类型；1-售中；2-售后
	_disputeType int64
}

// NewTmallExchangeRefusereasonGetRequest 初始化TmallExchangeRefusereasonGetAPIRequest对象
func NewTmallExchangeRefusereasonGetRequest() *TmallExchangeRefusereasonGetAPIRequest {
	return &TmallExchangeRefusereasonGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallExchangeRefusereasonGetAPIRequest) GetApiMethodName() string {
	return "tmall.exchange.refusereason.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallExchangeRefusereasonGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is DisputeId Setter
// 换货单号ID
func (r *TmallExchangeRefusereasonGetAPIRequest) SetDisputeId(_disputeId int64) error {
	r._disputeId = _disputeId
	r.Set("dispute_id", _disputeId)
	return nil
}

// Get DisputeId Getter
func (r TmallExchangeRefusereasonGetAPIRequest) GetDisputeId() int64 {
	return r._disputeId
}

// Set is Fields Setter
// 返回字段
func (r *TmallExchangeRefusereasonGetAPIRequest) SetFields(_fields []string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// Get Fields Getter
func (r TmallExchangeRefusereasonGetAPIRequest) GetFields() []string {
	return r._fields
}

// Set is DisputeType Setter
// 换货申请类型：0-任意类型；1-售中；2-售后
func (r *TmallExchangeRefusereasonGetAPIRequest) SetDisputeType(_disputeType int64) error {
	r._disputeType = _disputeType
	r.Set("dispute_type", _disputeType)
	return nil
}

// Get DisputeType Getter
func (r TmallExchangeRefusereasonGetAPIRequest) GetDisputeType() int64 {
	return r._disputeType
}
