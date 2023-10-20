package qianniu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoDaogoubaoOrderStatisticsTotalAPIRequest 销售订单总额统计 API请求
// taobao.daogoubao.order.statistics.total
//
// 对接千牛端数字中心
type TaobaoDaogoubaoOrderStatisticsTotalAPIRequest struct {
	model.Params
	// 调试时用的传入id
	_debugId string
	// 需要的字段名
	_field string
}

// NewTaobaoDaogoubaoOrderStatisticsTotalRequest 初始化TaobaoDaogoubaoOrderStatisticsTotalAPIRequest对象
func NewTaobaoDaogoubaoOrderStatisticsTotalRequest() *TaobaoDaogoubaoOrderStatisticsTotalAPIRequest {
	return &TaobaoDaogoubaoOrderStatisticsTotalAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoDaogoubaoOrderStatisticsTotalAPIRequest) GetApiMethodName() string {
	return "taobao.daogoubao.order.statistics.total"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoDaogoubaoOrderStatisticsTotalAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoDaogoubaoOrderStatisticsTotalAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDebugId is DebugId Setter
// 调试时用的传入id
func (r *TaobaoDaogoubaoOrderStatisticsTotalAPIRequest) SetDebugId(_debugId string) error {
	r._debugId = _debugId
	r.Set("debug_id", _debugId)
	return nil
}

// GetDebugId DebugId Getter
func (r TaobaoDaogoubaoOrderStatisticsTotalAPIRequest) GetDebugId() string {
	return r._debugId
}

// SetField is Field Setter
// 需要的字段名
func (r *TaobaoDaogoubaoOrderStatisticsTotalAPIRequest) SetField(_field string) error {
	r._field = _field
	r.Set("field", _field)
	return nil
}

// GetField Field Getter
func (r TaobaoDaogoubaoOrderStatisticsTotalAPIRequest) GetField() string {
	return r._field
}
