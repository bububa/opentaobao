package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
药品激活状态同步 API请求
alibaba.alihealth.drug.kyt.queryactivetime

根据赋码资源（CodeVersion + resCode）获得最新激活时间
应用于各地市对接前进行药品目录匹配，医保中心存在的药品可能比较陈旧杂乱
*/
type AlibabaAlihealthDrugKytQueryactivetimeRequest struct {
    model.Params
    // 社保局(所属地市名称)
    bureauName   string
    // 请求终端名称
    terminalName   string
    // 终端类型：1005100-零售，1005200-医疗
    terminalType   string
    // 调用方式：formal-正式、test-测试
    invocation   string
    // 码段的数组
    resProdCodeList   []string
}

// 初始化AlibabaAlihealthDrugKytQueryactivetimeRequest对象
func NewAlibabaAlihealthDrugKytQueryactivetimeRequest() *AlibabaAlihealthDrugKytQueryactivetimeRequest{
    return &AlibabaAlihealthDrugKytQueryactivetimeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytQueryactivetimeRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.queryactivetime"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytQueryactivetimeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BureauName Setter
// 社保局(所属地市名称)
func (r *AlibabaAlihealthDrugKytQueryactivetimeRequest) SetBureauName(bureauName string) error {
    r.bureauName = bureauName
    r.Set("bureau_name", bureauName)
    return nil
}

// BureauName Getter
func (r AlibabaAlihealthDrugKytQueryactivetimeRequest) GetBureauName() string {
    return r.bureauName
}
// TerminalName Setter
// 请求终端名称
func (r *AlibabaAlihealthDrugKytQueryactivetimeRequest) SetTerminalName(terminalName string) error {
    r.terminalName = terminalName
    r.Set("terminal_name", terminalName)
    return nil
}

// TerminalName Getter
func (r AlibabaAlihealthDrugKytQueryactivetimeRequest) GetTerminalName() string {
    return r.terminalName
}
// TerminalType Setter
// 终端类型：1005100-零售，1005200-医疗
func (r *AlibabaAlihealthDrugKytQueryactivetimeRequest) SetTerminalType(terminalType string) error {
    r.terminalType = terminalType
    r.Set("terminal_type", terminalType)
    return nil
}

// TerminalType Getter
func (r AlibabaAlihealthDrugKytQueryactivetimeRequest) GetTerminalType() string {
    return r.terminalType
}
// Invocation Setter
// 调用方式：formal-正式、test-测试
func (r *AlibabaAlihealthDrugKytQueryactivetimeRequest) SetInvocation(invocation string) error {
    r.invocation = invocation
    r.Set("invocation", invocation)
    return nil
}

// Invocation Getter
func (r AlibabaAlihealthDrugKytQueryactivetimeRequest) GetInvocation() string {
    return r.invocation
}
// ResProdCodeList Setter
// 码段的数组
func (r *AlibabaAlihealthDrugKytQueryactivetimeRequest) SetResProdCodeList(resProdCodeList []string) error {
    r.resProdCodeList = resProdCodeList
    r.Set("res_prod_code_list", resProdCodeList)
    return nil
}

// ResProdCodeList Getter
func (r AlibabaAlihealthDrugKytQueryactivetimeRequest) GetResProdCodeList() []string {
    return r.resProdCodeList
}
