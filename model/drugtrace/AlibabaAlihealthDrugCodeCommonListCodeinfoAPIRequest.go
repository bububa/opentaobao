package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugCodeCommonListCodeinfoAPIRequest
通用查询码接口 API请求
alibaba.alihealth.drug.code.common.list.codeinfo

通用查询码接口 */
type AlibabaAlihealthDrugCodeCommonListCodeinfoAPIRequest struct {
	model.Params
	// 企业refEntId
	_refEntId string
	// 标示医院业务
	_searchSource string
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
	// 验证权限企业id
	_authRefEntId string
}

// New
