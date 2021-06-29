package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
（新）新建创意 API请求
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

// 初始化TaobaoSimbaSalestarCreativeAddRequest对象
func NewTaobaoSimbaSalestarCreativeAddRequest() *TaobaoSimbaSalestarCreativeAddRequest{
    return &TaobaoSimbaSalestarCreativeAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaSalestarCreativeAddRequest) GetApiMethodName() string {
    return "taobao.simba.salestar.creative.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaSalestarCreativeAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AdgroupId Setter
// 推广组Id
func (r *TaobaoSimbaSalestarCreativeAddRequest) SetAdgroupId(adgroupId int64) error {
    r.adgroupId = adgroupId
    r.Set("adgroup_id", adgroupId)
    return nil
}

// AdgroupId Getter
func (r TaobaoSimbaSalestarCreativeAddRequest) GetAdgroupId() int64 {
    return r.adgroupId
}
// Title Setter
// 创意标题，最多20个汉字
func (r *TaobaoSimbaSalestarCreativeAddRequest) SetTitle(title string) error {
    r.title = title
    r.Set("title", title)
    return nil
}

// Title Getter
func (r TaobaoSimbaSalestarCreativeAddRequest) GetTitle() string {
    return r.title
}
// ImgUrl Setter
// 创意图片地址，必须是推广组对应商品的图片之一
func (r *TaobaoSimbaSalestarCreativeAddRequest) SetImgUrl(imgUrl string) error {
    r.imgUrl = imgUrl
    r.Set("img_url", imgUrl)
    return nil
}

// ImgUrl Getter
func (r TaobaoSimbaSalestarCreativeAddRequest) GetImgUrl() string {
    return r.imgUrl
}
// Nick Setter
// 主人昵称
func (r *TaobaoSimbaSalestarCreativeAddRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaSalestarCreativeAddRequest) GetNick() string {
    return r.nick
}
