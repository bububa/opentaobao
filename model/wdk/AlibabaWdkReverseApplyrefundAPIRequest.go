package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkReverseApplyrefundAPIRequest
逆向申请接口 API请求
alibaba.wdk.reverse.applyrefund

逆向渲染 */
type AlibabaWdkReverseApplyrefundAPIRequest struct {
	model.Params
	// 入参
	_paramApplyReverseReq *ApplyReverseReq
}

// New
