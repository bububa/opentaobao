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
type TaobaoTaotvVideoPlaylistPageAPIRequest struct {
    model.Params
    // 客户端信息
    _systemInfo   string
    // 当前页（从1开始）
    _pageNo   int64
}

// 初始化TaobaoTaotvVideoPlaylistPageAPIRequest对象
func NewTaobaoTaotvVideoPlaylistPageRequest() *TaobaoTaotvVideoPlaylistPageAPIRequest{
    return &TaobaoTaotvVideoPlaylistPageAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTaotvVideoPlaylistPageAPIRequest) GetApiMethodName() string {
    return "taobao.taotv.video.playlist.page"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTaotvVideoPlaylistPageAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SystemInfo Setter
// 客户端信息
func (r *TaobaoTaotvVideoPlaylistPageAPIRequest) SetSystemInfo(_systemInfo string) error {
    r._systemInfo = _systemInfo
    r.Set("system_info", _systemInfo)
    return nil
}

// SystemInfo Getter
func (r TaobaoTaotvVideoPlaylistPageAPIRequest) GetSystemInfo() string {
    return r._systemInfo
}
// PageNo Setter
// 当前页（从1开始）
func (r *TaobaoTaotvVideoPlaylistPageAPIRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoTaotvVideoPlaylistPageAPIRequest) GetPageNo() int64 {
    return r._pageNo
}
