package mos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
pos机获取线下最大小票号 APIRequest
alibaba.mj.oc.offline.maxticketno.get

给pos机提供线下最大小票号查询
*/
type AlibabaMjOcOfflineMaxticketnoGetRequest struct {
    model.Params

    // 收银机号
    posNo   string 

    // 外部门店号
    storeNo   string 

    // 日期
    datetime   string 

}

func NewAlibabaMjOcOfflineMaxticketnoGetRequest() *AlibabaMjOcOfflineMaxticketnoGetRequest{
    return &AlibabaMjOcOfflineMaxticketnoGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaMjOcOfflineMaxticketnoGetRequest) GetApiMethodName() string {
    return "alibaba.mj.oc.offline.maxticketno.get"
}

func (r AlibabaMjOcOfflineMaxticketnoGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaMjOcOfflineMaxticketnoGetRequest) SetPosNo(posNo string) error {
    r.posNo = posNo
    r.Set("pos_no", posNo)
    return nil
}

func (r AlibabaMjOcOfflineMaxticketnoGetRequest) GetPosNo() string {
    return r.posNo
}

func (r *AlibabaMjOcOfflineMaxticketnoGetRequest) SetStoreNo(storeNo string) error {
    r.storeNo = storeNo
    r.Set("store_no", storeNo)
    return nil
}

func (r AlibabaMjOcOfflineMaxticketnoGetRequest) GetStoreNo() string {
    return r.storeNo
}

func (r *AlibabaMjOcOfflineMaxticketnoGetRequest) SetDatetime(datetime string) error {
    r.datetime = datetime
    r.Set("datetime", datetime)
    return nil
}

func (r AlibabaMjOcOfflineMaxticketnoGetRequest) GetDatetime() string {
    return r.datetime
}

