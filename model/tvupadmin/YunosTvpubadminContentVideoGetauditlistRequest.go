package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
迎客松视频审核记录查询 API请求
yunos.tvpubadmin.content.video.getauditlist

迎客松视频审核记录查询
*/
type YunosTvpubadminContentVideoGetauditlistRequest struct {
    model.Params
    // 查询
    query   string
}

// 初始化YunosTvpubadminContentVideoGetauditlistRequest对象
func NewYunosTvpubadminContentVideoGetauditlistRequest() *YunosTvpubadminContentVideoGetauditlistRequest{
    return &YunosTvpubadminContentVideoGetauditlistRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminContentVideoGetauditlistRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.content.video.getauditlist"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminContentVideoGetauditlistRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Query Setter
// 查询
func (r *YunosTvpubadminContentVideoGetauditlistRequest) SetQuery(query string) error {
    r.query = query
    r.Set("query", query)
    return nil
}

// Query Getter
func (r YunosTvpubadminContentVideoGetauditlistRequest) GetQuery() string {
    return r.query
}
