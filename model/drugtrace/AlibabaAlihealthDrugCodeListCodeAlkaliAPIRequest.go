package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugCodeListCodeAlkaliAPIRequest
根据码查询码信息-黄麻碱 API请求
alibaba.alihealth.drug.code.list.code.alkali

服务描述
黄麻碱鉴证核查是基于在两定机构的药品管理（入库、出库或盘点）环节，增加扫码匹配
与验证鉴核流程；
此接口，针对有码药品，提供可通过追溯码获取该药品的基础信息和生产状况信息；
若所传的监管码是非最小包装监管码，且存在药品混包的情况，则此接口不支持。这种
情况下，需要分多次调用该接口。
核查平台优先过滤非8开头的，长度非20位数字的码信息。 */
type AlibabaAlihealthDrugCodeListCodeAlkaliAPIRequest struct {
	model.Params
	// 追溯码
	_codeList []string
	// 证件编号
	_certIsvNo string
	// 调用方式：formal-正式、test-测试
	_invocation string
	// 终端类型 1：零售
	_terminalType string
	// 调用零售药店名称
	_terminalName string
	// 城市名称
	_bureauName string
	// 错误信息
	_errorMessage string
}

// New
