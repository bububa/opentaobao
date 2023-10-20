package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytDrugdetailAPIRequest 查询药品详细信息 API请求
// alibaba.alihealth.drug.kyt.drugdetail
//
// 查询药品详细信息
type AlibabaAlihealthDrugKytDrugdetailAPIRequest struct {
	model.Params
	// 企业ID
	_refEntId string
	// 药品ID
	_drugEntBaseInfoId string
}

// NewAlibabaAlihealthDrugKytDrugdetailRequest 初始化AlibabaAlihealthDrugKytDrugdetailAPIRequest对象
func NewAlibabaAlihealthDrugKytDrugdetailRequest() *AlibabaAlihealthDrugKytDrugdetailAPIRequest {
	return &AlibabaAlihealthDrugKytDrugdetailAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugKytDrugdetailAPIRequest) Reset() {
	r._refEntId = ""
	r._drugEntBaseInfoId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytDrugdetailAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.drugdetail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytDrugdetailAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugKytDrugdetailAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 企业ID
func (r *AlibabaAlihealthDrugKytDrugdetailAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugKytDrugdetailAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetDrugEntBaseInfoId is DrugEntBaseInfoId Setter
// 药品ID
func (r *AlibabaAlihealthDrugKytDrugdetailAPIRequest) SetDrugEntBaseInfoId(_drugEntBaseInfoId string) error {
	r._drugEntBaseInfoId = _drugEntBaseInfoId
	r.Set("drug_ent_base_info_id", _drugEntBaseInfoId)
	return nil
}

// GetDrugEntBaseInfoId DrugEntBaseInfoId Getter
func (r AlibabaAlihealthDrugKytDrugdetailAPIRequest) GetDrugEntBaseInfoId() string {
	return r._drugEntBaseInfoId
}

var poolAlibabaAlihealthDrugKytDrugdetailAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugKytDrugdetailRequest()
	},
}

// GetAlibabaAlihealthDrugKytDrugdetailRequest 从 sync.Pool 获取 AlibabaAlihealthDrugKytDrugdetailAPIRequest
func GetAlibabaAlihealthDrugKytDrugdetailAPIRequest() *AlibabaAlihealthDrugKytDrugdetailAPIRequest {
	return poolAlibabaAlihealthDrugKytDrugdetailAPIRequest.Get().(*AlibabaAlihealthDrugKytDrugdetailAPIRequest)
}

// ReleaseAlibabaAlihealthDrugKytDrugdetailAPIRequest 将 AlibabaAlihealthDrugKytDrugdetailAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugKytDrugdetailAPIRequest(v *AlibabaAlihealthDrugKytDrugdetailAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugKytDrugdetailAPIRequest.Put(v)
}
