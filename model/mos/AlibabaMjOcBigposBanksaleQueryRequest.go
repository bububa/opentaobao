package mos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
大pos银行卡查账接口 APIRequest
alibaba.mj.oc.bigpos.banksale.query

大pos银行卡查账接口，给收银员查询银行卡销售记录，便于调账
*/
type AlibabaMjOcBigposBanksaleQueryRequest struct {
    model.Params

    // 开始时间
    startTime   string 

    // 卡号
    cardNo   string 

    // 外部门店号
    outStoreNo   string 

    // 结束时间
    endTime   string 

}

func NewAlibabaMjOcBigposBanksaleQueryRequest() *AlibabaMjOcBigposBanksaleQueryRequest{
    return &AlibabaMjOcBigposBanksaleQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaMjOcBigposBanksaleQueryRequest) GetApiMethodName() string {
    return "alibaba.mj.oc.bigpos.banksale.query"
}

func (r AlibabaMjOcBigposBanksaleQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaMjOcBigposBanksaleQueryRequest) SetStartTime(startTime string) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

func (r AlibabaMjOcBigposBanksaleQueryRequest) GetStartTime() string {
    return r.startTime
}

func (r *AlibabaMjOcBigposBanksaleQueryRequest) SetCardNo(cardNo string) error {
    r.cardNo = cardNo
    r.Set("card_no", cardNo)
    return nil
}

func (r AlibabaMjOcBigposBanksaleQueryRequest) GetCardNo() string {
    return r.cardNo
}

func (r *AlibabaMjOcBigposBanksaleQueryRequest) SetOutStoreNo(outStoreNo string) error {
    r.outStoreNo = outStoreNo
    r.Set("out_store_no", outStoreNo)
    return nil
}

func (r AlibabaMjOcBigposBanksaleQueryRequest) GetOutStoreNo() string {
    return r.outStoreNo
}

func (r *AlibabaMjOcBigposBanksaleQueryRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

func (r AlibabaMjOcBigposBanksaleQueryRequest) GetEndTime() string {
    return r.endTime
}

