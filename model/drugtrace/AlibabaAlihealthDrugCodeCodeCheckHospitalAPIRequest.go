package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugCodeCodeCheckHospitalAPIRequest
码核查状态同步-医院 API请求
alibaba.alihealth.drug.code.code.check.hospital

码核查状态同步-医院 */
type AlibabaAlihealthDrugCodeCodeCheckHospitalAPIRequest struct {
	model.Params
	// 认证企业refEntId
	_authRefEntId string
	// 企业refEntId
	_refEntId string
	// 城市名
	_bureauName string
	// 终端名称
	_terminalName string
	// 终端类型
	_terminalType string
	// 核销类型
	_cType string
	// 码列表
	_codes []string
}

// New
