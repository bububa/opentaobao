package media

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取商家视频列表 APIRequest
taobao.media.video.list

用于获取授权商家的视频列表
*/
type TaobaoMediaVideoListRequest struct {
    model.Params

    // 搜索条件
    searchCondition   *VideoSearchCondition2 

}

func NewTaobaoMediaVideoListRequest() *TaobaoMediaVideoListRequest{
    return &TaobaoMediaVideoListRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoMediaVideoListRequest) GetApiMethodName() string {
    return "taobao.media.video.list"
}

func (r TaobaoMediaVideoListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoMediaVideoListRequest) SetSearchCondition(searchCondition *VideoSearchCondition2) error {
    r.searchCondition = searchCondition
    r.Set("search_condition", searchCondition)
    return nil
}

func (r TaobaoMediaVideoListRequest) GetSearchCondition() *VideoSearchCondition2 {
    return r.searchCondition
}

