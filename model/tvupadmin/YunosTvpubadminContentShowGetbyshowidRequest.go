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
type YunosTvpubadminContentShowGetbyshowidAPIRequest struct {
    model.Params
    // 节目字符串id
    _showId   string
}

// 初始化YunosTvpubadminContentShowGetbyshowidAPIRequest对象
func NewYunosTvpubadminContentShowGetbyshowidRequest() *YunosTvpubadminContentShowGetbyshowidAPIRequest{
    return &YunosTvpubadminContentShowGetbyshowidAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminContentShowGetbyshowidAPIRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.content.show.getbyshowid"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminContentShowGetbyshowidAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ShowId Setter
// 节目字符串id
func (r *YunosTvpubadminContentShowGetbyshowidAPIRequest) SetShowId(_showId string) error {
    r._showId = _showId
    r.Set("show_id", _showId)
    return nil
}

// ShowId Getter
func (r YunosTvpubadminContentShowGetbyshowidAPIRequest) GetShowId() string {
    return r._showId
}
