package media

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoPictureUploadAPIRequest
上传单张图片 API请求
taobao.picture.upload

图片空间上传接口 */
type TaobaoPictureUploadAPIRequest struct {
	model.Params
	// 图片分类ID，设置具体某个分类ID或设置0上传到默认分类，只能传入一个分类
	_pictureCategoryId int64
	// 图片二进制文件流,不能为空,允许png、jpg、gif图片格式,3M以内。
	_img *model.File
	// 包括后缀名的图片标题,不能为空，如Bule.jpg,有些卖家希望图片上传后取图片文件的默认名
	_imageInputTitle string
	// 图片标题,如果为空,传的图片标题就取去掉后缀名的image_input_title,超过50字符长度会截取50字符,重名会在标题末尾加"(1)";标题末尾已经有"(数字)"了，则数字加1
	_title string
	// 图片上传的来源，有电脑版本宝贝发布，手机版本宝贝发布client:computer电脑版本宝贝使用，client:phone手机版本宝贝使用。注意：当client:phone时，图片限制为宽度在480-620之间，长度不能超过960，否则会报错。
	_clientType string
	// 是否获取https连接
	_isHttps bool
}

// New
