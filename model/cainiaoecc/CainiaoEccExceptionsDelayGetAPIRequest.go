package cainiaoecc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoEccExceptionsDelayGetAPIRequest
菜鸟控制塔包裹滞留异常信息获取 API请求
cainiao.ecc.exceptions.delay.get

菜鸟控制塔包裹滞留异常信息获取 */
type CainiaoEccExceptionsDelayGetAPIRequest struct {
	model.Params
	// 运单号
	_mailNo string
}

// New
