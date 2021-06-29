package mos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
pos机获取线下最大小票号 API请求
alibaba.mj.oc.offline.maxticketno.get

给pos机提供线下最大小票号查询
*/
type AlibabaMjOcOfflineMaxticketnoGetRequest struct {
    model.Params
    // 收银机号
    _posNo   string
    // 外部门店号
    _storeNo   string
    // 日期
    _datetime   string
}

// 初始化AlibabaMjOcOfflineMaxticketnoGetRequest对象
func NewAlibabaMjOcOfflineMaxticketnoGetRequest() *AlibabaMjOcOfflineMaxticketnoGetRequest{
    return &AlibabaMjOcOfflineMaxticketnoGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMjOcOfflineMaxticketnoGetRequest) GetApiMethodName() string {
    return "alibaba.mj.oc.offline.maxticketno.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMjOcOfflineMaxticketnoGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PosNo Setter
// 收银机号
func (r *AlibabaMjOcOfflineMaxticketnoGetRequest) SetPosNo(_posNo string) error {
    r._posNo = _posNo
    r.Set("pos_no", _posNo)
    return nil
}

// PosNo Getter
func (r AlibabaMjOcOfflineMaxticketnoGetRequest) GetPosNo() string {
    return r._posNo
}
// StoreNo Setter
// 外部门店号
func (r *AlibabaMjOcOfflineMaxticketnoGetRequest) SetStoreNo(_storeNo string) error {
    r._storeNo = _storeNo
    r.Set("store_no", _storeNo)
    return nil
}

// StoreNo Getter
func (r AlibabaMjOcOfflineMaxticketnoGetRequest) GetStoreNo() string {
    return r._storeNo
}
// Datetime Setter
// 日期
func (r *AlibabaMjOcOfflineMaxticketnoGetRequest) SetDatetime(_datetime string) error {
    r._datetime = _datetime
    r.Set("datetime", _datetime)
    return nil
}

// Datetime Getter
func (r AlibabaMjOcOfflineMaxticketnoGetRequest) GetDatetime() string {
    return r._datetime
}
