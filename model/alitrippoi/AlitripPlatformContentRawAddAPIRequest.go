package alitrippoi

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripPlatformContentRawAddAPIRequest
穷游内容写入接口 API请求
alitrip.platform.content.raw.add

穷游内容写入飞猪接口 */
type AlitripPlatformContentRawAddAPIRequest struct {
	model.Params
	// 写入入参
	_fliggyContentRequest *FliggyContentRequest
}

// New
