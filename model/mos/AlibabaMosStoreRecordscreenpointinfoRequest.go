package mos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
云屏埋点数据记录接口 API请求
alibaba.mos.store.recordscreenpointinfo

记录云屏埋点数据
*/
type AlibabaMosStoreRecordscreenpointinfoRequest struct {
    model.Params
    // 云屏埋点信息
    screenPointInfo   string
}

// 初始化AlibabaMosStoreRecordscreenpointinfoRequest对象
func NewAlibabaMosStoreRecordscreenpointinfoRequest() *AlibabaMosStoreRecordscreenpointinfoRequest{
    return &AlibabaMosStoreRecordscreenpointinfoRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMosStoreRecordscreenpointinfoRequest) GetApiMethodName() string {
    return "alibaba.mos.store.recordscreenpointinfo"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMosStoreRecordscreenpointinfoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ScreenPointInfo Setter
// 云屏埋点信息
func (r *AlibabaMosStoreRecordscreenpointinfoRequest) SetScreenPointInfo(screenPointInfo string) error {
    r.screenPointInfo = screenPointInfo
    r.Set("screen_point_info", screenPointInfo)
    return nil
}

// ScreenPointInfo Getter
func (r AlibabaMosStoreRecordscreenpointinfoRequest) GetScreenPointInfo() string {
    return r.screenPointInfo
}
