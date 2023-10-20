package drugtrace

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugKytQuerydruginfoAPIRequest) Reset() {
	r._codeList = r._codeList[:0]
	r._wuliuRefEntId = ""
	r._huozhuRefEntId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytQuerydruginfoAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.querydruginfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytQuerydruginfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugKytQuerydruginfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCodeList is CodeList Setter
// 码列表
func (r *AlibabaAlihealthDrugKytQuerydruginfoAPIRequest) SetCodeList(_codeList []string) error {
	r._codeList = _codeList
	r.Set("code_list", _codeList)
	return nil
}

// GetCodeList CodeList Getter
func (r AlibabaAlihealthDrugKytQuerydruginfoAPIRequest) GetCodeList() []string {
	return r._codeList
}

// SetWuliuRefEntId is WuliuRefEntId Setter
// 物流企业refentid
func (r *AlibabaAlihealthDrugKytQuerydruginfoAPIRequest) SetWuliuRefEntId(_wuliuRefEntId string) error {
	r._wuliuRefEntId = _wuliuRefEntId
	r.Set("wuliu_ref_ent_id", _wuliuRefEntId)
	return nil
}

// GetWuliuRefEntId WuliuRefEntId Getter
func (r AlibabaAlihealthDrugKytQuerydruginfoAPIRequest) GetWuliuRefEntId() string {
	return r._wuliuRefEntId
}

// SetHuozhuRefEntId is HuozhuRefEntId Setter
// 生产企业refentid
func (r *AlibabaAlihealthDrugKytQuerydruginfoAPIRequest) SetHuozhuRefEntId(_huozhuRefEntId string) error {
	r._huozhuRefEntId = _huozhuRefEntId
	r.Set("huozhu_ref_ent_id", _huozhuRefEntId)
	return nil
}

// GetHuozhuRefEntId HuozhuRefEntId Getter
func (r AlibabaAlihealthDrugKytQuerydruginfoAPIRequest) GetHuozhuRefEntId() string {
	return r._huozhuRefEntId
}

var poolAlibabaAlihealthDrugKytQuerydruginfoAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugKytQuerydruginfoRequest()
	},
}

// GetAlibabaAlihealthDrugKytQuerydruginfoRequest 从 sync.Pool 获取 AlibabaAlihealthDrugKytQuerydruginfoAPIRequest
func GetAlibabaAlihealthDrugKytQuerydruginfoAPIRequest() *AlibabaAlihealthDrugKytQuerydruginfoAPIRequest {
	return poolAlibabaAlihealthDrugKytQuerydruginfoAPIRequest.Get().(*AlibabaAlihealthDrugKytQuerydruginfoAPIRequest)
}

// ReleaseAlibabaAlihealthDrugKytQuerydruginfoAPIRequest 将 AlibabaAlihealthDrugKytQuerydruginfoAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugKytQuerydruginfoAPIRequest(v *AlibabaAlihealthDrugKytQuerydruginfoAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugKytQuerydruginfoAPIRequest.Put(v)
}
