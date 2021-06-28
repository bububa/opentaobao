package media

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
上传单张图片 APIRequest
taobao.picture.upload

图片空间上传接口
*/
type TaobaoPictureUploadRequest struct {
    model.Params

    // 图片分类ID，设置具体某个分类ID或设置0上传到默认分类，只能传入一个分类
    pictureCategoryId   int64 

    // 图片二进制文件流,不能为空,允许png、jpg、gif图片格式,3M以内。
    img   []*model.File 

    // 包括后缀名的图片标题,不能为空，如Bule.jpg,有些卖家希望图片上传后取图片文件的默认名
    imageInputTitle   string 

    // 图片标题,如果为空,传的图片标题就取去掉后缀名的image_input_title,超过50字符长度会截取50字符,重名会在标题末尾加"(1)";标题末尾已经有"(数字)"了，则数字加1
    title   string 

    // 图片上传的来源，有电脑版本宝贝发布，手机版本宝贝发布client:computer电脑版本宝贝使用，client:phone手机版本宝贝使用。注意：当client:phone时，图片限制为宽度在480-620之间，长度不能超过960，否则会报错。
    clientType   string 

    // 是否获取https连接
    isHttps   bool 

}

func NewTaobaoPictureUploadRequest() *TaobaoPictureUploadRequest{
    return &TaobaoPictureUploadRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoPictureUploadRequest) GetApiMethodName() string {
    return "taobao.picture.upload"
}

func (r TaobaoPictureUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoPictureUploadRequest) SetPictureCategoryId(pictureCategoryId int64) error {
    r.pictureCategoryId = pictureCategoryId
    r.Set("picture_category_id", pictureCategoryId)
    return nil
}

func (r TaobaoPictureUploadRequest) GetPictureCategoryId() int64 {
    return r.pictureCategoryId
}

func (r *TaobaoPictureUploadRequest) SetImg(img []*model.File) error {
    r.img = img
    r.Set("img", img)
    return nil
}

func (r TaobaoPictureUploadRequest) GetImg() []*model.File {
    return r.img
}

func (r *TaobaoPictureUploadRequest) SetImageInputTitle(imageInputTitle string) error {
    r.imageInputTitle = imageInputTitle
    r.Set("image_input_title", imageInputTitle)
    return nil
}

func (r TaobaoPictureUploadRequest) GetImageInputTitle() string {
    return r.imageInputTitle
}

func (r *TaobaoPictureUploadRequest) SetTitle(title string) error {
    r.title = title
    r.Set("title", title)
    return nil
}

func (r TaobaoPictureUploadRequest) GetTitle() string {
    return r.title
}

func (r *TaobaoPictureUploadRequest) SetClientType(clientType string) error {
    r.clientType = clientType
    r.Set("client_type", clientType)
    return nil
}

func (r TaobaoPictureUploadRequest) GetClientType() string {
    return r.clientType
}

func (r *TaobaoPictureUploadRequest) SetIsHttps(isHttps bool) error {
    r.isHttps = isHttps
    r.Set("is_https", isHttps)
    return nil
}

func (r TaobaoPictureUploadRequest) GetIsHttps() bool {
    return r.isHttps
}

