package media

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除图片空间图片 APIRequest
taobao.picture.delete

删除图片空间图片
*/
type TaobaoPictureDeleteRequest struct {
    model.Params

    // 图片ID字符串,可以一个也可以一组,用英文逗号间隔,如450,120,155.限制数量是100
    pictureIds   []string 

}

func NewTaobaoPictureDeleteRequest() *TaobaoPictureDeleteRequest{
    return &TaobaoPictureDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoPictureDeleteRequest) GetApiMethodName() string {
    return "taobao.picture.delete"
}

func (r TaobaoPictureDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoPictureDeleteRequest) SetPictureIds(pictureIds []string) error {
    r.pictureIds = pictureIds
    r.Set("picture_ids", pictureIds)
    return nil
}

func (r TaobaoPictureDeleteRequest) GetPictureIds() []string {
    return r.pictureIds
}

