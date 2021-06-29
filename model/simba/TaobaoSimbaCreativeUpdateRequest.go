package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改创意与 API请求
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

// 初始化TaobaoSimbaCreativeUpdateRequest对象
func NewTaobaoSimbaCreativeUpdateRequest() *TaobaoSimbaCreativeUpdateRequest{
    return &TaobaoSimbaCreativeUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaCreativeUpdateRequest) GetApiMethodName() string {
    return "taobao.simba.creative.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaCreativeUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AdgroupId Setter
// 推广组Id
func (r *TaobaoSimbaCreativeUpdateRequest) SetAdgroupId(adgroupId int64) error {
    r.adgroupId = adgroupId
    r.Set("adgroup_id", adgroupId)
    return nil
}

// AdgroupId Getter
func (r TaobaoSimbaCreativeUpdateRequest) GetAdgroupId() int64 {
    return r.adgroupId
}
// CreativeId Setter
// 创意Id
func (r *TaobaoSimbaCreativeUpdateRequest) SetCreativeId(creativeId int64) error {
    r.creativeId = creativeId
    r.Set("creative_id", creativeId)
    return nil
}

// CreativeId Getter
func (r TaobaoSimbaCreativeUpdateRequest) GetCreativeId() int64 {
    return r.creativeId
}
// Title Setter
// 创意标题，最多20个汉字
func (r *TaobaoSimbaCreativeUpdateRequest) SetTitle(title string) error {
    r.title = title
    r.Set("title", title)
    return nil
}

// Title Getter
func (r TaobaoSimbaCreativeUpdateRequest) GetTitle() string {
    return r.title
}
// ImgUrl Setter
// 创意图片地址，必须是推广组对应商品的图片之一
func (r *TaobaoSimbaCreativeUpdateRequest) SetImgUrl(imgUrl string) error {
    r.imgUrl = imgUrl
    r.Set("img_url", imgUrl)
    return nil
}

// ImgUrl Getter
func (r TaobaoSimbaCreativeUpdateRequest) GetImgUrl() string {
    return r.imgUrl
}
// Nick Setter
// 主人昵称
func (r *TaobaoSimbaCreativeUpdateRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaCreativeUpdateRequest) GetNick() string {
    return r.nick
}
// PictureId Setter
// 如果用户开通了创意本地上传图片功能的，可以使用该用户图片空间的图片来修改创意，pictureId为图片空间中图片的pictureId，img_url为图片空间中图片链接地址，如果是使用的主图或副图修改创意，则pictureId必须为空
func (r *TaobaoSimbaCreativeUpdateRequest) SetPictureId(pictureId int64) error {
    r.pictureId = pictureId
    r.Set("picture_id", pictureId)
    return nil
}

// PictureId Getter
func (r TaobaoSimbaCreativeUpdateRequest) GetPictureId() int64 {
    return r.pictureId
}
