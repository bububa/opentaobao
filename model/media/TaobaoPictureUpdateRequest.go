package media

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改图片名字 APIRequest
taobao.picture.update

修改指定图片的图片名
*/
type TaobaoPictureUpdateRequest struct {
    model.Params

    // 要更改名字的图片的id
    pictureId   int64 

    // 新的图片名，最大长度50字符，不能为空
    newName   string 

}

func NewTaobaoPictureUpdateRequest() *TaobaoPictureUpdateRequest{
    return &TaobaoPictureUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoPictureUpdateRequest) GetApiMethodName() string {
    return "taobao.picture.update"
}

func (r TaobaoPictureUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoPictureUpdateRequest) SetPictureId(pictureId int64) error {
    r.pictureId = pictureId
    r.Set("picture_id", pictureId)
    return nil
}

func (r TaobaoPictureUpdateRequest) GetPictureId() int64 {
    return r.pictureId
}

func (r *TaobaoPictureUpdateRequest) SetNewName(newName string) error {
    r.newName = newName
    r.Set("new_name", newName)
    return nil
}

func (r TaobaoPictureUpdateRequest) GetNewName() string {
    return r.newName
}

