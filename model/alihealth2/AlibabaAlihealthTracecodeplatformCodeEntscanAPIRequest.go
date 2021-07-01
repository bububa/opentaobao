package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthTracecodeplatformCodeEntscanAPIRequest
药品商家扫码 API请求
alibaba.alihealth.tracecodeplatform.code.entscan

药品商家扫描药品监管码，只有该商家的药才返回 */
type AlibabaAlihealthTracecodeplatformCodeEntscanAPIRequest struct {
	model.Params
	// 药监码
	_code string
	// 不同企业有不同的标识
	_serviceFlag string
}

// New
