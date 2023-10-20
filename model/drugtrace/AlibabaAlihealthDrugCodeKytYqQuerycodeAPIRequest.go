package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugcodekytyqquerycodeAPIRequest 查询追溯码对应的药品信息（疫情） API请求
// alibaba.alihealth.drug.code.kyt.yq.querycode
//
// 通过追溯码码得到 药品名称、包装规格、剂型、剂型规格”、有效期至等信息。
type AlibabaalihealthdrugcodekytyqquerycodeAPIRequest struct {
	model.Params
	// 追溯码
	_codeList []string
	// 调用零售药店名称
	_terminalName string
	// 门店所属的行政区域ID
	_bureauId string
}

// NewAlibabaalihealthdrugcodekytyqquerycodeRequest 初始化AlibabaalihealthdrugcodekytyqquerycodeAPIRequest对象
func NewAlibabaalihealthdrugcodekytyqquerycodeRequest() *AlibabaalihealthdrugcodekytyqquerycodeAPIRequest {
	return &AlibabaalihealthdrugcodekytyqquerycodeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugcodekytyqquerycodeAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.code.kyt.yq.querycode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugcodekytyqquerycodeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugcodekytyqquerycodeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCodeList is CodeList Setter
// 追溯码
func (r *AlibabaalihealthdrugcodekytyqquerycodeAPIRequest) SetCodeList(_codeList []string) error {
	r._codeList = _codeList
	r.Set("code_list", _codeList)
	return nil
}

// GetCodeList CodeList Getter
func (r AlibabaalihealthdrugcodekytyqquerycodeAPIRequest) GetCodeList() []string {
	return r._codeList
}

// SetTerminalName is TerminalName Setter
// 调用零售药店名称
func (r *AlibabaalihealthdrugcodekytyqquerycodeAPIRequest) SetTerminalName(_terminalName string) error {
	r._terminalName = _terminalName
	r.Set("terminal_name", _terminalName)
	return nil
}

// GetTerminalName TerminalName Getter
func (r AlibabaalihealthdrugcodekytyqquerycodeAPIRequest) GetTerminalName() string {
	return r._terminalName
}

// SetBureauId is BureauId Setter
// 门店所属的行政区域ID
func (r *AlibabaalihealthdrugcodekytyqquerycodeAPIRequest) SetBureauId(_bureauId string) error {
	r._bureauId = _bureauId
	r.Set("bureau_id", _bureauId)
	return nil
}

// GetBureauId BureauId Getter
func (r AlibabaalihealthdrugcodekytyqquerycodeAPIRequest) GetBureauId() string {
	return r._bureauId
}
