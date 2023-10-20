package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugtraceTopLsydServiceGetenddateAPIRequest 获取企业服务截止时间 API请求
// alibaba.alihealth.drugtrace.top.lsyd.service.getenddate
//
// 获取企业服务截止时间
type AlibabaAlihealthDrugtraceTopLsydServiceGetenddateAPIRequest struct {
	model.Params
	// 调用接口的企业ID
	_refEntId string
	// 药 行业线：传 1
	_business int64
	// 基础版：传 11
	_service int64
}

// NewAlibabaAlihealthDrugtraceTopLsydServiceGetenddateRequest 初始化AlibabaAlihealthDrugtraceTopLsydServiceGetenddateAPIRequest对象
func NewAlibabaAlihealthDrugtraceTopLsydServiceGetenddateRequest() *AlibabaAlihealthDrugtraceTopLsydServiceGetenddateAPIRequest {
	return &AlibabaAlihealthDrugtraceTopLsydServiceGetenddateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugtraceTopLsydServiceGetenddateAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drugtrace.top.lsyd.service.getenddate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugtraceTopLsydServiceGetenddateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugtraceTopLsydServiceGetenddateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 调用接口的企业ID
func (r *AlibabaAlihealthDrugtraceTopLsydServiceGetenddateAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugtraceTopLsydServiceGetenddateAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetBusiness is Business Setter
// 药 行业线：传 1
func (r *AlibabaAlihealthDrugtraceTopLsydServiceGetenddateAPIRequest) SetBusiness(_business int64) error {
	r._business = _business
	r.Set("business", _business)
	return nil
}

// GetBusiness Business Getter
func (r AlibabaAlihealthDrugtraceTopLsydServiceGetenddateAPIRequest) GetBusiness() int64 {
	return r._business
}

// SetService is Service Setter
// 基础版：传 11
func (r *AlibabaAlihealthDrugtraceTopLsydServiceGetenddateAPIRequest) SetService(_service int64) error {
	r._service = _service
	r.Set("service", _service)
	return nil
}

// GetService Service Getter
func (r AlibabaAlihealthDrugtraceTopLsydServiceGetenddateAPIRequest) GetService() int64 {
	return r._service
}
