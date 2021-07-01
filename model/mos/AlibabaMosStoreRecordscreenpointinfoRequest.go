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
type AlibabaMosStoreRecordscreenpointinfoAPIRequest struct {
    model.Params
    // 云屏埋点信息
    _screenPointInfo   string
}

// 初始化AlibabaMosStoreRecordscreenpointinfoAPIRequest对象
func NewAlibabaMosStoreRecordscreenpointinfoRequest() *AlibabaMosStoreRecordscreenpointinfoAPIRequest{
    return &AlibabaMosStoreRecordscreenpointinfoAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMosStoreRecordscreenpointinfoAPIRequest) GetApiMethodName() string {
    return "alibaba.mos.store.recordscreenpointinfo"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMosStoreRecordscreenpointinfoAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ScreenPointInfo Setter
// 云屏埋点信息
func (r *AlibabaMosStoreRecordscreenpointinfoAPIRequest) SetScreenPointInfo(_screenPointInfo string) error {
    r._screenPointInfo = _screenPointInfo
    r.Set("screen_point_info", _screenPointInfo)
    return nil
}

// ScreenPointInfo Getter
func (r AlibabaMosStoreRecordscreenpointinfoAPIRequest) GetScreenPointInfo() string {
    return r._screenPointInfo
}
