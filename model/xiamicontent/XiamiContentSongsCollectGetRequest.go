package xiamicontent

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取歌单详情接口 API请求
xiami.content.songs.collect.get

根据歌单id，获取歌单详情
*/
type XiamiContentSongsCollectGetRequest struct {
    model.Params
    // 歌单id
    _collectId   int64
    // 分页信息
    _page   *PagingVo
}

// 初始化XiamiContentSongsCollectGetRequest对象
func NewXiamiContentSongsCollectGetRequest() *XiamiContentSongsCollectGetRequest{
    return &XiamiContentSongsCollectGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r XiamiContentSongsCollectGetRequest) GetApiMethodName() string {
    return "xiami.content.songs.collect.get"
}

// IRequest interface 方法, 获取API参数
func (r XiamiContentSongsCollectGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CollectId Setter
// 歌单id
func (r *XiamiContentSongsCollectGetRequest) SetCollectId(_collectId int64) error {
    r._collectId = _collectId
    r.Set("collect_id", _collectId)
    return nil
}

// CollectId Getter
func (r XiamiContentSongsCollectGetRequest) GetCollectId() int64 {
    return r._collectId
}
// Page Setter
// 分页信息
func (r *XiamiContentSongsCollectGetRequest) SetPage(_page *PagingVo) error {
    r._page = _page
    r.Set("page", _page)
    return nil
}

// Page Getter
func (r XiamiContentSongsCollectGetRequest) GetPage() *PagingVo {
    return r._page
}
