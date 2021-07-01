package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaChongzhiQueryflowAPIRequest
查询指定商家的可用的流量宝贝接口 API请求
alibaba.chongzhi.queryflow

查询指定商家的可用的流量宝贝 */
type AlibabaChongzhiQueryflowAPIRequest struct {
	model.Params
	// 号码
	_mobile int64
	// 来源
	_clientSource string
}

// New
