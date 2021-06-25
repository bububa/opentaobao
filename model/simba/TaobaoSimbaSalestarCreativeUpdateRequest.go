package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
销量明星更新创意相关接口 APIRequest
taobao.simba.salestar.creative.update

更新一个创意的信息，可以设置创意标题、创意图片
*/
type TaobaoSimbaSalestarCreativeUpdateRequest struct {
    model.Params

    // 推广组Id
    adgroupId   int64 

    // 创意Id
    creativeId   int64 

    // 创意标题，最多20个汉字
    title   string 

    // 创意图片地址，必须是推广组对应商品的图片之一
    imgUrl   string 

    // 如果用户开通了创意本地上传图片功能的，可以使用该用户图片空间的图片来修改创意，pictureId为图片空间中图片的pictureId，img_url为图片空间中图片链接地址，如果是使用的主图或副图修改创意，则pictureId必须为空
    pictureId   int64 

}

func NewTaobaoSimbaSalestarCreativeUpdateRequest() *TaobaoSimbaSalestarCreativeUpdateRequest{
    return &TaobaoSimbaSalestarCreativeUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSimbaSalestarCreativeUpdateRequest) GetApiMethodName() string {
    return "taobao.simba.salestar.creative.update"
}

func (r TaobaoSimbaSalestarCreativeUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSimbaSalestarCreativeUpdateRequest) SetAdgroupId(adgroupId int64) error {
    r.adgroupId = adgroupId
    r.Set("adgroup_id", adgroupId)
    return nil
}

func (r TaobaoSimbaSalestarCreativeUpdateRequest) GetAdgroupId() int64 {
    return r.adgroupId
}

func (r *TaobaoSimbaSalestarCreativeUpdateRequest) SetCreativeId(creativeId int64) error {
    r.creativeId = creativeId
    r.Set("creative_id", creativeId)
    return nil
}

func (r TaobaoSimbaSalestarCreativeUpdateRequest) GetCreativeId() int64 {
    return r.creativeId
}

func (r *TaobaoSimbaSalestarCreativeUpdateRequest) SetTitle(title string) error {
    r.title = title
    r.Set("title", title)
    return nil
}

func (r TaobaoSimbaSalestarCreativeUpdateRequest) GetTitle() string {
    return r.title
}

func (r *TaobaoSimbaSalestarCreativeUpdateRequest) SetImgUrl(imgUrl string) error {
    r.imgUrl = imgUrl
    r.Set("img_url", imgUrl)
    return nil
}

func (r TaobaoSimbaSalestarCreativeUpdateRequest) GetImgUrl() string {
    return r.imgUrl
}

func (r *TaobaoSimbaSalestarCreativeUpdateRequest) SetPictureId(pictureId int64) error {
    r.pictureId = pictureId
    r.Set("picture_id", pictureId)
    return nil
}

func (r TaobaoSimbaSalestarCreativeUpdateRequest) GetPictureId() int64 {
    return r.pictureId
}

