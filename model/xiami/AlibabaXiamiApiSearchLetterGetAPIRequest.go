package xiami

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaXiamiApiSearchLetterGetAPIRequest
搜索接口（首字母） API请求
alibaba.xiami.api.search.letter.get

搜索接口（首字母） */
type AlibabaXiamiApiSearchLetterGetAPIRequest struct {
	model.Params
	// 搜索关键字
	_key string
}

// New
