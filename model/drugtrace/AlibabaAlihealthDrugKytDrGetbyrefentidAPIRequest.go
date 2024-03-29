package drugtrace

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugKytDrGetbyrefentidAPIRequest) Reset() {
	r._refEntId = ""
	r._destRefEntId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytDrGetbyrefentidAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.dr.getbyrefentid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytDrGetbyrefentidAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugKytDrGetbyrefentidAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 接口调用企业的唯一标识（接口调用者）
func (r *AlibabaAlihealthDrugKytDrGetbyrefentidAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugKytDrGetbyrefentidAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetDestRefEntId is DestRefEntId Setter
// 准备要查询的企业唯一标识（返回该唯一标识企业的详细信息）
func (r *AlibabaAlihealthDrugKytDrGetbyrefentidAPIRequest) SetDestRefEntId(_destRefEntId string) error {
	r._destRefEntId = _destRefEntId
	r.Set("dest_ref_ent_id", _destRefEntId)
	return nil
}

// GetDestRefEntId DestRefEntId Getter
func (r AlibabaAlihealthDrugKytDrGetbyrefentidAPIRequest) GetDestRefEntId() string {
	return r._destRefEntId
}

var poolAlibabaAlihealthDrugKytDrGetbyrefentidAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugKytDrGetbyrefentidRequest()
	},
}

// GetAlibabaAlihealthDrugKytDrGetbyrefentidRequest 从 sync.Pool 获取 AlibabaAlihealthDrugKytDrGetbyrefentidAPIRequest
func GetAlibabaAlihealthDrugKytDrGetbyrefentidAPIRequest() *AlibabaAlihealthDrugKytDrGetbyrefentidAPIRequest {
	return poolAlibabaAlihealthDrugKytDrGetbyrefentidAPIRequest.Get().(*AlibabaAlihealthDrugKytDrGetbyrefentidAPIRequest)
}

// ReleaseAlibabaAlihealthDrugKytDrGetbyrefentidAPIRequest 将 AlibabaAlihealthDrugKytDrGetbyrefentidAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugKytDrGetbyrefentidAPIRequest(v *AlibabaAlihealthDrugKytDrGetbyrefentidAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugKytDrGetbyrefentidAPIRequest.Put(v)
}
