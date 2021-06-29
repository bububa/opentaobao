package taotv

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分页获取所有播单 API请求
taobao.taotv.video.playlist.page

获取所有播单信息（分页）
*/
type TaobaoTaotvVideoPlaylistPageRequest struct {
    model.Params
    // 客户端信息
    systemInfo   string
    // 当前页（从1开始）
    pageNo   int64
}

// 初始化TaobaoTaotvVideoPlaylistPageRequest对象
func NewTaobaoTaotvVideoPlaylistPageRequest() *TaobaoTaotvVideoPlaylistPageRequest{
    return &TaobaoTaotvVideoPlaylistPageRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTaotvVideoPlaylistPageRequest) GetApiMethodName() string {
    return "taobao.taotv.video.playlist.page"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTaotvVideoPlaylistPageRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SystemInfo Setter
// 客户端信息
func (r *TaobaoTaotvVideoPlaylistPageRequest) SetSystemInfo(systemInfo string) error {
    r.systemInfo = systemInfo
    r.Set("system_info", systemInfo)
    return nil
}

// SystemInfo Getter
func (r TaobaoTaotvVideoPlaylistPageRequest) GetSystemInfo() string {
    return r.systemInfo
}
// PageNo Setter
// 当前页（从1开始）
func (r *TaobaoTaotvVideoPlaylistPageRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoTaotvVideoPlaylistPageRequest) GetPageNo() int64 {
    return r.pageNo
}
