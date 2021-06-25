package tbk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
淘宝客-公用-淘口令生成 APIRequest
taobao.tbk.tpwd.create

提供淘口令生成接口。提交需要生成淘口令的内容、logo、url等参数，生成淘口令，其中关键信息为￥SADadW￥，后续可基于淘口令进行文案包装组装用于传播。
*/
type TaobaoTbkTpwdCreateRequest struct {
    model.Params

    // 联盟官方渠道获取的淘客推广链接，请注意，不要随意篡改官方生成的链接，否则可能无法生成淘口令
    url   string 

    // 兼容旧版本api参数，无实际作用
    text   string 

    // 兼容旧版本api参数，无实际作用
    logo   string 

    // 兼容旧版本api参数，无实际作用
    ext   string 

    // 兼容旧版本api参数，无实际作用
    userId   string 

}

func NewTaobaoTbkTpwdCreateRequest() *TaobaoTbkTpwdCreateRequest{
    return &TaobaoTbkTpwdCreateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTbkTpwdCreateRequest) GetApiMethodName() string {
    return "taobao.tbk.tpwd.create"
}

func (r TaobaoTbkTpwdCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTbkTpwdCreateRequest) SetUrl(url string) error {
    r.url = url
    r.Set("url", url)
    return nil
}

func (r TaobaoTbkTpwdCreateRequest) GetUrl() string {
    return r.url
}

func (r *TaobaoTbkTpwdCreateRequest) SetText(text string) error {
    r.text = text
    r.Set("text", text)
    return nil
}

func (r TaobaoTbkTpwdCreateRequest) GetText() string {
    return r.text
}

func (r *TaobaoTbkTpwdCreateRequest) SetLogo(logo string) error {
    r.logo = logo
    r.Set("logo", logo)
    return nil
}

func (r TaobaoTbkTpwdCreateRequest) GetLogo() string {
    return r.logo
}

func (r *TaobaoTbkTpwdCreateRequest) SetExt(ext string) error {
    r.ext = ext
    r.Set("ext", ext)
    return nil
}

func (r TaobaoTbkTpwdCreateRequest) GetExt() string {
    return r.ext
}

func (r *TaobaoTbkTpwdCreateRequest) SetUserId(userId string) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

func (r TaobaoTbkTpwdCreateRequest) GetUserId() string {
    return r.userId
}

