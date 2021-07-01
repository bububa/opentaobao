package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugKytQueryactivetimeAPIRequest
药品激活状态同步 API请求
alibaba.alihealth.drug.kyt.queryactivetime

根据赋码资源（CodeVersion + resCode）获得最新激活时间
应用于各地市对接前进行药品目录匹配，医保中心存在的药品可能比较陈旧杂乱 */
type AlibabaAlihealthDrugKytQueryactivetimeAPIRequest struct {
	model.Params
	// 社保局(所属地市名称)
	_bureauName string
	// 请求终端名称
	_terminalName string
	// 终端类型：1005100-零售，1005200-医疗
	_terminalType string
	// 调用方式：formal-正式、test-测试
	_invocation string
	// 码段的数组
	_resProdCodeList []string
}

// New
