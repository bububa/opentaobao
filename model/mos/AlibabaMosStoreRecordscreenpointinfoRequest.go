package mos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
云屏埋点数据记录接口 APIRequest
alibaba.mos.store.recordscreenpointinfo

记录云屏埋点数据
*/
type AlibabaMosStoreRecordscreenpointinfoRequest struct {
    model.Params

    // 云屏埋点信息
    screenPointInfo   string 

}

func NewAlibabaMosStoreRecordscreenpointinfoRequest() *AlibabaMosStoreRecordscreenpointinfoRequest{
    return &AlibabaMosStoreRecordscreenpointinfoRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaMosStoreRecordscreenpointinfoRequest) GetApiMethodName() string {
    return "alibaba.mos.store.recordscreenpointinfo"
}

func (r AlibabaMosStoreRecordscreenpointinfoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaMosStoreRecordscreenpointinfoRequest) SetScreenPointInfo(screenPointInfo string) error {
    r.screenPointInfo = screenPointInfo
    r.Set("screen_point_info", screenPointInfo)
    return nil
}

func (r AlibabaMosStoreRecordscreenpointinfoRequest) GetScreenPointInfo() string {
    return r.screenPointInfo
}

