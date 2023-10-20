package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidAPIRequest 根据企业唯一标识查看企业详细信息 API请求
// alibaba.alihealth.drugtrace.top.lsyd.query.getbyrefentid
//
// 根据企业唯一标识查看企业详细信息
type AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidAPIRequest struct {
	model.Params
	// 接口调用企业的唯一标识（接口调用者）
	_refEntId string
	// 准备要查询的企业唯一标识（返回该唯一标识企业的详细信息）
	_destRefEntId string
}

// NewAlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidRequest 初始化AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidAPIRequest对象
func NewAlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidRequest() *AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidAPIRequest {
	return &AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidAPIRequest) Reset() {
	r._refEntId = ""
	r._destRefEntId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drugtrace.top.lsyd.query.getbyrefentid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 接口调用企业的唯一标识（接口调用者）
func (r *AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetDestRefEntId is DestRefEntId Setter
// 准备要查询的企业唯一标识（返回该唯一标识企业的详细信息）
func (r *AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidAPIRequest) SetDestRefEntId(_destRefEntId string) error {
	r._destRefEntId = _destRefEntId
	r.Set("dest_ref_ent_id", _destRefEntId)
	return nil
}

// GetDestRefEntId DestRefEntId Getter
func (r AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidAPIRequest) GetDestRefEntId() string {
	return r._destRefEntId
}

var poolAlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidRequest()
	},
}

// GetAlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidRequest 从 sync.Pool 获取 AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidAPIRequest
func GetAlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidAPIRequest() *AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidAPIRequest {
	return poolAlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidAPIRequest.Get().(*AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidAPIRequest)
}

// ReleaseAlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidAPIRequest 将 AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidAPIRequest(v *AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidAPIRequest.Put(v)
}
