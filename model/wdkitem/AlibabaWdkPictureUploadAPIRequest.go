package wdkitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkPictureUploadAPIRequest 图片上传接口 API请求
// alibaba.wdk.picture.upload
//
// 上传图片
type AlibabaWdkPictureUploadAPIRequest struct {
	model.Params
	// 包括后缀名的图片标题,不能为空，如Bule.jpg,有些卖家希望图片上传后取图片文件的默认名
	_imgInputTitle string
	// 图片标题,如果为空,传的图片标题就取去掉后缀名的image_input_title,超过50字符长度会截取50字符,重名会在标题末尾加"(1)";标题末尾已经有"(数字)"了，则数字加1
	_title string
	// 图片分类ID，设置具体某个分类ID或设置0上传到默认分类，只能传入一个分类
	_pictureCategoryId int64
	// 图片二进制文件流,不能为空,允许png、jpg、gif图片格式,3M以内
	_img *model.File
}

// NewAlibabaWdkPictureUploadRequest 初始化AlibabaWdkPictureUploadAPIRequest对象
func NewAlibabaWdkPictureUploadRequest() *AlibabaWdkPictureUploadAPIRequest {
	return &AlibabaWdkPictureUploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkPictureUploadAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.picture.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkPictureUploadAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetImgInputTitle is ImgInputTitle Setter
// 包括后缀名的图片标题,不能为空，如Bule.jpg,有些卖家希望图片上传后取图片文件的默认名
func (r *AlibabaWdkPictureUploadAPIRequest) SetImgInputTitle(_imgInputTitle string) error {
	r._imgInputTitle = _imgInputTitle
	r.Set("img_input_title", _imgInputTitle)
	return nil
}

// GetImgInputTitle ImgInputTitle Getter
func (r AlibabaWdkPictureUploadAPIRequest) GetImgInputTitle() string {
	return r._imgInputTitle
}

// SetTitle is Title Setter
// 图片标题,如果为空,传的图片标题就取去掉后缀名的image_input_title,超过50字符长度会截取50字符,重名会在标题末尾加&#34;(1)&#34;;标题末尾已经有&#34;(数字)&#34;了，则数字加1
func (r *AlibabaWdkPictureUploadAPIRequest) SetTitle(_title string) error {
	r._title = _title
	r.Set("title", _title)
	return nil
}

// GetTitle Title Getter
func (r AlibabaWdkPictureUploadAPIRequest) GetTitle() string {
	return r._title
}

// SetPictureCategoryId is PictureCategoryId Setter
// 图片分类ID，设置具体某个分类ID或设置0上传到默认分类，只能传入一个分类
func (r *AlibabaWdkPictureUploadAPIRequest) SetPictureCategoryId(_pictureCategoryId int64) error {
	r._pictureCategoryId = _pictureCategoryId
	r.Set("picture_category_id", _pictureCategoryId)
	return nil
}

// GetPictureCategoryId PictureCategoryId Getter
func (r AlibabaWdkPictureUploadAPIRequest) GetPictureCategoryId() int64 {
	return r._pictureCategoryId
}

// SetImg is Img Setter
// 图片二进制文件流,不能为空,允许png、jpg、gif图片格式,3M以内
func (r *AlibabaWdkPictureUploadAPIRequest) SetImg(_img *model.File) error {
	r._img = _img
	r.Set("img", _img)
	return nil
}

// GetImg Img Getter
func (r AlibabaWdkPictureUploadAPIRequest) GetImg() *model.File {
	return r._img
}
