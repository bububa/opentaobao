package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytGetbyrefentidAPIRequest 根据企业唯一标识查看企业详细信息 API请求
// alibaba.alihealth.drug.kyt.getbyrefentid
//
// 根据企业唯一标识查看企业详细信息
type AlibabaAlihealthDrugKytGetbyrefentidAPIRequest struct {
	model.Params
	// 接口调用企业的唯一标识（接口调用者）
	_refEntId string
	// 准备要查询的企业唯一标识（返回该唯一标识企业的详细信息）
	_destRefEntId string
}

// NewAlibabaAlihealthDrugKytGetbyrefentidRequest 初始化AlibabaAlihealthDrugKytGetbyrefentidAPIRequest对象
func NewAlibabaAlihealthDrugKytGetbyrefentidRequest() *AlibabaAlihealthDrugKytGetbyrefentidAPIRequest {
	return &AlibabaAlihealthDrugKytGetbyrefentidAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugKytGetbyrefentidAPIRequest) Reset() {
	r._refEntId = ""
	r._destRefEntId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytGetbyrefentidAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.getbyrefentid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytGetbyrefentidAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugKytGetbyrefentidAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 接口调用企业的唯一标识（接口调用者）
func (r *AlibabaAlihealthDrugKytGetbyrefentidAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugKytGetbyrefentidAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetDestRefEntId is DestRefEntId Setter
// 准备要查询的企业唯一标识（返回该唯一标识企业的详细信息）
func (r *AlibabaAlihealthDrugKytGetbyrefentidAPIRequest) SetDestRefEntId(_destRefEntId string) error {
	r._destRefEntId = _destRefEntId
	r.Set("dest_ref_ent_id", _destRefEntId)
	return nil
}

// GetDestRefEntId DestRefEntId Getter
func (r AlibabaAlihealthDrugKytGetbyrefentidAPIRequest) GetDestRefEntId() string {
	return r._destRefEntId
}

var poolAlibabaAlihealthDrugKytGetbyrefentidAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugKytGetbyrefentidRequest()
	},
}

// GetAlibabaAlihealthDrugKytGetbyrefentidRequest 从 sync.Pool 获取 AlibabaAlihealthDrugKytGetbyrefentidAPIRequest
func GetAlibabaAlihealthDrugKytGetbyrefentidAPIRequest() *AlibabaAlihealthDrugKytGetbyrefentidAPIRequest {
	return poolAlibabaAlihealthDrugKytGetbyrefentidAPIRequest.Get().(*AlibabaAlihealthDrugKytGetbyrefentidAPIRequest)
}

// ReleaseAlibabaAlihealthDrugKytGetbyrefentidAPIRequest 将 AlibabaAlihealthDrugKytGetbyrefentidAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugKytGetbyrefentidAPIRequest(v *AlibabaAlihealthDrugKytGetbyrefentidAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugKytGetbyrefentidAPIRequest.Put(v)
}
