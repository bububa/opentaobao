package tbk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
淘宝客-公用-淘口令生成 API请求
taobao.tbk.tpwd.create

提供淘口令生成接口。提交需要生成淘口令的内容、logo、url等参数，生成淘口令，其中关键信息为￥SADadW￥，后续可基于淘口令进行文案包装组装用于传播。
*/
type TaobaoTbkTpwdCreateRequest struct {
    model.Params
    // 联盟官方渠道获取的淘客推广链接，请注意，不要随意篡改官方生成的链接，否则可能无法生成淘口令
    _url   string
    // 兼容旧版本api参数，无实际作用
    _text   string
    // 兼容旧版本api参数，无实际作用
    _logo   string
    // 兼容旧版本api参数，无实际作用
    _ext   string
    // 兼容旧版本api参数，无实际作用
    _userId   string
}

// 初始化TaobaoTbkTpwdCreateRequest对象
func NewTaobaoTbkTpwdCreateRequest() *TaobaoTbkTpwdCreateRequest{
    return &TaobaoTbkTpwdCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTbkTpwdCreateRequest) GetApiMethodName() string {
    return "taobao.tbk.tpwd.create"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTbkTpwdCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Url Setter
// 联盟官方渠道获取的淘客推广链接，请注意，不要随意篡改官方生成的链接，否则可能无法生成淘口令
func (r *TaobaoTbkTpwdCreateRequest) SetUrl(_url string) error {
    r._url = _url
    r.Set("url", _url)
    return nil
}

// Url Getter
func (r TaobaoTbkTpwdCreateRequest) GetUrl() string {
    return r._url
}
// Text Setter
// 兼容旧版本api参数，无实际作用
func (r *TaobaoTbkTpwdCreateRequest) SetText(_text string) error {
    r._text = _text
    r.Set("text", _text)
    return nil
}

// Text Getter
func (r TaobaoTbkTpwdCreateRequest) GetText() string {
    return r._text
}
// Logo Setter
// 兼容旧版本api参数，无实际作用
func (r *TaobaoTbkTpwdCreateRequest) SetLogo(_logo string) error {
    r._logo = _logo
    r.Set("logo", _logo)
    return nil
}

// Logo Getter
func (r TaobaoTbkTpwdCreateRequest) GetLogo() string {
    return r._logo
}
// Ext Setter
// 兼容旧版本api参数，无实际作用
func (r *TaobaoTbkTpwdCreateRequest) SetExt(_ext string) error {
    r._ext = _ext
    r.Set("ext", _ext)
    return nil
}

// Ext Getter
func (r TaobaoTbkTpwdCreateRequest) GetExt() string {
    return r._ext
}
// UserId Setter
// 兼容旧版本api参数，无实际作用
func (r *TaobaoTbkTpwdCreateRequest) SetUserId(_userId string) error {
    r._userId = _userId
    r.Set("user_id", _userId)
    return nil
}

// UserId Getter
func (r TaobaoTbkTpwdCreateRequest) GetUserId() string {
    return r._userId
}
