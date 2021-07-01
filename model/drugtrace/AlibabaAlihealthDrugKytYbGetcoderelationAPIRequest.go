package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugKytYbGetcoderelationAPIRequest
医保-查询码的所有子码 API请求
alibaba.alihealth.drug.kyt.yb.getcoderelation

应用于药店或医院入库环节，通过扫码获取下级码进行入库；
通过码查询所有子码以及包装比例 */
type AlibabaAlihealthDrugKytYbGetcoderelationAPIRequest struct {
	model.Params
	// 社保局(所属地市名称)
	_bureauName string
	// 请求终端名称
	_terminalName string
	// 终端类型：1005100-零售，1005200-医疗
	_terminalType string
	// 调用方式：formal-正式、test-测试
	_invocation string
	// 追溯码
	_code string
}

// New
