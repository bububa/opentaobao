package mos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
大pos银行卡查账接口 API请求
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

// 初始化AlibabaMjOcBigposBanksaleQueryRequest对象
func NewAlibabaMjOcBigposBanksaleQueryRequest() *AlibabaMjOcBigposBanksaleQueryRequest{
    return &AlibabaMjOcBigposBanksaleQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMjOcBigposBanksaleQueryRequest) GetApiMethodName() string {
    return "alibaba.mj.oc.bigpos.banksale.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMjOcBigposBanksaleQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StartTime Setter
// 开始时间
func (r *AlibabaMjOcBigposBanksaleQueryRequest) SetStartTime(startTime string) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

// StartTime Getter
func (r AlibabaMjOcBigposBanksaleQueryRequest) GetStartTime() string {
    return r.startTime
}
// CardNo Setter
// 卡号
func (r *AlibabaMjOcBigposBanksaleQueryRequest) SetCardNo(cardNo string) error {
    r.cardNo = cardNo
    r.Set("card_no", cardNo)
    return nil
}

// CardNo Getter
func (r AlibabaMjOcBigposBanksaleQueryRequest) GetCardNo() string {
    return r.cardNo
}
// OutStoreNo Setter
// 外部门店号
func (r *AlibabaMjOcBigposBanksaleQueryRequest) SetOutStoreNo(outStoreNo string) error {
    r.outStoreNo = outStoreNo
    r.Set("out_store_no", outStoreNo)
    return nil
}

// OutStoreNo Getter
func (r AlibabaMjOcBigposBanksaleQueryRequest) GetOutStoreNo() string {
    return r.outStoreNo
}
// EndTime Setter
// 结束时间
func (r *AlibabaMjOcBigposBanksaleQueryRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

// EndTime Getter
func (r AlibabaMjOcBigposBanksaleQueryRequest) GetEndTime() string {
    return r.endTime
}
