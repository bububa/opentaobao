package media

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
替换图片 APIRequest
taobao.picture.replace

替换一张图片，只替换图片数据，图片名称，图片分类等保持不变。
*/
type TaobaoPictureReplaceRequest struct {
    model.Params

    // 要替换的图片的id，必须大于0
    pictureId   int64 

    // 图片二进制文件流,不能为空,允许png、jpg、gif图片格式
    imageData   []*model.File 

}

func NewTaobaoPictureReplaceRequest() *TaobaoPictureReplaceRequest{
    return &TaobaoPictureReplaceRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoPictureReplaceRequest) GetApiMethodName() string {
    return "taobao.picture.replace"
}

func (r TaobaoPictureReplaceRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoPictureReplaceRequest) SetPictureId(pictureId int64) error {
    r.pictureId = pictureId
    r.Set("picture_id", pictureId)
    return nil
}

func (r TaobaoPictureReplaceRequest) GetPictureId() int64 {
    return r.pictureId
}

func (r *TaobaoPictureReplaceRequest) SetImageData(imageData []*model.File) error {
    r.imageData = imageData
    r.Set("image_data", imageData)
    return nil
}

func (r TaobaoPictureReplaceRequest) GetImageData() []*model.File {
    return r.imageData
}

