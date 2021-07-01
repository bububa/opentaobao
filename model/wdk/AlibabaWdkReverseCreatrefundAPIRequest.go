package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkReverseCreatrefundAPIRequest
逆向提交 API请求
alibaba.wdk.reverse.creatrefund

逆向申请提交 */
type AlibabaWdkReverseCreatrefundAPIRequest struct {
	model.Params
	// CreateReverseReq
	_paramCreateReverseReq *CreateReverseReq
}

// New
