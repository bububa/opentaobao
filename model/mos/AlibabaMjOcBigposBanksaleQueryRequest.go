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
    _startTime   string
    // 卡号
    _cardNo   string
    // 外部门店号
    _outStoreNo   string
    // 结束时间
    _endTime   string
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
func (r *AlibabaMjOcBigposBanksaleQueryRequest) SetStartTime(_startTime string) error {
    r._startTime = _startTime
    r.Set("start_time", _startTime)
    return nil
}

// StartTime Getter
func (r AlibabaMjOcBigposBanksaleQueryRequest) GetStartTime() string {
    return r._startTime
}
// CardNo Setter
// 卡号
func (r *AlibabaMjOcBigposBanksaleQueryRequest) SetCardNo(_cardNo string) error {
    r._cardNo = _cardNo
    r.Set("card_no", _cardNo)
    return nil
}

// CardNo Getter
func (r AlibabaMjOcBigposBanksaleQueryRequest) GetCardNo() string {
    return r._cardNo
}
// OutStoreNo Setter
// 外部门店号
func (r *AlibabaMjOcBigposBanksaleQueryRequest) SetOutStoreNo(_outStoreNo string) error {
    r._outStoreNo = _outStoreNo
    r.Set("out_store_no", _outStoreNo)
    return nil
}

// OutStoreNo Getter
func (r AlibabaMjOcBigposBanksaleQueryRequest) GetOutStoreNo() string {
    return r._outStoreNo
}
// EndTime Setter
// 结束时间
func (r *AlibabaMjOcBigposBanksaleQueryRequest) SetEndTime(_endTime string) error {
    r._endTime = _endTime
    r.Set("end_time", _endTime)
    return nil
}

// EndTime Getter
func (r AlibabaMjOcBigposBanksaleQueryRequest) GetEndTime() string {
    return r._endTime
}
