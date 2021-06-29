package taotv

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取播单列表 API请求
taobao.taotv.video.playlist.all

根据牌照和视频源等获取播单列表
*/
type TaobaoTaotvVideoPlaylistAllRequest struct {
    model.Params
    // 系统信息
    systemInfo   string
}

// 初始化TaobaoTaotvVideoPlaylistAllRequest对象
func NewTaobaoTaotvVideoPlaylistAllRequest() *TaobaoTaotvVideoPlaylistAllRequest{
    return &TaobaoTaotvVideoPlaylistAllRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTaotvVideoPlaylistAllRequest) GetApiMethodName() string {
    return "taobao.taotv.video.playlist.all"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTaotvVideoPlaylistAllRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SystemInfo Setter
// 系统信息
func (r *TaobaoTaotvVideoPlaylistAllRequest) SetSystemInfo(systemInfo string) error {
    r.systemInfo = systemInfo
    r.Set("system_info", systemInfo)
    return nil
}

// SystemInfo Getter
func (r TaobaoTaotvVideoPlaylistAllRequest) GetSystemInfo() string {
    return r.systemInfo
}
