package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugtraceTopYljgServiceGetenddateAPIRequest 获取服务截止日期 API请求
// alibaba.alihealth.drugtrace.top.yljg.service.getenddate
//
// 获取企业服务截止时间
type AlibabaAlihealthDrugtraceTopYljgServiceGetenddateAPIRequest struct {
	model.Params
	// 调用接口的企业ID
	_refEntId string
	// 药 行业线：传 1
	_business int64
	// 基础版：传 11
	_service int64
}

// NewAlibabaAlihealthDrugtraceTopYljgServiceGetenddateRequest 初始化AlibabaAlihealthDrugtraceTopYljgServiceGetenddateAPIRequest对象
func NewAlibabaAlihealthDrugtraceTopYljgServiceGetenddateRequest() *AlibabaAlihealthDrugtraceTopYljgServiceGetenddateAPIRequest {
	return &AlibabaAlihealthDrugtraceTopYljgServiceGetenddateAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugtraceTopYljgServiceGetenddateAPIRequest) Reset() {
	r._refEntId = ""
	r._business = 0
	r._service = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugtraceTopYljgServiceGetenddateAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drugtrace.top.yljg.service.getenddate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugtraceTopYljgServiceGetenddateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugtraceTopYljgServiceGetenddateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 调用接口的企业ID
func (r *AlibabaAlihealthDrugtraceTopYljgServiceGetenddateAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugtraceTopYljgServiceGetenddateAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetBusiness is Business Setter
// 药 行业线：传 1
func (r *AlibabaAlihealthDrugtraceTopYljgServiceGetenddateAPIRequest) SetBusiness(_business int64) error {
	r._business = _business
	r.Set("business", _business)
	return nil
}

// GetBusiness Business Getter
func (r AlibabaAlihealthDrugtraceTopYljgServiceGetenddateAPIRequest) GetBusiness() int64 {
	return r._business
}

// SetService is Service Setter
// 基础版：传 11
func (r *AlibabaAlihealthDrugtraceTopYljgServiceGetenddateAPIRequest) SetService(_service int64) error {
	r._service = _service
	r.Set("service", _service)
	return nil
}

// GetService Service Getter
func (r AlibabaAlihealthDrugtraceTopYljgServiceGetenddateAPIRequest) GetService() int64 {
	return r._service
}

var poolAlibabaAlihealthDrugtraceTopYljgServiceGetenddateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugtraceTopYljgServiceGetenddateRequest()
	},
}

// GetAlibabaAlihealthDrugtraceTopYljgServiceGetenddateRequest 从 sync.Pool 获取 AlibabaAlihealthDrugtraceTopYljgServiceGetenddateAPIRequest
func GetAlibabaAlihealthDrugtraceTopYljgServiceGetenddateAPIRequest() *AlibabaAlihealthDrugtraceTopYljgServiceGetenddateAPIRequest {
	return poolAlibabaAlihealthDrugtraceTopYljgServiceGetenddateAPIRequest.Get().(*AlibabaAlihealthDrugtraceTopYljgServiceGetenddateAPIRequest)
}

// ReleaseAlibabaAlihealthDrugtraceTopYljgServiceGetenddateAPIRequest 将 AlibabaAlihealthDrugtraceTopYljgServiceGetenddateAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugtraceTopYljgServiceGetenddateAPIRequest(v *AlibabaAlihealthDrugtraceTopYljgServiceGetenddateAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugtraceTopYljgServiceGetenddateAPIRequest.Put(v)
}
