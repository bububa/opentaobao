package mos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据屏编号获取专柜集 APIRequest
alibaba.mos.store.getstorelist

根据屏编号获取专柜集
*/
type AlibabaMosStoreGetstorelistRequest struct {
    model.Params

    // 屏编号
    screenNo   string 

}

func NewAlibabaMosStoreGetstorelistRequest() *AlibabaMosStoreGetstorelistRequest{
    return &AlibabaMosStoreGetstorelistRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaMosStoreGetstorelistRequest) GetApiMethodName() string {
    return "alibaba.mos.store.getstorelist"
}

func (r AlibabaMosStoreGetstorelistRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaMosStoreGetstorelistRequest) SetScreenNo(screenNo string) error {
    r.screenNo = screenNo
    r.Set("screen_no", screenNo)
    return nil
}

func (r AlibabaMosStoreGetstorelistRequest) GetScreenNo() string {
    return r.screenNo
}

