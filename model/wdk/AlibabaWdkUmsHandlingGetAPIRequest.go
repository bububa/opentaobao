package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkUmsHandlingGetAPIRequest
加工单-回流单（新接口） API请求
alibaba.wdk.ums.handling.get

加工单-回流单（新接口） */
type AlibabaWdkUmsHandlingGetAPIRequest struct {
	model.Params
	// 仓库编码
	_warehouseCode string
}

// New
