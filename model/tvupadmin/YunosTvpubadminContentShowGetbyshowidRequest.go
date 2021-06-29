package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
迎客松根据节目id获取节目元数据 API请求
yunos.tvpubadmin.content.show.getbyshowid

迎客松根据节目id获取节目元数据
*/
type YunosTvpubadminContentShowGetbyshowidRequest struct {
    model.Params
    // 节目字符串id
    showId   string
}

// 初始化YunosTvpubadminContentShowGetbyshowidRequest对象
func NewYunosTvpubadminContentShowGetbyshowidRequest() *YunosTvpubadminContentShowGetbyshowidRequest{
    return &YunosTvpubadminContentShowGetbyshowidRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminContentShowGetbyshowidRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.content.show.getbyshowid"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminContentShowGetbyshowidRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ShowId Setter
// 节目字符串id
func (r *YunosTvpubadminContentShowGetbyshowidRequest) SetShowId(showId string) error {
    r.showId = showId
    r.Set("show_id", showId)
    return nil
}

// ShowId Getter
func (r YunosTvpubadminContentShowGetbyshowidRequest) GetShowId() string {
    return r.showId
}
