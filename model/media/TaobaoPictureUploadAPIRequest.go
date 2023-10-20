package media

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPictureUploadAPIRequest 上传单张图片 API请求
// taobao.picture.upload
//
// 图片空间上传接口
type TaobaoPictureUploadAPIRequest struct {
	model.Params
	// 包括后缀名的图片标题,不能为空，如Bule.jpg,有些卖家希望图片上传后取图片文件的默认名
	_imageInputTitle string
	// 图片标题,如果为空,传的图片标题就取去掉后缀名的image_input_title,超过50字符长度会截取50字符,重名会在标题末尾加"(1)";标题末尾已经有"(数字)"了，则数字加1
	_title string
	// 图片上传的来源，有电脑版本宝贝发布，手机版本宝贝发布client:computer电脑版本宝贝使用，client:phone手机版本宝贝使用。注意：当client:phone时，图片限制为宽度在480-620之间，长度不能超过960，否则会报错。
	_clientType string
	// 图片分类ID，设置具体某个分类ID或设置0上传到默认分类，只能传入一个分类
	_pictureCategoryId int64
	// 图片二进制文件流,不能为空,允许png、jpg、gif图片格式,3M以内。
	_img *model.File
	// 是否获取https连接
	_isHttps bool
}

// NewTaobaoPictureUploadRequest 初始化TaobaoPictureUploadAPIRequest对象
func NewTaobaoPictureUploadRequest() *TaobaoPictureUploadAPIRequest {
	return &TaobaoPictureUploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPictureUploadAPIRequest) GetApiMethodName() string {
	return "taobao.picture.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPictureUploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoPictureUploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetImageInputTitle is ImageInputTitle Setter
// 包括后缀名的图片标题,不能为空，如Bule.jpg,有些卖家希望图片上传后取图片文件的默认名
func (r *TaobaoPictureUploadAPIRequest) SetImageInputTitle(_imageInputTitle string) error {
	r._imageInputTitle = _imageInputTitle
	r.Set("image_input_title", _imageInputTitle)
	return nil
}

// GetImageInputTitle ImageInputTitle Getter
func (r TaobaoPictureUploadAPIRequest) GetImageInputTitle() string {
	return r._imageInputTitle
}

// SetTitle is Title Setter
// 图片标题,如果为空,传的图片标题就取去掉后缀名的image_input_title,超过50字符长度会截取50字符,重名会在标题末尾加&#34;(1)&#34;;标题末尾已经有&#34;(数字)&#34;了，则数字加1
func (r *TaobaoPictureUploadAPIRequest) SetTitle(_title string) error {
	r._title = _title
	r.Set("title", _title)
	return nil
}

// GetTitle Title Getter
func (r TaobaoPictureUploadAPIRequest) GetTitle() string {
	return r._title
}

// SetClientType is ClientType Setter
// 图片上传的来源，有电脑版本宝贝发布，手机版本宝贝发布client:computer电脑版本宝贝使用，client:phone手机版本宝贝使用。注意：当client:phone时，图片限制为宽度在480-620之间，长度不能超过960，否则会报错。
func (r *TaobaoPictureUploadAPIRequest) SetClientType(_clientType string) error {
	r._clientType = _clientType
	r.Set("client_type", _clientType)
	return nil
}

// GetClientType ClientType Getter
func (r TaobaoPictureUploadAPIRequest) GetClientType() string {
	return r._clientType
}

// SetPictureCategoryId is PictureCategoryId Setter
// 图片分类ID，设置具体某个分类ID或设置0上传到默认分类，只能传入一个分类
func (r *TaobaoPictureUploadAPIRequest) SetPictureCategoryId(_pictureCategoryId int64) error {
	r._pictureCategoryId = _pictureCategoryId
	r.Set("picture_category_id", _pictureCategoryId)
	return nil
}

// GetPictureCategoryId PictureCategoryId Getter
func (r TaobaoPictureUploadAPIRequest) GetPictureCategoryId() int64 {
	return r._pictureCategoryId
}

// SetImg is Img Setter
// 图片二进制文件流,不能为空,允许png、jpg、gif图片格式,3M以内。
func (r *TaobaoPictureUploadAPIRequest) SetImg(_img *model.File) error {
	r._img = _img
	r.Set("img", _img)
	return nil
}

// GetImg Img Getter
func (r TaobaoPictureUploadAPIRequest) GetImg() *model.File {
	return r._img
}

// SetIsHttps is IsHttps Setter
// 是否获取https连接
func (r *TaobaoPictureUploadAPIRequest) SetIsHttps(_isHttps bool) error {
	r._isHttps = _isHttps
	r.Set("is_https", _isHttps)
	return nil
}

// GetIsHttps IsHttps Getter
func (r TaobaoPictureUploadAPIRequest) GetIsHttps() bool {
	return r._isHttps
}
