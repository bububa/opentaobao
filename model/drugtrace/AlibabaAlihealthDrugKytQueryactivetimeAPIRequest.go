package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytQueryactivetimeAPIRequest 药品激活状态同步 API请求
// alibaba.alihealth.drug.kyt.queryactivetime
//
// 根据赋码资源（CodeVersion + resCode）获得最新激活时间
// 应用于各地市对接前进行药品目录匹配，医保中心存在的药品可能比较陈旧杂乱
type AlibabaAlihealthDrugKytQueryactivetimeAPIRequest struct {
	model.Params
	// 码段的数组
	_resProdCodeList []string
	// 社保局(所属地市名称)
	_bureauName string
	// 请求终端名称
	_terminalName string
	// 终端类型：1005100-零售，1005200-医疗
	_terminalType string
	// 调用方式：formal-正式、test-测试
	_invocation string
}

// NewAlibabaAlihealthDrugKytQueryactivetimeRequest 初始化AlibabaAlihealthDrugKytQueryactivetimeAPIRequest对象
func NewAlibabaAlihealthDrugKytQueryactivetimeRequest() *AlibabaAlihealthDrugKytQueryactivetimeAPIRequest {
	return &AlibabaAlihealthDrugKytQueryactivetimeAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugKytQueryactivetimeAPIRequest) Reset() {
	r._resProdCodeList = r._resProdCodeList[:0]
	r._bureauName = ""
	r._terminalName = ""
	r._terminalType = ""
	r._invocation = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytQueryactivetimeAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.queryactivetime"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytQueryactivetimeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugKytQueryactivetimeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetResProdCodeList is ResProdCodeList Setter
// 码段的数组
func (r *AlibabaAlihealthDrugKytQueryactivetimeAPIRequest) SetResProdCodeList(_resProdCodeList []string) error {
	r._resProdCodeList = _resProdCodeList
	r.Set("res_prod_code_list", _resProdCodeList)
	return nil
}

// GetResProdCodeList ResProdCodeList Getter
func (r AlibabaAlihealthDrugKytQueryactivetimeAPIRequest) GetResProdCodeList() []string {
	return r._resProdCodeList
}

// SetBureauName is BureauName Setter
// 社保局(所属地市名称)
func (r *AlibabaAlihealthDrugKytQueryactivetimeAPIRequest) SetBureauName(_bureauName string) error {
	r._bureauName = _bureauName
	r.Set("bureau_name", _bureauName)
	return nil
}

// GetBureauName BureauName Getter
func (r AlibabaAlihealthDrugKytQueryactivetimeAPIRequest) GetBureauName() string {
	return r._bureauName
}

// SetTerminalName is TerminalName Setter
// 请求终端名称
func (r *AlibabaAlihealthDrugKytQueryactivetimeAPIRequest) SetTerminalName(_terminalName string) error {
	r._terminalName = _terminalName
	r.Set("terminal_name", _terminalName)
	return nil
}

// GetTerminalName TerminalName Getter
func (r AlibabaAlihealthDrugKytQueryactivetimeAPIRequest) GetTerminalName() string {
	return r._terminalName
}

// SetTerminalType is TerminalType Setter
// 终端类型：1005100-零售，1005200-医疗
func (r *AlibabaAlihealthDrugKytQueryactivetimeAPIRequest) SetTerminalType(_terminalType string) error {
	r._terminalType = _terminalType
	r.Set("terminal_type", _terminalType)
	return nil
}

// GetTerminalType TerminalType Getter
func (r AlibabaAlihealthDrugKytQueryactivetimeAPIRequest) GetTerminalType() string {
	return r._terminalType
}

// SetInvocation is Invocation Setter
// 调用方式：formal-正式、test-测试
func (r *AlibabaAlihealthDrugKytQueryactivetimeAPIRequest) SetInvocation(_invocation string) error {
	r._invocation = _invocation
	r.Set("invocation", _invocation)
	return nil
}

// GetInvocation Invocation Getter
func (r AlibabaAlihealthDrugKytQueryactivetimeAPIRequest) GetInvocation() string {
	return r._invocation
}

var poolAlibabaAlihealthDrugKytQueryactivetimeAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugKytQueryactivetimeRequest()
	},
}

// GetAlibabaAlihealthDrugKytQueryactivetimeRequest 从 sync.Pool 获取 AlibabaAlihealthDrugKytQueryactivetimeAPIRequest
func GetAlibabaAlihealthDrugKytQueryactivetimeAPIRequest() *AlibabaAlihealthDrugKytQueryactivetimeAPIRequest {
	return poolAlibabaAlihealthDrugKytQueryactivetimeAPIRequest.Get().(*AlibabaAlihealthDrugKytQueryactivetimeAPIRequest)
}

// ReleaseAlibabaAlihealthDrugKytQueryactivetimeAPIRequest 将 AlibabaAlihealthDrugKytQueryactivetimeAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugKytQueryactivetimeAPIRequest(v *AlibabaAlihealthDrugKytQueryactivetimeAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugKytQueryactivetimeAPIRequest.Put(v)
}
