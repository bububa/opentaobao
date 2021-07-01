package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaChongzhiQueryecardsAPIRequest
查询指定商家的可用的话费宝贝接口 API请求
alibaba.chongzhi.queryecards

查询指定商家的可用的话费宝贝 */
type AlibabaChongzhiQueryecardsAPIRequest struct {
	model.Params
	// 号码
	_mobile int64
	// 来源
	_clientSource string
}

// New
