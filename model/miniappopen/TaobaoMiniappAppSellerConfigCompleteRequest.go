package miniappopen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家完成小程序相关配置 APIRequest
taobao.miniapp.app.seller.config.complete

通过该接口告知平台商家已经完成小程序相关的必要设置，可进行后续操作。主要用于小部件、客服插件等场景。
*/
type TaobaoMiniappAppSellerConfigCompleteRequest struct {
    model.Params

    // 商家已完成配置的小部件/B端插件的appid
    appId   int64 

    // 小部件必传，B端插件不用传。与app_id对应的已完成配置的版本号
    version   string 

}

func NewTaobaoMiniappAppSellerConfigCompleteRequest() *TaobaoMiniappAppSellerConfigCompleteRequest{
    return &TaobaoMiniappAppSellerConfigCompleteRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoMiniappAppSellerConfigCompleteRequest) GetApiMethodName() string {
    return "taobao.miniapp.app.seller.config.complete"
}

func (r TaobaoMiniappAppSellerConfigCompleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoMiniappAppSellerConfigCompleteRequest) SetAppId(appId int64) error {
    r.appId = appId
    r.Set("app_id", appId)
    return nil
}

func (r TaobaoMiniappAppSellerConfigCompleteRequest) GetAppId() int64 {
    return r.appId
}

func (r *TaobaoMiniappAppSellerConfigCompleteRequest) SetVersion(version string) error {
    r.version = version
    r.Set("version", version)
    return nil
}

func (r TaobaoMiniappAppSellerConfigCompleteRequest) GetVersion() string {
    return r.version
}

