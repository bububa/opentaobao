package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugcodeApplycertAPIRequest
申请证书为对接方 API请求
alibaba.alihealth.drugcode.applycert

申请证书 为对接方(当前是药厂和中心化系统) */
type AlibabaAlihealthDrugcodeApplycertAPIRequest struct {
	model.Params
	// 设备唯一标识编号
	_serialNum string
	// 企业Id
	_refEntId string
	// 企业名称
	_entName string
	// 证书签名请求
	_csr string
	// 证书丢失时的操作类型 (true：证书丢失)
	_certLostFlag bool
}

// New
