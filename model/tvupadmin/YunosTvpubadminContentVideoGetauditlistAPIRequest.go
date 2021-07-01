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
type YunosTvpubadminContentVideoGetauditlistAPIRequest struct {
    model.Params
    // 查询
    _query   string
}

// 初始化YunosTvpubadminContentVideoGetauditlistAPIRequest对象
func NewYunosTvpubadminContentVideoGetauditlistRequest() *YunosTvpubadminContentVideoGetauditlistAPIRequest{
    return &YunosTvpubadminContentVideoGetauditlistAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminContentVideoGetauditlistAPIRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.content.video.getauditlist"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminContentVideoGetauditlistAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Query Setter
// 查询
func (r *YunosTvpubadminContentVideoGetauditlistAPIRequest) SetQuery(_query string) error {
    r._query = _query
    r.Set("query", _query)
    return nil
}

// Query Getter
func (r YunosTvpubadminContentVideoGetauditlistAPIRequest) GetQuery() string {
    return r._query
}
