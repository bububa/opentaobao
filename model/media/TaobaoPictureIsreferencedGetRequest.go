package media

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
图片是否被引用 APIRequest
taobao.picture.isreferenced.get

查询图片是否被引用，被引用返回true，未被引用返回false
*/
type TaobaoPictureIsreferencedGetRequest struct {
    model.Params

    // 图片id
    pictureId   int64 

}

func NewTaobaoPictureIsreferencedGetRequest() *TaobaoPictureIsreferencedGetRequest{
    return &TaobaoPictureIsreferencedGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoPictureIsreferencedGetRequest) GetApiMethodName() string {
    return "taobao.picture.isreferenced.get"
}

func (r TaobaoPictureIsreferencedGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoPictureIsreferencedGetRequest) SetPictureId(pictureId int64) error {
    r.pictureId = pictureId
    r.Set("picture_id", pictureId)
    return nil
}

func (r TaobaoPictureIsreferencedGetRequest) GetPictureId() int64 {
    return r.pictureId
}

