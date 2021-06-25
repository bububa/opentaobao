package mos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
大pos银行卡调账申请 APIRequest
alibaba.mj.oc.bigpos.banksale.adjustment.apply

大pos银行卡调账申请
*/
type AlibabaMjOcBigposBanksaleAdjustmentApplyRequest struct {
    model.Params

    // 门店号
    storeNo   string 

    // 调账金额
    amount   int64 

    // 卡号
    cardNo   string 

    // 交易时间
    operTime   string 

    // 收银员号
    operator   string 

    // 调账收银机号
    posNo   string 

}

func NewAlibabaMjOcBigposBanksaleAdjustmentApplyRequest() *AlibabaMjOcBigposBanksaleAdjustmentApplyRequest{
    return &AlibabaMjOcBigposBanksaleAdjustmentApplyRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaMjOcBigposBanksaleAdjustmentApplyRequest) GetApiMethodName() string {
    return "alibaba.mj.oc.bigpos.banksale.adjustment.apply"
}

func (r AlibabaMjOcBigposBanksaleAdjustmentApplyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaMjOcBigposBanksaleAdjustmentApplyRequest) SetStoreNo(storeNo string) error {
    r.storeNo = storeNo
    r.Set("store_no", storeNo)
    return nil
}

func (r AlibabaMjOcBigposBanksaleAdjustmentApplyRequest) GetStoreNo() string {
    return r.storeNo
}

func (r *AlibabaMjOcBigposBanksaleAdjustmentApplyRequest) SetAmount(amount int64) error {
    r.amount = amount
    r.Set("amount", amount)
    return nil
}

func (r AlibabaMjOcBigposBanksaleAdjustmentApplyRequest) GetAmount() int64 {
    return r.amount
}

func (r *AlibabaMjOcBigposBanksaleAdjustmentApplyRequest) SetCardNo(cardNo string) error {
    r.cardNo = cardNo
    r.Set("card_no", cardNo)
    return nil
}

func (r AlibabaMjOcBigposBanksaleAdjustmentApplyRequest) GetCardNo() string {
    return r.cardNo
}

func (r *AlibabaMjOcBigposBanksaleAdjustmentApplyRequest) SetOperTime(operTime string) error {
    r.operTime = operTime
    r.Set("oper_time", operTime)
    return nil
}

func (r AlibabaMjOcBigposBanksaleAdjustmentApplyRequest) GetOperTime() string {
    return r.operTime
}

func (r *AlibabaMjOcBigposBanksaleAdjustmentApplyRequest) SetOperator(operator string) error {
    r.operator = operator
    r.Set("operator", operator)
    return nil
}

func (r AlibabaMjOcBigposBanksaleAdjustmentApplyRequest) GetOperator() string {
    return r.operator
}

func (r *AlibabaMjOcBigposBanksaleAdjustmentApplyRequest) SetPosNo(posNo string) error {
    r.posNo = posNo
    r.Set("pos_no", posNo)
    return nil
}

func (r AlibabaMjOcBigposBanksaleAdjustmentApplyRequest) GetPosNo() string {
    return r.posNo
}

