package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWirelessPictureCheckAPIRequest
无线开放图片内容安全检查 API请求
taobao.wireless.picture.check

无线开放内容检查，提供涉黄暴力政治图片检查。更详情介绍见 <a href="https://help.aliyun.com/document_detail/70292.html" target="blank">阿里云内容安全</a>
此API会进行两个场景审核，平均RT为1s。 */
type TaobaoWirelessPictureCheckAPIRequest struct {
	model.Params
	// 图片的URL,URL必须为淘系安全域名地址。图片格式支持png,jpg,webp
	_url string
}

// New
