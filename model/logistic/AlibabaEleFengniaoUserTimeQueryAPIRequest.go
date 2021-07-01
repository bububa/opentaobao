package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEleFengniaoUserTimeQueryAPIRequest
蜂鸟询用户T API请求
alibaba.ele.fengniao.user.time.query

蜂鸟询用户T */
type AlibabaEleFengniaoUserTimeQueryAPIRequest struct {
	model.Params
	// 询T入参
	_param *PredictDeliveryTimeParam
}

// New
