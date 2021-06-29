package xiamicontent

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取mv详情 API请求
xiami.content.mv.detail.get

获取mv详情
*/
type XiamiContentMvDetailGetRequest struct {
    model.Params
    // mvId
    mvIds   []int64
}

// 初始化XiamiContentMvDetailGetRequest对象
func NewXiamiContentMvDetailGetRequest() *XiamiContentMvDetailGetRequest{
    return &XiamiContentMvDetailGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r XiamiContentMvDetailGetRequest) GetApiMethodName() string {
    return "xiami.content.mv.detail.get"
}

// IRequest interface 方法, 获取API参数
func (r XiamiContentMvDetailGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MvIds Setter
// mvId
func (r *XiamiContentMvDetailGetRequest) SetMvIds(mvIds []int64) error {
    r.mvIds = mvIds
    r.Set("mv_ids", mvIds)
    return nil
}

// MvIds Getter
func (r XiamiContentMvDetailGetRequest) GetMvIds() []int64 {
    return r.mvIds
}
