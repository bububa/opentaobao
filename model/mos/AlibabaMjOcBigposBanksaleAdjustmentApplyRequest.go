package mos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
大pos银行卡调账申请 API请求
alibaba.mj.oc.bigpos.banksale.adjustment.apply

大pos银行卡调账申请
*/
type AlibabaMjOcBigposBanksaleAdjustmentApplyRequest struct {
    model.Params
    // 门店号
    _storeNo   string
    // 调账金额
    _amount   int64
    // 卡号
    _cardNo   string
    // 交易时间
    _operTime   string
    // 收银员号
    _operator   string
    // 调账收银机号
    _posNo   string
}

// 初始化AlibabaMjOcBigposBanksaleAdjustmentApplyRequest对象
func NewAlibabaMjOcBigposBanksaleAdjustmentApplyRequest() *AlibabaMjOcBigposBanksaleAdjustmentApplyRequest{
    return &AlibabaMjOcBigposBanksaleAdjustmentApplyRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMjOcBigposBanksaleAdjustmentApplyRequest) GetApiMethodName() string {
    return "alibaba.mj.oc.bigpos.banksale.adjustment.apply"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMjOcBigposBanksaleAdjustmentApplyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StoreNo Setter
// 门店号
func (r *AlibabaMjOcBigposBanksaleAdjustmentApplyRequest) SetStoreNo(_storeNo string) error {
    r._storeNo = _storeNo
    r.Set("store_no", _storeNo)
    return nil
}

// StoreNo Getter
func (r AlibabaMjOcBigposBanksaleAdjustmentApplyRequest) GetStoreNo() string {
    return r._storeNo
}
// Amount Setter
// 调账金额
func (r *AlibabaMjOcBigposBanksaleAdjustmentApplyRequest) SetAmount(_amount int64) error {
    r._amount = _amount
    r.Set("amount", _amount)
    return nil
}

// Amount Getter
func (r AlibabaMjOcBigposBanksaleAdjustmentApplyRequest) GetAmount() int64 {
    return r._amount
}
// CardNo Setter
// 卡号
func (r *AlibabaMjOcBigposBanksaleAdjustmentApplyRequest) SetCardNo(_cardNo string) error {
    r._cardNo = _cardNo
    r.Set("card_no", _cardNo)
    return nil
}

// CardNo Getter
func (r AlibabaMjOcBigposBanksaleAdjustmentApplyRequest) GetCardNo() string {
    return r._cardNo
}
// OperTime Setter
// 交易时间
func (r *AlibabaMjOcBigposBanksaleAdjustmentApplyRequest) SetOperTime(_operTime string) error {
    r._operTime = _operTime
    r.Set("oper_time", _operTime)
    return nil
}

// OperTime Getter
func (r AlibabaMjOcBigposBanksaleAdjustmentApplyRequest) GetOperTime() string {
    return r._operTime
}
// Operator Setter
// 收银员号
func (r *AlibabaMjOcBigposBanksaleAdjustmentApplyRequest) SetOperator(_operator string) error {
    r._operator = _operator
    r.Set("operator", _operator)
    return nil
}

// Operator Getter
func (r AlibabaMjOcBigposBanksaleAdjustmentApplyRequest) GetOperator() string {
    return r._operator
}
// PosNo Setter
// 调账收银机号
func (r *AlibabaMjOcBigposBanksaleAdjustmentApplyRequest) SetPosNo(_posNo string) error {
    r._posNo = _posNo
    r.Set("pos_no", _posNo)
    return nil
}

// PosNo Getter
func (r AlibabaMjOcBigposBanksaleAdjustmentApplyRequest) GetPosNo() string {
    return r._posNo
}
