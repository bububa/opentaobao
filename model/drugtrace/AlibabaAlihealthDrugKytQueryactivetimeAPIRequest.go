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

// NewAlibabaAlihealthDrugKytQueryactivetimeRequest 初始化AlibabaAlihealthDrugKytQueryactivetimeAPIRequest对象
func NewAlibabaAlihealthDrugKytQueryactivetimeRequest() *AlibabaAlihealthDrugKytQueryactivetimeAPIRequest {
	return &AlibabaAlihealthDrugKytQueryactivetimeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytQueryactivetimeAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.queryactivetime"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytQueryactivetimeAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is BureauName Setter
// 社保局(所属地市名称)
func (r *AlibabaAlihealthDrugKytQueryactivetimeAPIRequest) SetBureauName(_bureauName string) error {
	r._bureauName = _bureauName
	r.Set("bureau_name", _bureauName)
	return nil
}

// Get BureauName Getter
func (r AlibabaAlihealthDrugKytQueryactivetimeAPIRequest) GetBureauName() string {
	return r._bureauName
}

// Set is TerminalName Setter
// 请求终端名称
func (r *AlibabaAlihealthDrugKytQueryactivetimeAPIRequest) SetTerminalName(_terminalName string) error {
	r._terminalName = _terminalName
	r.Set("terminal_name", _terminalName)
	return nil
}

// Get TerminalName Getter
func (r AlibabaAlihealthDrugKytQueryactivetimeAPIRequest) GetTerminalName() string {
	return r._terminalName
}

// Set is TerminalType Setter
// 终端类型：1005100-零售，1005200-医疗
func (r *AlibabaAlihealthDrugKytQueryactivetimeAPIRequest) SetTerminalType(_terminalType string) error {
	r._terminalType = _terminalType
	r.Set("terminal_type", _terminalType)
	return nil
}

// Get TerminalType Getter
func (r AlibabaAlihealthDrugKytQueryactivetimeAPIRequest) GetTerminalType() string {
	return r._terminalType
}

// Set is Invocation Setter
// 调用方式：formal-正式、test-测试
func (r *AlibabaAlihealthDrugKytQueryactivetimeAPIRequest) SetInvocation(_invocation string) error {
	r._invocation = _invocation
	r.Set("invocation", _invocation)
	return nil
}

// Get Invocation Getter
func (r AlibabaAlihealthDrugKytQueryactivetimeAPIRequest) GetInvocation() string {
	return r._invocation
}

// Set is ResProdCodeList Setter
// 码段的数组
func (r *AlibabaAlihealthDrugKytQueryactivetimeAPIRequest) SetResProdCodeList(_resProdCodeList []string) error {
	r._resProdCodeList = _resProdCodeList
	r.Set("res_prod_code_list", _resProdCodeList)
	return nil
}

// Get ResProdCodeList Getter
func (r AlibabaAlihealthDrugKytQueryactivetimeAPIRequest) GetResProdCodeList() []string {
	return r._resProdCodeList
}
