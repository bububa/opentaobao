package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthTraceCodeSearchGetDrugresourcetopAPIRequest
根据码获取码信息 API请求
alibaba.alihealth.trace.code.search.get.drugresourcetop

根据码获取码信息 */
type AlibabaAlihealthTraceCodeSearchGetDrugresourcetopAPIRequest struct {
	model.Params
	// 追溯码
	_code string
	// 校验值
	_token string
	// 查询app名称
	_queryAppName string
	// 客户端ip
	_clientId string
	// 用户id
	_tbUserId int64
	// 设备号
	_deviceUtdid string
}

// New
