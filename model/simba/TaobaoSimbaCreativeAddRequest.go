package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
增加创意 APIRequest
taobao.simba.creative.add

创建一个创意
*/
type TaobaoSimbaCreativeAddRequest struct {
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

func NewTaobaoSimbaCreativeAddRequest() *TaobaoSimbaCreativeAddRequest{
    return &TaobaoSimbaCreativeAddRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSimbaCreativeAddRequest) GetApiMethodName() string {
    return "taobao.simba.creative.add"
}

func (r TaobaoSimbaCreativeAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSimbaCreativeAddRequest) SetAdgroupId(adgroupId int64) error {
    r.adgroupId = adgroupId
    r.Set("adgroup_id", adgroupId)
    return nil
}

func (r TaobaoSimbaCreativeAddRequest) GetAdgroupId() int64 {
    return r.adgroupId
}

func (r *TaobaoSimbaCreativeAddRequest) SetTitle(title string) error {
    r.title = title
    r.Set("title", title)
    return nil
}

func (r TaobaoSimbaCreativeAddRequest) GetTitle() string {
    return r.title
}

func (r *TaobaoSimbaCreativeAddRequest) SetImgUrl(imgUrl string) error {
    r.imgUrl = imgUrl
    r.Set("img_url", imgUrl)
    return nil
}

func (r TaobaoSimbaCreativeAddRequest) GetImgUrl() string {
    return r.imgUrl
}

func (r *TaobaoSimbaCreativeAddRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

func (r TaobaoSimbaCreativeAddRequest) GetNick() string {
    return r.nick
}

