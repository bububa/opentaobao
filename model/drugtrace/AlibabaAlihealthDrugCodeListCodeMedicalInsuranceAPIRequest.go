package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugCodeListCodeMedicalInsuranceAPIRequest
根据码获取码信息接口-医保 API请求
alibaba.alihealth.drug.code.list.code.medical.insurance

服务描述
医保鉴证核查是基于在两定机构的药品管理（入库、出库或盘点）环节，增加扫码匹配
与验证鉴核流程；
此接口，针对有码药品，提供可通过追溯码获取该药品的基础信息和生产状况信息；
若所传的监管码是非最小包装监管码，且存在药品混包的情况，则此接口不支持。这种
情况下，需要分多次调用该接口。
核查平台优先过滤非8开头的，长度非20位数字的码信息。 */
type AlibabaAlihealthDrugCodeListCodeMedicalInsuranceAPIRequest struct {
	model.Params
	// 追溯码
	_codeList []string
	// ISV开放平台帐号标识
	_certIsvNo string
	// 调用方式：formal-正式、test-测试
	_invocation string
	// 终端类型 1005100-零售药店 ；10052-医疗机构
	_terminalType string
	// 调用零售药店名称
	_terminalName string
	// 门店所属的行政区域ID
	_bureauId string
	// 零售终端id
	_terminalEntId string
}

// New
