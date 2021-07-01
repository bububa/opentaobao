package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugKytRecordinfoAPIRequest
快易通健康检查 API请求
alibaba.alihealth.drug.kyt.recordinfo

快易通健康检查 */
type AlibabaAlihealthDrugKytRecordinfoAPIRequest struct {
	model.Params
	// 服务名
	_serviceName string
	// 类型
	_serviceType string
	// 输入参数
	_inputParam string
	// 其他参数
	_otherParam string
	// 级别
	_logLevel string
}

// New
