package yunosappstore

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询HpPad appList API请求
yunos.appstore.pad.hp.applist

提供hp pad应用群数据
*/
type YunosAppstorePadHpApplistAPIRequest struct {
    model.Params
    // 获取的应用群code
    _code   string
}

// 初始化YunosAppstorePadHpApplistAPIRequest对象
func NewYunosAppstorePadHpApplistRequest() *YunosAppstorePadHpApplistAPIRequest{
    return &YunosAppstorePadHpApplistAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosAppstorePadHpApplistAPIRequest) GetApiMethodName() string {
    return "yunos.appstore.pad.hp.applist"
}

// IRequest interface 方法, 获取API参数
func (r YunosAppstorePadHpApplistAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Code Setter
// 获取的应用群code
func (r *YunosAppstorePadHpApplistAPIRequest) SetCode(_code string) error {
    r._code = _code
    r.Set("code", _code)
    return nil
}

// Code Getter
func (r YunosAppstorePadHpApplistAPIRequest) GetCode() string {
    return r._code
}
