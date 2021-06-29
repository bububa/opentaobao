package wdkitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
图片上传接口 API请求
alibaba.wdk.picture.upload

上传图片
*/
type AlibabaWdkPictureUploadRequest struct {
    model.Params
    // 图片分类ID，设置具体某个分类ID或设置0上传到默认分类，只能传入一个分类
    _pictureCategoryId   int64
    // 图片二进制文件流,不能为空,允许png、jpg、gif图片格式,3M以内
    _img   *model.File
    // 包括后缀名的图片标题,不能为空，如Bule.jpg,有些卖家希望图片上传后取图片文件的默认名
    _imgInputTitle   string
    // 图片标题,如果为空,传的图片标题就取去掉后缀名的image_input_title,超过50字符长度会截取50字符,重名会在标题末尾加"(1)";标题末尾已经有"(数字)"了，则数字加1
    _title   string
}

// 初始化AlibabaWdkPictureUploadRequest对象
func NewAlibabaWdkPictureUploadRequest() *AlibabaWdkPictureUploadRequest{
    return &AlibabaWdkPictureUploadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkPictureUploadRequest) GetApiMethodName() string {
    return "alibaba.wdk.picture.upload"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkPictureUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PictureCategoryId Setter
// 图片分类ID，设置具体某个分类ID或设置0上传到默认分类，只能传入一个分类
func (r *AlibabaWdkPictureUploadRequest) SetPictureCategoryId(_pictureCategoryId int64) error {
    r._pictureCategoryId = _pictureCategoryId
    r.Set("picture_category_id", _pictureCategoryId)
    return nil
}

// PictureCategoryId Getter
func (r AlibabaWdkPictureUploadRequest) GetPictureCategoryId() int64 {
    return r._pictureCategoryId
}
// Img Setter
// 图片二进制文件流,不能为空,允许png、jpg、gif图片格式,3M以内
func (r *AlibabaWdkPictureUploadRequest) SetImg(_img *model.File) error {
    r._img = _img
    r.Set("img", _img)
    return nil
}

// Img Getter
func (r AlibabaWdkPictureUploadRequest) GetImg() *model.File {
    return r._img
}
// ImgInputTitle Setter
// 包括后缀名的图片标题,不能为空，如Bule.jpg,有些卖家希望图片上传后取图片文件的默认名
func (r *AlibabaWdkPictureUploadRequest) SetImgInputTitle(_imgInputTitle string) error {
    r._imgInputTitle = _imgInputTitle
    r.Set("img_input_title", _imgInputTitle)
    return nil
}

// ImgInputTitle Getter
func (r AlibabaWdkPictureUploadRequest) GetImgInputTitle() string {
    return r._imgInputTitle
}
// Title Setter
// 图片标题,如果为空,传的图片标题就取去掉后缀名的image_input_title,超过50字符长度会截取50字符,重名会在标题末尾加"(1)";标题末尾已经有"(数字)"了，则数字加1
func (r *AlibabaWdkPictureUploadRequest) SetTitle(_title string) error {
    r._title = _title
    r.Set("title", _title)
    return nil
}

// Title Getter
func (r AlibabaWdkPictureUploadRequest) GetTitle() string {
    return r._title
}
