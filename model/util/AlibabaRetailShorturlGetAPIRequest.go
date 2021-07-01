package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaRetailShorturlGetAPIRequest
短链接获取 API请求
alibaba.retail.shorturl.get

短链接获取 */
type AlibabaRetailShorturlGetAPIRequest struct {
	model.Params
	// 源url
	_sourceUrl string
	// 系统自动生成
	_options *ShortUrlOption
}

// New
