package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询追溯码对应的药品信息（疫情） API请求
alibaba.alihealth.drug.code.kyt.yq.querycode

通过追溯码码得到 药品名称、包装规格、剂型、剂型规格”、有效期至等信息。
*/
type AlibabaAlihealthDrugCodeKytYqQuerycodeRequest struct {
    model.Params
    // 追溯码
    _codeList   []string
    // 调用零售药店名称
    _terminalName   string
    // 门店所属的行政区域ID
    _bureauId   string
}

// 初始化AlibabaAlihealthDrugCodeKytYqQuerycodeRequest对象
func NewAlibabaAlihealthDrugCodeKytYqQuerycodeRequest() *AlibabaAlihealthDrugCodeKytYqQuerycodeRequest{
    return &AlibabaAlihealthDrugCodeKytYqQuerycodeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugCodeKytYqQuerycodeRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.code.kyt.yq.querycode"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugCodeKytYqQuerycodeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CodeList Setter
// 追溯码
func (r *AlibabaAlihealthDrugCodeKytYqQuerycodeRequest) SetCodeList(_codeList []string) error {
    r._codeList = _codeList
    r.Set("code_list", _codeList)
    return nil
}

// CodeList Getter
func (r AlibabaAlihealthDrugCodeKytYqQuerycodeRequest) GetCodeList() []string {
    return r._codeList
}
// TerminalName Setter
// 调用零售药店名称
func (r *AlibabaAlihealthDrugCodeKytYqQuerycodeRequest) SetTerminalName(_terminalName string) error {
    r._terminalName = _terminalName
    r.Set("terminal_name", _terminalName)
    return nil
}

// TerminalName Getter
func (r AlibabaAlihealthDrugCodeKytYqQuerycodeRequest) GetTerminalName() string {
    return r._terminalName
}
// BureauId Setter
// 门店所属的行政区域ID
func (r *AlibabaAlihealthDrugCodeKytYqQuerycodeRequest) SetBureauId(_bureauId string) error {
    r._bureauId = _bureauId
    r.Set("bureau_id", _bureauId)
    return nil
}

// BureauId Getter
func (r AlibabaAlihealthDrugCodeKytYqQuerycodeRequest) GetBureauId() string {
    return r._bureauId
}
