package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytGetcodebaseinfoAPIRequest 码的药品信息查询 API请求
// alibaba.alihealth.drug.kyt.getcodebaseinfo
//
// 提供根据码查询码基本信息接口
type AlibabaAlihealthDrugKytGetcodebaseinfoAPIRequest struct {
	model.Params
	// 企业唯一标识
	_refEntId string
	// 码
	_code string
}

// NewAlibabaAlihealthDrugKytGetcodebaseinfoRequest 初始化AlibabaAlihealthDrugKytGetcodebaseinfoAPIRequest对象
func NewAlibabaAlihealthDrugKytGetcodebaseinfoRequest() *AlibabaAlihealthDrugKytGetcodebaseinfoAPIRequest {
	return &AlibabaAlihealthDrugKytGetcodebaseinfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytGetcodebaseinfoAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.getcodebaseinfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytGetcodebaseinfoAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetRefEntId is RefEntId Setter
// 企业唯一标识
func (r *AlibabaAlihealthDrugKytGetcodebaseinfoAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugKytGetcodebaseinfoAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetCode is Code Setter
// 码
func (r *AlibabaAlihealthDrugKytGetcodebaseinfoAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r AlibabaAlihealthDrugKytGetcodebaseinfoAPIRequest) GetCode() string {
	return r._code
}
