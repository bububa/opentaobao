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
    pictureCategoryId   int64
    // 图片二进制文件流,不能为空,允许png、jpg、gif图片格式,3M以内
    img   []*model.File
    // 包括后缀名的图片标题,不能为空，如Bule.jpg,有些卖家希望图片上传后取图片文件的默认名
    imgInputTitle   string
    // 图片标题,如果为空,传的图片标题就取去掉后缀名的image_input_title,超过50字符长度会截取50字符,重名会在标题末尾加"(1)";标题末尾已经有"(数字)"了，则数字加1
    title   string
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
func (r *AlibabaWdkPictureUploadRequest) SetPictureCategoryId(pictureCategoryId int64) error {
    r.pictureCategoryId = pictureCategoryId
    r.Set("picture_category_id", pictureCategoryId)
    return nil
}

// PictureCategoryId Getter
func (r AlibabaWdkPictureUploadRequest) GetPictureCategoryId() int64 {
    return r.pictureCategoryId
}
// Img Setter
// 图片二进制文件流,不能为空,允许png、jpg、gif图片格式,3M以内
func (r *AlibabaWdkPictureUploadRequest) SetImg(img []*model.File) error {
    r.img = img
    r.Set("img", img)
    return nil
}

// Img Getter
func (r AlibabaWdkPictureUploadRequest) GetImg() []*model.File {
    return r.img
}
// ImgInputTitle Setter
// 包括后缀名的图片标题,不能为空，如Bule.jpg,有些卖家希望图片上传后取图片文件的默认名
func (r *AlibabaWdkPictureUploadRequest) SetImgInputTitle(imgInputTitle string) error {
    r.imgInputTitle = imgInputTitle
    r.Set("img_input_title", imgInputTitle)
    return nil
}

// ImgInputTitle Getter
func (r AlibabaWdkPictureUploadRequest) GetImgInputTitle() string {
    return r.imgInputTitle
}
// Title Setter
// 图片标题,如果为空,传的图片标题就取去掉后缀名的image_input_title,超过50字符长度会截取50字符,重名会在标题末尾加"(1)";标题末尾已经有"(数字)"了，则数字加1
func (r *AlibabaWdkPictureUploadRequest) SetTitle(title string) error {
    r.title = title
    r.Set("title", title)
    return nil
}

// Title Getter
func (r AlibabaWdkPictureUploadRequest) GetTitle() string {
    return r.title
}
