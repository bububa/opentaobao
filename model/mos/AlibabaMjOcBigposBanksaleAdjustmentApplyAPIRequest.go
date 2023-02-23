package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMjOcBigposBanksaleAdjustmentApplyAPIRequest 大pos银行卡调账申请 API请求
// alibaba.mj.oc.bigpos.banksale.adjustment.apply
//
// 大pos银行卡调账申请
type AlibabaMjOcBigposBanksaleAdjustmentApplyAPIRequest struct {
	model.Params
	// 门店号
	_storeNo string
	// 卡号
	_cardNo string
	// 交易时间
	_operTime string
	// 收银员号
	_operator string
	// 调账收银机号
	_posNo string
	// 调账金额
	_amount int64
}

// NewAlibabaMjOcBigposBanksaleAdjustmentApplyRequest 初始化AlibabaMjOcBigposBanksaleAdjustmentApplyAPIRequest对象
func NewAlibabaMjOcBigposBanksaleAdjustmentApplyRequest() *AlibabaMjOcBigposBanksaleAdjustmentApplyAPIRequest {
	return &AlibabaMjOcBigposBanksaleAdjustmentApplyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMjOcBigposBanksaleAdjustmentApplyAPIRequest) GetApiMethodName() string {
	return "alibaba.mj.oc.bigpos.banksale.adjustment.apply"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMjOcBigposBanksaleAdjustmentApplyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMjOcBigposBanksaleAdjustmentApplyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStoreNo is StoreNo Setter
// 门店号
func (r *AlibabaMjOcBigposBanksaleAdjustmentApplyAPIRequest) SetStoreNo(_storeNo string) error {
	r._storeNo = _storeNo
	r.Set("store_no", _storeNo)
	return nil
}

// GetStoreNo StoreNo Getter
func (r AlibabaMjOcBigposBanksaleAdjustmentApplyAPIRequest) GetStoreNo() string {
	return r._storeNo
}

// SetCardNo is CardNo Setter
// 卡号
func (r *AlibabaMjOcBigposBanksaleAdjustmentApplyAPIRequest) SetCardNo(_cardNo string) error {
	r._cardNo = _cardNo
	r.Set("card_no", _cardNo)
	return nil
}

// GetCardNo CardNo Getter
func (r AlibabaMjOcBigposBanksaleAdjustmentApplyAPIRequest) GetCardNo() string {
	return r._cardNo
}

// SetOperTime is OperTime Setter
// 交易时间
func (r *AlibabaMjOcBigposBanksaleAdjustmentApplyAPIRequest) SetOperTime(_operTime string) error {
	r._operTime = _operTime
	r.Set("oper_time", _operTime)
	return nil
}

// GetOperTime OperTime Getter
func (r AlibabaMjOcBigposBanksaleAdjustmentApplyAPIRequest) GetOperTime() string {
	return r._operTime
}

// SetOperator is Operator Setter
// 收银员号
func (r *AlibabaMjOcBigposBanksaleAdjustmentApplyAPIRequest) SetOperator(_operator string) error {
	r._operator = _operator
	r.Set("operator", _operator)
	return nil
}

// GetOperator Operator Getter
func (r AlibabaMjOcBigposBanksaleAdjustmentApplyAPIRequest) GetOperator() string {
	return r._operator
}

// SetPosNo is PosNo Setter
// 调账收银机号
func (r *AlibabaMjOcBigposBanksaleAdjustmentApplyAPIRequest) SetPosNo(_posNo string) error {
	r._posNo = _posNo
	r.Set("pos_no", _posNo)
	return nil
}

// GetPosNo PosNo Getter
func (r AlibabaMjOcBigposBanksaleAdjustmentApplyAPIRequest) GetPosNo() string {
	return r._posNo
}

// SetAmount is Amount Setter
// 调账金额
func (r *AlibabaMjOcBigposBanksaleAdjustmentApplyAPIRequest) SetAmount(_amount int64) error {
	r._amount = _amount
	r.Set("amount", _amount)
	return nil
}

// GetAmount Amount Getter
func (r AlibabaMjOcBigposBanksaleAdjustmentApplyAPIRequest) GetAmount() int64 {
	return r._amount
}
