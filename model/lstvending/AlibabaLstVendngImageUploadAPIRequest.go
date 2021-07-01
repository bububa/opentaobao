package lstvending

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLstVendngImageUploadAPIRequest
售货机商品图片上传 API请求
alibaba.lst.vendng.image.upload

零售通自动售货机商品图片上传接口，主要为ISV厂商提供图片同步的通道，从而建立统一的商品图片库。 */
type AlibabaLstVendngImageUploadAPIRequest struct {
	model.Params
	// 图片文件字节数组
	_imgBytes *model.File
}

// New
