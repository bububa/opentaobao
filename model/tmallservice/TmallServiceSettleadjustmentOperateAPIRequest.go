package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallservicesettleadjustmentoperateAPIRequest 天猫服务调整单操作 API请求
// tmall.service.settleadjustment.operate
//
// 提供给服务商在对结算有异议时，发起结算调整单。
// 通过说明调整单ID，调整费用值，调整原因进行结算调整单修改。
type TmallservicesettleadjustmentoperateAPIRequest struct {
	model.Params
	// 审批备注
	_memo string
	// 操作动作
	_operateCode string
	// 调整单ID（也即退款单ID）
	_settlementAdjustmentOrderId int64
}

// NewTmallservicesettleadjustmentoperateRequest 初始化TmallservicesettleadjustmentoperateAPIRequest对象
func NewTmallservicesettleadjustmentoperateRequest() *TmallservicesettleadjustmentoperateAPIRequest {
	return &TmallservicesettleadjustmentoperateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallservicesettleadjustmentoperateAPIRequest) GetApiMethodName() string {
	return "tmall.service.settleadjustment.operate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallservicesettleadjustmentoperateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallservicesettleadjustmentoperateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMemo is Memo Setter
// 审批备注
func (r *TmallservicesettleadjustmentoperateAPIRequest) SetMemo(_memo string) error {
	r._memo = _memo
	r.Set("memo", _memo)
	return nil
}

// GetMemo Memo Getter
func (r TmallservicesettleadjustmentoperateAPIRequest) GetMemo() string {
	return r._memo
}

// SetOperateCode is OperateCode Setter
// 操作动作
func (r *TmallservicesettleadjustmentoperateAPIRequest) SetOperateCode(_operateCode string) error {
	r._operateCode = _operateCode
	r.Set("operate_code", _operateCode)
	return nil
}

// GetOperateCode OperateCode Getter
func (r TmallservicesettleadjustmentoperateAPIRequest) GetOperateCode() string {
	return r._operateCode
}

// SetSettlementAdjustmentOrderId is SettlementAdjustmentOrderId Setter
// 调整单ID（也即退款单ID）
func (r *TmallservicesettleadjustmentoperateAPIRequest) SetSettlementAdjustmentOrderId(_settlementAdjustmentOrderId int64) error {
	r._settlementAdjustmentOrderId = _settlementAdjustmentOrderId
	r.Set("settlement_adjustment_order_id", _settlementAdjustmentOrderId)
	return nil
}

// GetSettlementAdjustmentOrderId SettlementAdjustmentOrderId Getter
func (r TmallservicesettleadjustmentoperateAPIRequest) GetSettlementAdjustmentOrderId() int64 {
	return r._settlementAdjustmentOrderId
}
