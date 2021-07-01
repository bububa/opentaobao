package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugCodeCodeCheckMedicalInsuranceAPIRequest
码核查状态同步-医保 API请求
alibaba.alihealth.drug.code.code.check.medical.insurance

服务描述
核查平台在进行医保单据鉴证核查时，会记录单据中所有提交的追溯码信息；单据中的
追溯码包含所有正常和异常的数据；
此接口，针对正式鉴核的单据中提交的有效的、正常状态的追溯码，提供可由核查平台
发起，按单据鉴核时间顺序组织，向码上放心平台同步每笔单据核销的码状态信息；
入参采用数组方式提供，一次同步最多支持100条记录 */
type AlibabaAlihealthDrugCodeCodeCheckMedicalInsuranceAPIRequest struct {
	model.Params
	// 行政区域
	_bureauName string
	// 终端id
	_terminalId string
	// 终端类型（1005100-零售药店；1005200-医疗机构）
	_terminalType string
	// 核销类型(1012100：核销；1012900：退库)
	_cType string
	// 码列表
	_codes []string
	// 平台返回的终端id
	_terminalEntId string
}

// New
