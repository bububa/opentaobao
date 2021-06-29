package xiamicontent

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取歌单详情接口 APIRequest
xiami.content.songs.collect.get

根据歌单id，获取歌单详情
*/
type XiamiContentSongsCollectGetRequest struct {
    model.Params

    // 歌单id
    collectId   int64 

    // 分页信息
    page   *PagingVo 

}

func NewXiamiContentSongsCollectGetRequest() *XiamiContentSongsCollectGetRequest{
    return &XiamiContentSongsCollectGetRequest{
        Params: model.NewParams(),
    }
}

func (r XiamiContentSongsCollectGetRequest) GetApiMethodName() string {
    return "xiami.content.songs.collect.get"
}

func (r XiamiContentSongsCollectGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *XiamiContentSongsCollectGetRequest) SetCollectId(collectId int64) error {
    r.collectId = collectId
    r.Set("collect_id", collectId)
    return nil
}

func (r XiamiContentSongsCollectGetRequest) GetCollectId() int64 {
    return r.collectId
}

func (r *XiamiContentSongsCollectGetRequest) SetPage(page *PagingVo) error {
    r.page = page
    r.Set("page", page)
    return nil
}

func (r XiamiContentSongsCollectGetRequest) GetPage() *PagingVo {
    return r.page
}

