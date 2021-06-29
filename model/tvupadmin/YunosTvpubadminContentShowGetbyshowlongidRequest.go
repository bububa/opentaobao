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
type YunosTvpubadminContentShowGetbyshowlongidRequest struct {
    model.Params
    // 节目longid
    showLongId   int64
}

// 初始化YunosTvpubadminContentShowGetbyshowlongidRequest对象
func NewYunosTvpubadminContentShowGetbyshowlongidRequest() *YunosTvpubadminContentShowGetbyshowlongidRequest{
    return &YunosTvpubadminContentShowGetbyshowlongidRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminContentShowGetbyshowlongidRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.content.show.getbyshowlongid"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminContentShowGetbyshowlongidRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ShowLongId Setter
// 节目longid
func (r *YunosTvpubadminContentShowGetbyshowlongidRequest) SetShowLongId(showLongId int64) error {
    r.showLongId = showLongId
    r.Set("show_long_id", showLongId)
    return nil
}

// ShowLongId Getter
func (r YunosTvpubadminContentShowGetbyshowlongidRequest) GetShowLongId() int64 {
    return r.showLongId
}
