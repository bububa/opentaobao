package exchange

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallexchangerefusereasongetAPIRequest 获取拒绝换货原因列表 API请求
// tmall.exchange.refusereason.get
//
// 获取拒绝换货原因列表
type TmallexchangerefusereasongetAPIRequest struct {
	model.Params
	// 返回字段
	_fields []string
	// 换货单号ID
	_disputeId int64
	// 换货申请类型：0-任意类型；1-售中；2-售后
	_disputeType int64
}

// NewTmallexchangerefusereasongetRequest 初始化TmallexchangerefusereasongetAPIRequest对象
func NewTmallexchangerefusereasongetRequest() *TmallexchangerefusereasongetAPIRequest {
	return &TmallexchangerefusereasongetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallexchangerefusereasongetAPIRequest) GetApiMethodName() string {
	return "tmall.exchange.refusereason.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallexchangerefusereasongetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallexchangerefusereasongetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFields is Fields Setter
// 返回字段
func (r *TmallexchangerefusereasongetAPIRequest) SetFields(_fields []string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TmallexchangerefusereasongetAPIRequest) GetFields() []string {
	return r._fields
}

// SetDisputeId is DisputeId Setter
// 换货单号ID
func (r *TmallexchangerefusereasongetAPIRequest) SetDisputeId(_disputeId int64) error {
	r._disputeId = _disputeId
	r.Set("dispute_id", _disputeId)
	return nil
}

// GetDisputeId DisputeId Getter
func (r TmallexchangerefusereasongetAPIRequest) GetDisputeId() int64 {
	return r._disputeId
}

// SetDisputeType is DisputeType Setter
// 换货申请类型：0-任意类型；1-售中；2-售后
func (r *TmallexchangerefusereasongetAPIRequest) SetDisputeType(_disputeType int64) error {
	r._disputeType = _disputeType
	r.Set("dispute_type", _disputeType)
	return nil
}

// GetDisputeType DisputeType Getter
func (r TmallexchangerefusereasongetAPIRequest) GetDisputeType() int64 {
	return r._disputeType
}
