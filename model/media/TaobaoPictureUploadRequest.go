package media

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
上传单张图片 API请求
taobao.picture.upload

图片空间上传接口
*/
type TaobaoPictureUploadRequest struct {
    model.Params
    // 图片分类ID，设置具体某个分类ID或设置0上传到默认分类，只能传入一个分类
    _pictureCategoryId   int64
    // 图片二进制文件流,不能为空,允许png、jpg、gif图片格式,3M以内。
    _img   *model.File
    // 包括后缀名的图片标题,不能为空，如Bule.jpg,有些卖家希望图片上传后取图片文件的默认名
    _imageInputTitle   string
    // 图片标题,如果为空,传的图片标题就取去掉后缀名的image_input_title,超过50字符长度会截取50字符,重名会在标题末尾加"(1)";标题末尾已经有"(数字)"了，则数字加1
    _title   string
    // 图片上传的来源，有电脑版本宝贝发布，手机版本宝贝发布client:computer电脑版本宝贝使用，client:phone手机版本宝贝使用。注意：当client:phone时，图片限制为宽度在480-620之间，长度不能超过960，否则会报错。
    _clientType   string
    // 是否获取https连接
    _isHttps   bool
}

// 初始化TaobaoPictureUploadRequest对象
func NewTaobaoPictureUploadRequest() *TaobaoPictureUploadRequest{
    return &TaobaoPictureUploadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoPictureUploadRequest) GetApiMethodName() string {
    return "taobao.picture.upload"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoPictureUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PictureCategoryId Setter
// 图片分类ID，设置具体某个分类ID或设置0上传到默认分类，只能传入一个分类
func (r *TaobaoPictureUploadRequest) SetPictureCategoryId(_pictureCategoryId int64) error {
    r._pictureCategoryId = _pictureCategoryId
    r.Set("picture_category_id", _pictureCategoryId)
    return nil
}

// PictureCategoryId Getter
func (r TaobaoPictureUploadRequest) GetPictureCategoryId() int64 {
    return r._pictureCategoryId
}
// Img Setter
// 图片二进制文件流,不能为空,允许png、jpg、gif图片格式,3M以内。
func (r *TaobaoPictureUploadRequest) SetImg(_img *model.File) error {
    r._img = _img
    r.Set("img", _img)
    return nil
}

// Img Getter
func (r TaobaoPictureUploadRequest) GetImg() *model.File {
    return r._img
}
// ImageInputTitle Setter
// 包括后缀名的图片标题,不能为空，如Bule.jpg,有些卖家希望图片上传后取图片文件的默认名
func (r *TaobaoPictureUploadRequest) SetImageInputTitle(_imageInputTitle string) error {
    r._imageInputTitle = _imageInputTitle
    r.Set("image_input_title", _imageInputTitle)
    return nil
}

// ImageInputTitle Getter
func (r TaobaoPictureUploadRequest) GetImageInputTitle() string {
    return r._imageInputTitle
}
// Title Setter
// 图片标题,如果为空,传的图片标题就取去掉后缀名的image_input_title,超过50字符长度会截取50字符,重名会在标题末尾加"(1)";标题末尾已经有"(数字)"了，则数字加1
func (r *TaobaoPictureUploadRequest) SetTitle(_title string) error {
    r._title = _title
    r.Set("title", _title)
    return nil
}

// Title Getter
func (r TaobaoPictureUploadRequest) GetTitle() string {
    return r._title
}
// ClientType Setter
// 图片上传的来源，有电脑版本宝贝发布，手机版本宝贝发布client:computer电脑版本宝贝使用，client:phone手机版本宝贝使用。注意：当client:phone时，图片限制为宽度在480-620之间，长度不能超过960，否则会报错。
func (r *TaobaoPictureUploadRequest) SetClientType(_clientType string) error {
    r._clientType = _clientType
    r.Set("client_type", _clientType)
    return nil
}

// ClientType Getter
func (r TaobaoPictureUploadRequest) GetClientType() string {
    return r._clientType
}
// IsHttps Setter
// 是否获取https连接
func (r *TaobaoPictureUploadRequest) SetIsHttps(_isHttps bool) error {
    r._isHttps = _isHttps
    r.Set("is_https", _isHttps)
    return nil
}

// IsHttps Getter
func (r TaobaoPictureUploadRequest) GetIsHttps() bool {
    return r._isHttps
}
