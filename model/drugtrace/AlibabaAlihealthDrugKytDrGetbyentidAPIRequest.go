package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytDrGetbyentidAPIRequest 多融通过企业ID得到一个企业的详细信息 API请求
// alibaba.alihealth.drug.kyt.dr.getbyentid
//
// 根据企业主键查看企业详细信息
type AlibabaAlihealthDrugKytDrGetbyentidAPIRequest struct {
	model.Params
	// 接口调用企业的唯一标识（接口调用者）
	_refEntId string
	// 准备要查询的企业ID（返回该企业ID的详细信息）
	_entId string
}

// NewAlibabaAlihealthDrugKytDrGetbyentidRequest 初始化AlibabaAlihealthDrugKytDrGetbyentidAPIRequest对象
func NewAlibabaAlihealthDrugKytDrGetbyentidRequest() *AlibabaAlihealthDrugKytDrGetbyentidAPIRequest {
	return &AlibabaAlihealthDrugKytDrGetbyentidAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugKytDrGetbyentidAPIRequest) Reset() {
	r._refEntId = ""
	r._entId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytDrGetbyentidAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.dr.getbyentid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytDrGetbyentidAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugKytDrGetbyentidAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 接口调用企业的唯一标识（接口调用者）
func (r *AlibabaAlihealthDrugKytDrGetbyentidAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugKytDrGetbyentidAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetEntId is EntId Setter
// 准备要查询的企业ID（返回该企业ID的详细信息）
func (r *AlibabaAlihealthDrugKytDrGetbyentidAPIRequest) SetEntId(_entId string) error {
	r._entId = _entId
	r.Set("ent_id", _entId)
	return nil
}

// GetEntId EntId Getter
func (r AlibabaAlihealthDrugKytDrGetbyentidAPIRequest) GetEntId() string {
	return r._entId
}

var poolAlibabaAlihealthDrugKytDrGetbyentidAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugKytDrGetbyentidRequest()
	},
}

// GetAlibabaAlihealthDrugKytDrGetbyentidRequest 从 sync.Pool 获取 AlibabaAlihealthDrugKytDrGetbyentidAPIRequest
func GetAlibabaAlihealthDrugKytDrGetbyentidAPIRequest() *AlibabaAlihealthDrugKytDrGetbyentidAPIRequest {
	return poolAlibabaAlihealthDrugKytDrGetbyentidAPIRequest.Get().(*AlibabaAlihealthDrugKytDrGetbyentidAPIRequest)
}

// ReleaseAlibabaAlihealthDrugKytDrGetbyentidAPIRequest 将 AlibabaAlihealthDrugKytDrGetbyentidAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugKytDrGetbyentidAPIRequest(v *AlibabaAlihealthDrugKytDrGetbyentidAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugKytDrGetbyentidAPIRequest.Put(v)
}
