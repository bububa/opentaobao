package wangwang

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWangwangAbstractDeletewordAPIRequest
删除关键词 API请求
taobao.wangwang.abstract.deleteword

删除关键词，只支持json返回 */
type TaobaoWangwangAbstractDeletewordAPIRequest struct {
	model.Params
	// 关键词，长度大于0
	_word string
	// 传入参数的字符集
	_charset string
}

// New
