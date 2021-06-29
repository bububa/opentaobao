package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
迎客松批量查询节目某个牌照的免审状态 API请求
yunos.tvpubadmin.content.show.getshowexemptauditmap

迎客松批量查询节目某个牌照的免审状态
*/
type YunosTvpubadminContentShowGetshowexemptauditmapRequest struct {
    model.Params
    // 节目longid
    _showLongIds   string
    // 牌照id
    _license   int64
}

// 初始化YunosTvpubadminContentShowGetshowexemptauditmapRequest对象
func NewYunosTvpubadminContentShowGetshowexemptauditmapRequest() *YunosTvpubadminContentShowGetshowexemptauditmapRequest{
    return &YunosTvpubadminContentShowGetshowexemptauditmapRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminContentShowGetshowexemptauditmapRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.content.show.getshowexemptauditmap"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminContentShowGetshowexemptauditmapRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ShowLongIds Setter
// 节目longid
func (r *YunosTvpubadminContentShowGetshowexemptauditmapRequest) SetShowLongIds(_showLongIds string) error {
    r._showLongIds = _showLongIds
    r.Set("show_long_ids", _showLongIds)
    return nil
}

// ShowLongIds Getter
func (r YunosTvpubadminContentShowGetshowexemptauditmapRequest) GetShowLongIds() string {
    return r._showLongIds
}
// License Setter
// 牌照id
func (r *YunosTvpubadminContentShowGetshowexemptauditmapRequest) SetLicense(_license int64) error {
    r._license = _license
    r.Set("license", _license)
    return nil
}

// License Getter
func (r YunosTvpubadminContentShowGetshowexemptauditmapRequest) GetLicense() int64 {
    return r._license
}
