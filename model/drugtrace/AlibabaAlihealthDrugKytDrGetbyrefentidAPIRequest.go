package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytDrGetbyrefentidAPIRequest 多融通过一个企业唯一标识查询企业详细信息 API请求
// alibaba.alihealth.drug.kyt.dr.getbyrefentid
//
// 根据企业唯一标识查看企业详细信息
type AlibabaAlihealthDrugKytDrGetbyrefentidAPIRequest struct {
	model.Params
	// 接口调用企业的唯一标识（接口调用者）
	_refEntId string
	// 准备要查询的企业唯一标识（返回该唯一标识企业的详细信息）
	_destRefEntId string
}

// NewAlibabaAlihealthDrugKytDrGetbyrefentidRequest 初始化AlibabaAlihealthDrugKytDrGetbyrefentidAPIRequest对象
func NewAlibabaAlihealthDrugKytDrGetbyrefentidRequest() *AlibabaAlihealthDrugKytDrGetbyrefentidAPIRequest {
	return &AlibabaAlihealthDrugKytDrGetbyrefentidAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytDrGetbyrefentidAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.dr.getbyrefentid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytDrGetbyrefentidAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is RefEntId Setter
// 接口调用企业的唯一标识（接口调用者）
func (r *AlibabaAlihealthDrugKytDrGetbyrefentidAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// Get RefEntId Getter
func (r AlibabaAlihealthDrugKytDrGetbyrefentidAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// Set is DestRefEntId Setter
// 准备要查询的企业唯一标识（返回该唯一标识企业的详细信息）
func (r *AlibabaAlihealthDrugKytDrGetbyrefentidAPIRequest) SetDestRefEntId(_destRefEntId string) error {
	r._destRefEntId = _destRefEntId
	r.Set("dest_ref_ent_id", _destRefEntId)
	return nil
}

// Get DestRefEntId Getter
func (r AlibabaAlihealthDrugKytDrGetbyrefentidAPIRequest) GetDestRefEntId() string {
	return r._destRefEntId
}
