package yunosappstore

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询HpPad appList APIRequest
yunos.appstore.pad.hp.applist

提供hp pad应用群数据
*/
type YunosAppstorePadHpApplistRequest struct {
    model.Params

    // 获取的应用群code
    code   string 

}

func NewYunosAppstorePadHpApplistRequest() *YunosAppstorePadHpApplistRequest{
    return &YunosAppstorePadHpApplistRequest{
        Params: model.NewParams(),
    }
}

func (r YunosAppstorePadHpApplistRequest) GetApiMethodName() string {
    return "yunos.appstore.pad.hp.applist"
}

func (r YunosAppstorePadHpApplistRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosAppstorePadHpApplistRequest) SetCode(code string) error {
    r.code = code
    r.Set("code", code)
    return nil
}

func (r YunosAppstorePadHpApplistRequest) GetCode() string {
    return r.code
}

