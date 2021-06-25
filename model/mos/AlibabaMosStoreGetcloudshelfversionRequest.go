package mos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取云货架版本信息 APIRequest
alibaba.mos.store.getcloudshelfversion

根据屏编号获取云货架版本信息
*/
type AlibabaMosStoreGetcloudshelfversionRequest struct {
    model.Params

    // 屏编号
    screenNo   string 

}

func NewAlibabaMosStoreGetcloudshelfversionRequest() *AlibabaMosStoreGetcloudshelfversionRequest{
    return &AlibabaMosStoreGetcloudshelfversionRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaMosStoreGetcloudshelfversionRequest) GetApiMethodName() string {
    return "alibaba.mos.store.getcloudshelfversion"
}

func (r AlibabaMosStoreGetcloudshelfversionRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaMosStoreGetcloudshelfversionRequest) SetScreenNo(screenNo string) error {
    r.screenNo = screenNo
    r.Set("screen_no", screenNo)
    return nil
}

func (r AlibabaMosStoreGetcloudshelfversionRequest) GetScreenNo() string {
    return r.screenNo
}

