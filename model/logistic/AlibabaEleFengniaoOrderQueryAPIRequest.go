package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEleFengniaoOrderQueryAPIRequest
查询订单基本信息 API请求
alibaba.ele.fengniao.order.query

查询订单基本信息 */
type AlibabaEleFengniaoOrderQueryAPIRequest struct {
	model.Params
	// 参数
	_param *Param
}

// New
