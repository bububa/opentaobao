package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugkytquerydruginfoAPIRequest 码查询药品 API请求
// alibaba.alihealth.drug.kyt.querydruginfo
//
// 通过追溯码查询药品信息
type AlibabaalihealthdrugkytquerydruginfoAPIRequest struct {
	model.Params
	// 码列表
	_codeList []string
	// 物流企业refentid
	_wuliuRefEntId string
	// 生产企业refentid
	_huozhuRefEntId string
}

// NewAlibabaalihealthdrugkytquerydruginfoRequest 初始化AlibabaalihealthdrugkytquerydruginfoAPIRequest对象
func NewAlibabaalihealthdrugkytquerydruginfoRequest() *AlibabaalihealthdrugkytquerydruginfoAPIRequest {
	return &AlibabaalihealthdrugkytquerydruginfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugkytquerydruginfoAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.querydruginfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugkytquerydruginfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugkytquerydruginfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCodeList is CodeList Setter
// 码列表
func (r *AlibabaalihealthdrugkytquerydruginfoAPIRequest) SetCodeList(_codeList []string) error {
	r._codeList = _codeList
	r.Set("code_list", _codeList)
	return nil
}

// GetCodeList CodeList Getter
func (r AlibabaalihealthdrugkytquerydruginfoAPIRequest) GetCodeList() []string {
	return r._codeList
}

// SetWuliuRefEntId is WuliuRefEntId Setter
// 物流企业refentid
func (r *AlibabaalihealthdrugkytquerydruginfoAPIRequest) SetWuliuRefEntId(_wuliuRefEntId string) error {
	r._wuliuRefEntId = _wuliuRefEntId
	r.Set("wuliu_ref_ent_id", _wuliuRefEntId)
	return nil
}

// GetWuliuRefEntId WuliuRefEntId Getter
func (r AlibabaalihealthdrugkytquerydruginfoAPIRequest) GetWuliuRefEntId() string {
	return r._wuliuRefEntId
}

// SetHuozhuRefEntId is HuozhuRefEntId Setter
// 生产企业refentid
func (r *AlibabaalihealthdrugkytquerydruginfoAPIRequest) SetHuozhuRefEntId(_huozhuRefEntId string) error {
	r._huozhuRefEntId = _huozhuRefEntId
	r.Set("huozhu_ref_ent_id", _huozhuRefEntId)
	return nil
}

// GetHuozhuRefEntId HuozhuRefEntId Getter
func (r AlibabaalihealthdrugkytquerydruginfoAPIRequest) GetHuozhuRefEntId() string {
	return r._huozhuRefEntId
}
