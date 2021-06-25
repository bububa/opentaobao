package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
（新）新建创意 APIRequest
taobao.simba.salestar.creative.add

创建一个创意
*/
type TaobaoSimbaSalestarCreativeAddRequest struct {
    model.Params

    // 推广组Id
    adgroupId   int64 

    // 创意标题，最多20个汉字
    title   string 

    // 创意图片地址，必须是推广组对应商品的图片之一
    imgUrl   string 

    // 主人昵称
    nick   string 

}

func NewTaobaoSimbaSalestarCreativeAddRequest() *TaobaoSimbaSalestarCreativeAddRequest{
    return &TaobaoSimbaSalestarCreativeAddRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSimbaSalestarCreativeAddRequest) GetApiMethodName() string {
    return "taobao.simba.salestar.creative.add"
}

func (r TaobaoSimbaSalestarCreativeAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSimbaSalestarCreativeAddRequest) SetAdgroupId(adgroupId int64) error {
    r.adgroupId = adgroupId
    r.Set("adgroup_id", adgroupId)
    return nil
}

func (r TaobaoSimbaSalestarCreativeAddRequest) GetAdgroupId() int64 {
    return r.adgroupId
}

func (r *TaobaoSimbaSalestarCreativeAddRequest) SetTitle(title string) error {
    r.title = title
    r.Set("title", title)
    return nil
}

func (r TaobaoSimbaSalestarCreativeAddRequest) GetTitle() string {
    return r.title
}

func (r *TaobaoSimbaSalestarCreativeAddRequest) SetImgUrl(imgUrl string) error {
    r.imgUrl = imgUrl
    r.Set("img_url", imgUrl)
    return nil
}

func (r TaobaoSimbaSalestarCreativeAddRequest) GetImgUrl() string {
    return r.imgUrl
}

func (r *TaobaoSimbaSalestarCreativeAddRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

func (r TaobaoSimbaSalestarCreativeAddRequest) GetNick() string {
    return r.nick
}

