package tmallservice

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServiceSettleadjustmentOperateAPIRequest 天猫服务调整单操作 API请求
// tmall.service.settleadjustment.operate
//
// 提供给服务商在对结算有异议时，发起结算调整单。
// 通过说明调整单ID，调整费用值，调整原因进行结算调整单修改。
type TmallServiceSettleadjustmentOperateAPIRequest struct {
	model.Params
	// 审批备注
	_memo string
	// 操作动作
	_operateCode string
	// 调整单ID（也即退款单ID）
	_settlementAdjustmentOrderId int64
}

// NewTmallServiceSettleadjustmentOperateRequest 初始化TmallServiceSettleadjustmentOperateAPIRequest对象
func NewTmallServiceSettleadjustmentOperateRequest() *TmallServiceSettleadjustmentOperateAPIRequest {
	return &TmallServiceSettleadjustmentOperateAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallServiceSettleadjustmentOperateAPIRequest) Reset() {
	r._memo = ""
	r._operateCode = ""
	r._settlementAdjustmentOrderId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServiceSettleadjustmentOperateAPIRequest) GetApiMethodName() string {
	return "tmall.service.settleadjustment.operate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServiceSettleadjustmentOperateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallServiceSettleadjustmentOperateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMemo is Memo Setter
// 审批备注
func (r *TmallServiceSettleadjustmentOperateAPIRequest) SetMemo(_memo string) error {
	r._memo = _memo
	r.Set("memo", _memo)
	return nil
}

// GetMemo Memo Getter
func (r TmallServiceSettleadjustmentOperateAPIRequest) GetMemo() string {
	return r._memo
}

// SetOperateCode is OperateCode Setter
// 操作动作
func (r *TmallServiceSettleadjustmentOperateAPIRequest) SetOperateCode(_operateCode string) error {
	r._operateCode = _operateCode
	r.Set("operate_code", _operateCode)
	return nil
}

// GetOperateCode OperateCode Getter
func (r TmallServiceSettleadjustmentOperateAPIRequest) GetOperateCode() string {
	return r._operateCode
}

// SetSettlementAdjustmentOrderId is SettlementAdjustmentOrderId Setter
// 调整单ID（也即退款单ID）
func (r *TmallServiceSettleadjustmentOperateAPIRequest) SetSettlementAdjustmentOrderId(_settlementAdjustmentOrderId int64) error {
	r._settlementAdjustmentOrderId = _settlementAdjustmentOrderId
	r.Set("settlement_adjustment_order_id", _settlementAdjustmentOrderId)
	return nil
}

// GetSettlementAdjustmentOrderId SettlementAdjustmentOrderId Getter
func (r TmallServiceSettleadjustmentOperateAPIRequest) GetSettlementAdjustmentOrderId() int64 {
	return r._settlementAdjustmentOrderId
}

var poolTmallServiceSettleadjustmentOperateAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallServiceSettleadjustmentOperateRequest()
	},
}

// GetTmallServiceSettleadjustmentOperateRequest 从 sync.Pool 获取 TmallServiceSettleadjustmentOperateAPIRequest
func GetTmallServiceSettleadjustmentOperateAPIRequest() *TmallServiceSettleadjustmentOperateAPIRequest {
	return poolTmallServiceSettleadjustmentOperateAPIRequest.Get().(*TmallServiceSettleadjustmentOperateAPIRequest)
}

// ReleaseTmallServiceSettleadjustmentOperateAPIRequest 将 TmallServiceSettleadjustmentOperateAPIRequest 放入 sync.Pool
func ReleaseTmallServiceSettleadjustmentOperateAPIRequest(v *TmallServiceSettleadjustmentOperateAPIRequest) {
	v.Reset()
	poolTmallServiceSettleadjustmentOperateAPIRequest.Put(v)
}
