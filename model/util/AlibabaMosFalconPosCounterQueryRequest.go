package util

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
云POS查看专柜属性 APIRequest
alibaba.mos.falcon.pos.counter.query

银泰商业获取专柜是否支持小数等属性查看
*/
type AlibabaMosFalconPosCounterQueryRequest struct {
    model.Params

    // 设备序列号
    sn   string 

    // 门店号
    storeNo   string 

    // 专柜号
    counterNo   string 

}

func NewAlibabaMosFalconPosCounterQueryRequest() *AlibabaMosFalconPosCounterQueryRequest{
    return &AlibabaMosFalconPosCounterQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaMosFalconPosCounterQueryRequest) GetApiMethodName() string {
    return "alibaba.mos.falcon.pos.counter.query"
}

func (r AlibabaMosFalconPosCounterQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaMosFalconPosCounterQueryRequest) SetSn(sn string) error {
    r.sn = sn
    r.Set("sn", sn)
    return nil
}

func (r AlibabaMosFalconPosCounterQueryRequest) GetSn() string {
    return r.sn
}

func (r *AlibabaMosFalconPosCounterQueryRequest) SetStoreNo(storeNo string) error {
    r.storeNo = storeNo
    r.Set("store_no", storeNo)
    return nil
}

func (r AlibabaMosFalconPosCounterQueryRequest) GetStoreNo() string {
    return r.storeNo
}

func (r *AlibabaMosFalconPosCounterQueryRequest) SetCounterNo(counterNo string) error {
    r.counterNo = counterNo
    r.Set("counter_no", counterNo)
    return nil
}

func (r AlibabaMosFalconPosCounterQueryRequest) GetCounterNo() string {
    return r.counterNo
}

