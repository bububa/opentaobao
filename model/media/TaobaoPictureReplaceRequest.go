package media

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
替换图片 API请求
taobao.picture.replace

替换一张图片，只替换图片数据，图片名称，图片分类等保持不变。
*/
type TaobaoPictureReplaceRequest struct {
    model.Params
    // 要替换的图片的id，必须大于0
    _pictureId   int64
    // 图片二进制文件流,不能为空,允许png、jpg、gif图片格式
    _imageData   *model.File
}

// 初始化TaobaoPictureReplaceRequest对象
func NewTaobaoPictureReplaceRequest() *TaobaoPictureReplaceRequest{
    return &TaobaoPictureReplaceRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoPictureReplaceRequest) GetApiMethodName() string {
    return "taobao.picture.replace"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoPictureReplaceRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PictureId Setter
// 要替换的图片的id，必须大于0
func (r *TaobaoPictureReplaceRequest) SetPictureId(_pictureId int64) error {
    r._pictureId = _pictureId
    r.Set("picture_id", _pictureId)
    return nil
}

// PictureId Getter
func (r TaobaoPictureReplaceRequest) GetPictureId() int64 {
    return r._pictureId
}
// ImageData Setter
// 图片二进制文件流,不能为空,允许png、jpg、gif图片格式
func (r *TaobaoPictureReplaceRequest) SetImageData(_imageData *model.File) error {
    r._imageData = _imageData
    r.Set("image_data", _imageData)
    return nil
}

// ImageData Getter
func (r TaobaoPictureReplaceRequest) GetImageData() *model.File {
    return r._imageData
}
