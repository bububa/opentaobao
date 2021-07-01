package openim

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpenimSnfilterwordSetfilterAPIRequest
关键词过滤 API请求
taobao.openim.snfilterword.setfilter

设置openim关键词过滤 */
type TaobaoOpenimSnfilterwordSetfilterAPIRequest struct {
	model.Params
	// 上传者身份信息，区分不同上传者;只是记录，没有身份校验功能
	_creator string
	// 需要过滤的关键词
	_filterword string
	// 过滤原因描述
	_desc string
}

// New
