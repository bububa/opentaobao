package wdkitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkPictureUploadAPIRequest
图片上传接口 API请求
alibaba.wdk.picture.upload

上传图片 */
type AlibabaWdkPictureUploadAPIRequest struct {
	model.Params
	// 图片分类ID，设置具体某个分类ID或设置0上传到默认分类，只能传入一个分类
	_pictureCategoryId int64
	// 图片二进制文件流,不能为空,允许png、jpg、gif图片格式,3M以内
	_img *model.File
	// 包括后缀名的图片标题,不能为空，如Bule.jpg,有些卖家希望图片上传后取图片文件的默认名
	_imgInputTitle string
	// 图片标题,如果为空,传的图片标题就取去掉后缀名的image_input_title,超过50字符长度会截取50字符,重名会在标题末尾加"(1)";标题末尾已经有"(数字)"了，则数字加1
	_title string
}

// New
