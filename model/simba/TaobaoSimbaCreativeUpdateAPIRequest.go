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
type TaobaoSimbaCreativeUpdateAPIRequest struct {
    model.Params
    // 推广组Id
    _adgroupId   int64
    // 创意Id
    _creativeId   int64
    // 创意标题，最多20个汉字
    _title   string
    // 创意图片地址，必须是推广组对应商品的图片之一
    _imgUrl   string
    // 主人昵称
    _nick   string
    // 如果用户开通了创意本地上传图片功能的，可以使用该用户图片空间的图片来修改创意，pictureId为图片空间中图片的pictureId，img_url为图片空间中图片链接地址，如果是使用的主图或副图修改创意，则pictureId必须为空
    _pictureId   int64
}

// 初始化TaobaoSimbaCreativeUpdateAPIRequest对象
func NewTaobaoSimbaCreativeUpdateRequest() *TaobaoSimbaCreativeUpdateAPIRequest{
    return &TaobaoSimbaCreativeUpdateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaCreativeUpdateAPIRequest) GetApiMethodName() string {
    return "taobao.simba.creative.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaCreativeUpdateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AdgroupId Setter
// 推广组Id
func (r *TaobaoSimbaCreativeUpdateAPIRequest) SetAdgroupId(_adgroupId int64) error {
    r._adgroupId = _adgroupId
    r.Set("adgroup_id", _adgroupId)
    return nil
}

// AdgroupId Getter
func (r TaobaoSimbaCreativeUpdateAPIRequest) GetAdgroupId() int64 {
    return r._adgroupId
}
// CreativeId Setter
// 创意Id
func (r *TaobaoSimbaCreativeUpdateAPIRequest) SetCreativeId(_creativeId int64) error {
    r._creativeId = _creativeId
    r.Set("creative_id", _creativeId)
    return nil
}

// CreativeId Getter
func (r TaobaoSimbaCreativeUpdateAPIRequest) GetCreativeId() int64 {
    return r._creativeId
}
// Title Setter
// 创意标题，最多20个汉字
func (r *TaobaoSimbaCreativeUpdateAPIRequest) SetTitle(_title string) error {
    r._title = _title
    r.Set("title", _title)
    return nil
}

// Title Getter
func (r TaobaoSimbaCreativeUpdateAPIRequest) GetTitle() string {
    return r._title
}
// ImgUrl Setter
// 创意图片地址，必须是推广组对应商品的图片之一
func (r *TaobaoSimbaCreativeUpdateAPIRequest) SetImgUrl(_imgUrl string) error {
    r._imgUrl = _imgUrl
    r.Set("img_url", _imgUrl)
    return nil
}

// ImgUrl Getter
func (r TaobaoSimbaCreativeUpdateAPIRequest) GetImgUrl() string {
    return r._imgUrl
}
// Nick Setter
// 主人昵称
func (r *TaobaoSimbaCreativeUpdateAPIRequest) SetNick(_nick string) error {
    r._nick = _nick
    r.Set("nick", _nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaCreativeUpdateAPIRequest) GetNick() string {
    return r._nick
}
// PictureId Setter
// 如果用户开通了创意本地上传图片功能的，可以使用该用户图片空间的图片来修改创意，pictureId为图片空间中图片的pictureId，img_url为图片空间中图片链接地址，如果是使用的主图或副图修改创意，则pictureId必须为空
func (r *TaobaoSimbaCreativeUpdateAPIRequest) SetPictureId(_pictureId int64) error {
    r._pictureId = _pictureId
    r.Set("picture_id", _pictureId)
    return nil
}

// PictureId Getter
func (r TaobaoSimbaCreativeUpdateAPIRequest) GetPictureId() int64 {
    return r._pictureId
}
