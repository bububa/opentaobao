package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidAPIRequest 根据企业唯一标识查看企业详细信息 API请求
// alibaba.alihealth.drugtrace.top.yljg.query.getbyrefentid
//
// 根据企业唯一标识查看企业详细信息
type AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidAPIRequest struct {
	model.Params
	// 接口调用企业的唯一标识（接口调用者）
	_refEntId string
	// 准备要查询的企业唯一标识（返回该唯一标识企业的详细信息）
	_destRefEntId string
}

// NewAlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidRequest 初始化AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidAPIRequest对象
func NewAlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidRequest() *AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidAPIRequest {
	return &AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidAPIRequest) Reset() {
	r._refEntId = ""
	r._destRefEntId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drugtrace.top.yljg.query.getbyrefentid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 接口调用企业的唯一标识（接口调用者）
func (r *AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetDestRefEntId is DestRefEntId Setter
// 准备要查询的企业唯一标识（返回该唯一标识企业的详细信息）
func (r *AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidAPIRequest) SetDestRefEntId(_destRefEntId string) error {
	r._destRefEntId = _destRefEntId
	r.Set("dest_ref_ent_id", _destRefEntId)
	return nil
}

// GetDestRefEntId DestRefEntId Getter
func (r AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidAPIRequest) GetDestRefEntId() string {
	return r._destRefEntId
}

var poolAlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidRequest()
	},
}

// GetAlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidRequest 从 sync.Pool 获取 AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidAPIRequest
func GetAlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidAPIRequest() *AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidAPIRequest {
	return poolAlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidAPIRequest.Get().(*AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidAPIRequest)
}

// ReleaseAlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidAPIRequest 将 AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidAPIRequest(v *AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidAPIRequest.Put(v)
}
