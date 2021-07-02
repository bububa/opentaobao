package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytQuerydruginfoAPIRequest 码查询药品 API请求
// alibaba.alihealth.drug.kyt.querydruginfo
//
// 通过追溯码查询药品信息
type AlibabaAlihealthDrugKytQuerydruginfoAPIRequest struct {
	model.Params
	// 码列表
	_codeList []string
	// 物流企业refentid
	_wuliuRefEntId string
	// 生产企业refentid
	_huozhuRefEntId string
}

// NewAlibabaAlihealthDrugKytQuerydruginfoRequest 初始化AlibabaAlihealthDrugKytQuerydruginfoAPIRequest对象
func NewAlibabaAlihealthDrugKytQuerydruginfoRequest() *AlibabaAlihealthDrugKytQuerydruginfoAPIRequest {
	return &AlibabaAlihealthDrugKytQuerydruginfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytQuerydruginfoAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.querydruginfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytQuerydruginfoAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is CodeList Setter
// 码列表
func (r *AlibabaAlihealthDrugKytQuerydruginfoAPIRequest) SetCodeList(_codeList []string) error {
	r._codeList = _codeList
	r.Set("code_list", _codeList)
	return nil
}

// Get CodeList Getter
func (r AlibabaAlihealthDrugKytQuerydruginfoAPIRequest) GetCodeList() []string {
	return r._codeList
}

// Set is WuliuRefEntId Setter
// 物流企业refentid
func (r *AlibabaAlihealthDrugKytQuerydruginfoAPIRequest) SetWuliuRefEntId(_wuliuRefEntId string) error {
	r._wuliuRefEntId = _wuliuRefEntId
	r.Set("wuliu_ref_ent_id", _wuliuRefEntId)
	return nil
}

// Get WuliuRefEntId Getter
func (r AlibabaAlihealthDrugKytQuerydruginfoAPIRequest) GetWuliuRefEntId() string {
	return r._wuliuRefEntId
}

// Set is HuozhuRefEntId Setter
// 生产企业refentid
func (r *AlibabaAlihealthDrugKytQuerydruginfoAPIRequest) SetHuozhuRefEntId(_huozhuRefEntId string) error {
	r._huozhuRefEntId = _huozhuRefEntId
	r.Set("huozhu_ref_ent_id", _huozhuRefEntId)
	return nil
}

// Get HuozhuRefEntId Getter
func (r AlibabaAlihealthDrugKytQuerydruginfoAPIRequest) GetHuozhuRefEntId() string {
	return r._huozhuRefEntId
}
