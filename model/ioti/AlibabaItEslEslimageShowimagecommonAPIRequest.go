package ioti

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaItEslEslimageShowimagecommonAPIRequest
对混合云提供的刷图接口 API请求
alibaba.it.esl.eslimage.showimagecommon

混合云使用，提供给isv和我们混合云环境部署的应用刷图 */
type AlibabaItEslEslimageShowimagecommonAPIRequest struct {
	model.Params
	// ma地址
	_mac string
	// 图片的base64编码,图片要和价签大小一致
	_content2 string
	// 图片2的base64编码,图片要和价签大小一致
	_content string
	// 是否压缩，默认不传，字符串：yes，no
	_isCompress string
	// 是否手动刷图，默认不传，字符串：true，false
	_isManual string
}

// New
