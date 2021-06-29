package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
应用上下架操作 APIRequest
yunos.tvpubadmin.content.app.onoffappbylicense

应用上下架操作
*/
type YunosTvpubadminContentAppOnoffappbylicenseRequest struct {
    model.Params

    // com.ali.yunos.tvacs.domain.OnOffApp
    onOffApp   string 

}

func NewYunosTvpubadminContentAppOnoffappbylicenseRequest() *YunosTvpubadminContentAppOnoffappbylicenseRequest{
    return &YunosTvpubadminContentAppOnoffappbylicenseRequest{
        Params: model.NewParams(),
    }
}

func (r YunosTvpubadminContentAppOnoffappbylicenseRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.content.app.onoffappbylicense"
}

func (r YunosTvpubadminContentAppOnoffappbylicenseRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosTvpubadminContentAppOnoffappbylicenseRequest) SetOnOffApp(onOffApp string) error {
    r.onOffApp = onOffApp
    r.Set("on_off_app", onOffApp)
    return nil
}

func (r YunosTvpubadminContentAppOnoffappbylicenseRequest) GetOnOffApp() string {
    return r.onOffApp
}

