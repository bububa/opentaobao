package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabamjocbigposbanksaleadjustmentapplyAPIRequest 大pos银行卡调账申请 API请求
// alibaba.mj.oc.bigpos.banksale.adjustment.apply
//
// 大pos银行卡调账申请
type AlibabamjocbigposbanksaleadjustmentapplyAPIRequest struct {
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

// NewAlibabamjocbigposbanksaleadjustmentapplyRequest 初始化AlibabamjocbigposbanksaleadjustmentapplyAPIRequest对象
func NewAlibabamjocbigposbanksaleadjustmentapplyRequest() *AlibabamjocbigposbanksaleadjustmentapplyAPIRequest {
	return &AlibabamjocbigposbanksaleadjustmentapplyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabamjocbigposbanksaleadjustmentapplyAPIRequest) GetApiMethodName() string {
	return "alibaba.mj.oc.bigpos.banksale.adjustment.apply"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabamjocbigposbanksaleadjustmentapplyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabamjocbigposbanksaleadjustmentapplyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStoreNo is StoreNo Setter
// 门店号
func (r *AlibabamjocbigposbanksaleadjustmentapplyAPIRequest) SetStoreNo(_storeNo string) error {
	r._storeNo = _storeNo
	r.Set("store_no", _storeNo)
	return nil
}

// GetStoreNo StoreNo Getter
func (r AlibabamjocbigposbanksaleadjustmentapplyAPIRequest) GetStoreNo() string {
	return r._storeNo
}

// SetCardNo is CardNo Setter
// 卡号
func (r *AlibabamjocbigposbanksaleadjustmentapplyAPIRequest) SetCardNo(_cardNo string) error {
	r._cardNo = _cardNo
	r.Set("card_no", _cardNo)
	return nil
}

// GetCardNo CardNo Getter
func (r AlibabamjocbigposbanksaleadjustmentapplyAPIRequest) GetCardNo() string {
	return r._cardNo
}

// SetOperTime is OperTime Setter
// 交易时间
func (r *AlibabamjocbigposbanksaleadjustmentapplyAPIRequest) SetOperTime(_operTime string) error {
	r._operTime = _operTime
	r.Set("oper_time", _operTime)
	return nil
}

// GetOperTime OperTime Getter
func (r AlibabamjocbigposbanksaleadjustmentapplyAPIRequest) GetOperTime() string {
	return r._operTime
}

// SetOperator is Operator Setter
// 收银员号
func (r *AlibabamjocbigposbanksaleadjustmentapplyAPIRequest) SetOperator(_operator string) error {
	r._operator = _operator
	r.Set("operator", _operator)
	return nil
}

// GetOperator Operator Getter
func (r AlibabamjocbigposbanksaleadjustmentapplyAPIRequest) GetOperator() string {
	return r._operator
}

// SetPosNo is PosNo Setter
// 调账收银机号
func (r *AlibabamjocbigposbanksaleadjustmentapplyAPIRequest) SetPosNo(_posNo string) error {
	r._posNo = _posNo
	r.Set("pos_no", _posNo)
	return nil
}

// GetPosNo PosNo Getter
func (r AlibabamjocbigposbanksaleadjustmentapplyAPIRequest) GetPosNo() string {
	return r._posNo
}

// SetAmount is Amount Setter
// 调账金额
func (r *AlibabamjocbigposbanksaleadjustmentapplyAPIRequest) SetAmount(_amount int64) error {
	r._amount = _amount
	r.Set("amount", _amount)
	return nil
}

// GetAmount Amount Getter
func (r AlibabamjocbigposbanksaleadjustmentapplyAPIRequest) GetAmount() int64 {
	return r._amount
}
