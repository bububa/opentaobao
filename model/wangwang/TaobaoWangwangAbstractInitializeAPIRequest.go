package wangwang

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWangwangAbstractInitializeAPIRequest
模糊查询服务初始化 API请求
taobao.wangwang.abstract.initialize

模糊查询服务初始化，只支持json返回 */
type TaobaoWangwangAbstractInitializeAPIRequest struct {
	model.Params
	// 传入参数的字符集
	_charset string
}

// New
