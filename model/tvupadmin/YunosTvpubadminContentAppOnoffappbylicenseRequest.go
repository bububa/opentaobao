package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
应用上下架操作 API请求
yunos.tvpubadmin.content.app.onoffappbylicense

应用上下架操作
*/
type YunosTvpubadminContentAppOnoffappbylicenseRequest struct {
    model.Params
    // com.ali.yunos.tvacs.domain.OnOffApp
    _onOffApp   string
}

// 初始化YunosTvpubadminContentAppOnoffappbylicenseRequest对象
func NewYunosTvpubadminContentAppOnoffappbylicenseRequest() *YunosTvpubadminContentAppOnoffappbylicenseRequest{
    return &YunosTvpubadminContentAppOnoffappbylicenseRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminContentAppOnoffappbylicenseRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.content.app.onoffappbylicense"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminContentAppOnoffappbylicenseRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OnOffApp Setter
// com.ali.yunos.tvacs.domain.OnOffApp
func (r *YunosTvpubadminContentAppOnoffappbylicenseRequest) SetOnOffApp(_onOffApp string) error {
    r._onOffApp = _onOffApp
    r.Set("on_off_app", _onOffApp)
    return nil
}

// OnOffApp Getter
func (r YunosTvpubadminContentAppOnoffappbylicenseRequest) GetOnOffApp() string {
    return r._onOffApp
}
