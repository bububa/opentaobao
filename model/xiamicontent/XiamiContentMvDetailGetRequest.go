package xiamicontent

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取mv详情 APIRequest
xiami.content.mv.detail.get

获取mv详情
*/
type XiamiContentMvDetailGetRequest struct {
    model.Params

    // mvId
    mvIds   []int64 

}

func NewXiamiContentMvDetailGetRequest() *XiamiContentMvDetailGetRequest{
    return &XiamiContentMvDetailGetRequest{
        Params: model.NewParams(),
    }
}

func (r XiamiContentMvDetailGetRequest) GetApiMethodName() string {
    return "xiami.content.mv.detail.get"
}

func (r XiamiContentMvDetailGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *XiamiContentMvDetailGetRequest) SetMvIds(mvIds []int64) error {
    r.mvIds = mvIds
    r.Set("mv_ids", mvIds)
    return nil
}

func (r XiamiContentMvDetailGetRequest) GetMvIds() []int64 {
    return r.mvIds
}

