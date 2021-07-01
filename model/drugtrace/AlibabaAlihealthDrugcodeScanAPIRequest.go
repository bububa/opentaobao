package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugcodeScanAPIRequest
查询扫码信息 API请求
alibaba.alihealth.drugcode.scan

查询扫码信息 */
type AlibabaAlihealthDrugcodeScanAPIRequest struct {
	model.Params
	// 20位码
	_code string
	// 渠道
	_queryAppName string
	// 用户ip
	_clientId string
	// 设备标识
	_deviceUtdid string
	// 用户ID
	_userId string
}

// New
