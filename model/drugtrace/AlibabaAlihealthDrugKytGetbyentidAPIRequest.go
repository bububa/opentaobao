package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytGetbyentidAPIRequest 根据企业主键查看企业详细信息 API请求
// alibaba.alihealth.drug.kyt.getbyentid
//
// 根据企业主键查看企业详细信息
type AlibabaAlihealthDrugKytGetbyentidAPIRequest struct {
	model.Params
	// 接口调用企业的唯一标识（接口调用者）
	_refEntId string
	// 准备要查询的企业ID（返回该企业ID的详细信息）
	_entId string
}

// NewAlibabaAlihealthDrugKytGetbyentidRequest 初始化AlibabaAlihealthDrugKytGetbyentidAPIRequest对象
func NewAlibabaAlihealthDrugKytGetbyentidRequest() *AlibabaAlihealthDrugKytGetbyentidAPIRequest {
	return &AlibabaAlihealthDrugKytGetbyentidAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugKytGetbyentidAPIRequest) Reset() {
	r._refEntId = ""
	r._entId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytGetbyentidAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.getbyentid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytGetbyentidAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugKytGetbyentidAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 接口调用企业的唯一标识（接口调用者）
func (r *AlibabaAlihealthDrugKytGetbyentidAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugKytGetbyentidAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetEntId is EntId Setter
// 准备要查询的企业ID（返回该企业ID的详细信息）
func (r *AlibabaAlihealthDrugKytGetbyentidAPIRequest) SetEntId(_entId string) error {
	r._entId = _entId
	r.Set("ent_id", _entId)
	return nil
}

// GetEntId EntId Getter
func (r AlibabaAlihealthDrugKytGetbyentidAPIRequest) GetEntId() string {
	return r._entId
}

var poolAlibabaAlihealthDrugKytGetbyentidAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugKytGetbyentidRequest()
	},
}

// GetAlibabaAlihealthDrugKytGetbyentidRequest 从 sync.Pool 获取 AlibabaAlihealthDrugKytGetbyentidAPIRequest
func GetAlibabaAlihealthDrugKytGetbyentidAPIRequest() *AlibabaAlihealthDrugKytGetbyentidAPIRequest {
	return poolAlibabaAlihealthDrugKytGetbyentidAPIRequest.Get().(*AlibabaAlihealthDrugKytGetbyentidAPIRequest)
}

// ReleaseAlibabaAlihealthDrugKytGetbyentidAPIRequest 将 AlibabaAlihealthDrugKytGetbyentidAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugKytGetbyentidAPIRequest(v *AlibabaAlihealthDrugKytGetbyentidAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugKytGetbyentidAPIRequest.Put(v)
}
