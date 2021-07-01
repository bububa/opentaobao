package wangwang

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWangwangAbstractGetwordlistAPIRequest
获取关键词列表 API请求
taobao.wangwang.abstract.getwordlist

获取关键词列表，只支持json返回 */
type TaobaoWangwangAbstractGetwordlistAPIRequest struct {
	model.Params
	// 传入参数的字符集
	_charset string
}

// New
