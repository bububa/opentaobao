package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytDrQueryupbillcodeAPIRequest 查询上游企业出库单据号 API请求
// alibaba.alihealth.drug.kyt.dr.queryupbillcode
//
// 疫苗温度合规补充需求-增加一个查询上游出库单号的接口。疾控在扫码入库时，接口通过扫到的码判定这个码对应所属的出库单据号
type AlibabaAlihealthDrugKytDrQueryupbillcodeAPIRequest struct {
	model.Params
	// 追溯码
	_code string
	// 企业ID （一般为要查询单据的收货企业）
	_refEntId string
}

// NewAlibabaAlihealthDrugKytDrQueryupbillcodeRequest 初始化AlibabaAlihealthDrugKytDrQueryupbillcodeAPIRequest对象
func NewAlibabaAlihealthDrugKytDrQueryupbillcodeRequest() *AlibabaAlihealthDrugKytDrQueryupbillcodeAPIRequest {
	return &AlibabaAlihealthDrugKytDrQueryupbillcodeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytDrQueryupbillcodeAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.dr.queryupbillcode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytDrQueryupbillcodeAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Code Setter
// 追溯码
func (r *AlibabaAlihealthDrugKytDrQueryupbillcodeAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// Get Code Getter
func (r AlibabaAlihealthDrugKytDrQueryupbillcodeAPIRequest) GetCode() string {
	return r._code
}

// Set is RefEntId Setter
// 企业ID （一般为要查询单据的收货企业）
func (r *AlibabaAlihealthDrugKytDrQueryupbillcodeAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// Get RefEntId Getter
func (r AlibabaAlihealthDrugKytDrQueryupbillcodeAPIRequest) GetRefEntId() string {
	return r._refEntId
}
