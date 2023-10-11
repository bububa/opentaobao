package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugcheckcodeChecklastfourAPIRequest 校验追溯码的后4位是否正确 API请求
// alibaba.alihealth.drugcheckcode.checklastfour
//
// 校验追溯码的后4位是否正确
type AlibabaAlihealthDrugcheckcodeChecklastfourAPIRequest struct {
	model.Params
	// 追溯码
	_code string
}

// NewAlibabaAlihealthDrugcheckcodeChecklastfourRequest 初始化AlibabaAlihealthDrugcheckcodeChecklastfourAPIRequest对象
func NewAlibabaAlihealthDrugcheckcodeChecklastfourRequest() *AlibabaAlihealthDrugcheckcodeChecklastfourAPIRequest {
	return &AlibabaAlihealthDrugcheckcodeChecklastfourAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugcheckcodeChecklastfourAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drugcheckcode.checklastfour"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugcheckcodeChecklastfourAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugcheckcodeChecklastfourAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCode is Code Setter
// 追溯码
func (r *AlibabaAlihealthDrugcheckcodeChecklastfourAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r AlibabaAlihealthDrugcheckcodeChecklastfourAPIRequest) GetCode() string {
	return r._code
}
