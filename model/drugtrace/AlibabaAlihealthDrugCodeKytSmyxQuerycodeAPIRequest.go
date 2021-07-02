package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugCodeKytSmyxQuerycodeAPIRequest 扫码营销码查询 API请求
// alibaba.alihealth.drug.code.kyt.smyx.querycode
//
// 此接口针对有码药品，提供可通过追溯码获取该药品的基础信息和生产信息；
// 核查平台优先过滤非8开头的，长度非20位数字的码信息。
type AlibabaAlihealthDrugCodeKytSmyxQuerycodeAPIRequest struct {
	model.Params
	// 企业唯一标识
	_refEntId string
	// 码列表
	_codes []string
}

// NewAlibabaAlihealthDrugCodeKytSmyxQuerycodeRequest 初始化AlibabaAlihealthDrugCodeKytSmyxQuerycodeAPIRequest对象
func NewAlibabaAlihealthDrugCodeKytSmyxQuerycodeRequest() *AlibabaAlihealthDrugCodeKytSmyxQuerycodeAPIRequest {
	return &AlibabaAlihealthDrugCodeKytSmyxQuerycodeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugCodeKytSmyxQuerycodeAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.code.kyt.smyx.querycode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugCodeKytSmyxQuerycodeAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetRefEntId is RefEntId Setter
// 企业唯一标识
func (r *AlibabaAlihealthDrugCodeKytSmyxQuerycodeAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugCodeKytSmyxQuerycodeAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetCodes is Codes Setter
// 码列表
func (r *AlibabaAlihealthDrugCodeKytSmyxQuerycodeAPIRequest) SetCodes(_codes []string) error {
	r._codes = _codes
	r.Set("codes", _codes)
	return nil
}

// GetCodes Codes Getter
func (r AlibabaAlihealthDrugCodeKytSmyxQuerycodeAPIRequest) GetCodes() []string {
	return r._codes
}
