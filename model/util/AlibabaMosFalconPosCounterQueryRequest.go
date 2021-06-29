package util

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
云POS查看专柜属性 API请求
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

// 初始化AlibabaMosFalconPosCounterQueryRequest对象
func NewAlibabaMosFalconPosCounterQueryRequest() *AlibabaMosFalconPosCounterQueryRequest{
    return &AlibabaMosFalconPosCounterQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMosFalconPosCounterQueryRequest) GetApiMethodName() string {
    return "alibaba.mos.falcon.pos.counter.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMosFalconPosCounterQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Sn Setter
// 设备序列号
func (r *AlibabaMosFalconPosCounterQueryRequest) SetSn(sn string) error {
    r.sn = sn
    r.Set("sn", sn)
    return nil
}

// Sn Getter
func (r AlibabaMosFalconPosCounterQueryRequest) GetSn() string {
    return r.sn
}
// StoreNo Setter
// 门店号
func (r *AlibabaMosFalconPosCounterQueryRequest) SetStoreNo(storeNo string) error {
    r.storeNo = storeNo
    r.Set("store_no", storeNo)
    return nil
}

// StoreNo Getter
func (r AlibabaMosFalconPosCounterQueryRequest) GetStoreNo() string {
    return r.storeNo
}
// CounterNo Setter
// 专柜号
func (r *AlibabaMosFalconPosCounterQueryRequest) SetCounterNo(counterNo string) error {
    r.counterNo = counterNo
    r.Set("counter_no", counterNo)
    return nil
}

// CounterNo Getter
func (r AlibabaMosFalconPosCounterQueryRequest) GetCounterNo() string {
    return r.counterNo
}
