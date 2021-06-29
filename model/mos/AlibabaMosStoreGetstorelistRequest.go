package mos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据屏编号获取专柜集 API请求
alibaba.mos.store.getstorelist

根据屏编号获取专柜集
*/
type AlibabaMosStoreGetstorelistRequest struct {
    model.Params
    // 屏编号
    screenNo   string
}

// 初始化AlibabaMosStoreGetstorelistRequest对象
func NewAlibabaMosStoreGetstorelistRequest() *AlibabaMosStoreGetstorelistRequest{
    return &AlibabaMosStoreGetstorelistRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMosStoreGetstorelistRequest) GetApiMethodName() string {
    return "alibaba.mos.store.getstorelist"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMosStoreGetstorelistRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ScreenNo Setter
// 屏编号
func (r *AlibabaMosStoreGetstorelistRequest) SetScreenNo(screenNo string) error {
    r.screenNo = screenNo
    r.Set("screen_no", screenNo)
    return nil
}

// ScreenNo Getter
func (r AlibabaMosStoreGetstorelistRequest) GetScreenNo() string {
    return r.screenNo
}
