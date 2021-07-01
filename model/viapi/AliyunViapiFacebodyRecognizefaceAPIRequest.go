package viapi

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliyunViapiFacebodyRecognizefaceAPIRequest
人脸属性识别 API请求
aliyun.viapi.facebody.recognizeface

输入图片之后，在人脸检测定位返回结果的基础上，识别各个检测人脸的四种属性，返回性别（男/女）、年龄、表情（笑/不笑）、眼镜（戴/不戴）；并可返回人脸的1024维深度学习特征。(参数图片/链接必须通过以下方式获取: https://help.aliyun.com/document_detail/155645.html ) */
type AliyunViapiFacebodyRecognizefaceAPIRequest struct {
	model.Params
	// 待检测图片链接
	_imageUrl string
	// 图片类型, ,取值范围[0:ImageURL, 1:ImageContent]
	_imageType int64
}

// New
