package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改创意与 APIRequest
taobao.simba.creative.update

更新一个创意的信息，可以设置创意标题、创意图片
*/
type TaobaoSimbaCreativeUpdateRequest struct {
    model.Params

    // 推广组Id
    adgroupId   int64 

    // 创意Id
    creativeId   int64 

    // 创意标题，最多20个汉字
    title   string 

    // 创意图片地址，必须是推广组对应商品的图片之一
    imgUrl   string 

    // 主人昵称
    nick   string 

    // 如果用户开通了创意本地上传图片功能的，可以使用该用户图片空间的图片来修改创意，pictureId为图片空间中图片的pictureId，img_url为图片空间中图片链接地址，如果是使用的主图或副图修改创意，则pictureId必须为空
    pictureId   int64 

}

func NewTaobaoSimbaCreativeUpdateRequest() *TaobaoSimbaCreativeUpdateRequest{
    return &TaobaoSimbaCreativeUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSimbaCreativeUpdateRequest) GetApiMethodName() string {
    return "taobao.simba.creative.update"
}

func (r TaobaoSimbaCreativeUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSimbaCreativeUpdateRequest) SetAdgroupId(adgroupId int64) error {
    r.adgroupId = adgroupId
    r.Set("adgroup_id", adgroupId)
    return nil
}

func (r TaobaoSimbaCreativeUpdateRequest) GetAdgroupId() int64 {
    return r.adgroupId
}

func (r *TaobaoSimbaCreativeUpdateRequest) SetCreativeId(creativeId int64) error {
    r.creativeId = creativeId
    r.Set("creative_id", creativeId)
    return nil
}

func (r TaobaoSimbaCreativeUpdateRequest) GetCreativeId() int64 {
    return r.creativeId
}

func (r *TaobaoSimbaCreativeUpdateRequest) SetTitle(title string) error {
    r.title = title
    r.Set("title", title)
    return nil
}

func (r TaobaoSimbaCreativeUpdateRequest) GetTitle() string {
    return r.title
}

func (r *TaobaoSimbaCreativeUpdateRequest) SetImgUrl(imgUrl string) error {
    r.imgUrl = imgUrl
    r.Set("img_url", imgUrl)
    return nil
}

func (r TaobaoSimbaCreativeUpdateRequest) GetImgUrl() string {
    return r.imgUrl
}

func (r *TaobaoSimbaCreativeUpdateRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

func (r TaobaoSimbaCreativeUpdateRequest) GetNick() string {
    return r.nick
}

func (r *TaobaoSimbaCreativeUpdateRequest) SetPictureId(pictureId int64) error {
    r.pictureId = pictureId
    r.Set("picture_id", pictureId)
    return nil
}

func (r TaobaoSimbaCreativeUpdateRequest) GetPictureId() int64 {
    return r.pictureId
}

