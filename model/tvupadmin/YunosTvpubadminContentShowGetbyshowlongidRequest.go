package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
迎客松根据节目longid获取节目元数据 API请求
yunos.tvpubadmin.content.show.getbyshowlongid

迎客松根据节目longid获取节目元数据
*/
type YunosTvpubadminContentShowGetbyshowlongidAPIRequest struct {
    model.Params
    // 节目longid
    _showLongId   int64
}

// 初始化YunosTvpubadminContentShowGetbyshowlongidAPIRequest对象
func NewYunosTvpubadminContentShowGetbyshowlongidRequest() *YunosTvpubadminContentShowGetbyshowlongidAPIRequest{
    return &YunosTvpubadminContentShowGetbyshowlongidAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminContentShowGetbyshowlongidAPIRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.content.show.getbyshowlongid"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminContentShowGetbyshowlongidAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ShowLongId Setter
// 节目longid
func (r *YunosTvpubadminContentShowGetbyshowlongidAPIRequest) SetShowLongId(_showLongId int64) error {
    r._showLongId = _showLongId
    r.Set("show_long_id", _showLongId)
    return nil
}

// ShowLongId Getter
func (r YunosTvpubadminContentShowGetbyshowlongidAPIRequest) GetShowLongId() int64 {
    return r._showLongId
}
