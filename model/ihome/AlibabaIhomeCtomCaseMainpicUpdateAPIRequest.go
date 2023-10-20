package ihome

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaihomectomcasemainpicupdateAPIRequest 方案渲染图修改 API请求
// alibaba.ihome.ctom.case.mainpic.update
//
// 用于在门店工作台里的编辑器保存方案，由三维家后端调用阿里后端，保存方案信息
// 此接口只允许ihome业务使用，用于门店的编辑功能，只允许广东三维家信息科技有限公司一家公司调用，不适用于其他业务。
type AlibabaihomectomcasemainpicupdateAPIRequest struct {
	model.Params
	// 32位字符串
	_traceId string
	// 方案id
	_caseId string
	// 图片的地址
	_picUrl string
	// 图片类型
	_picType string
}

// NewAlibabaihomectomcasemainpicupdateRequest 初始化AlibabaihomectomcasemainpicupdateAPIRequest对象
func NewAlibabaihomectomcasemainpicupdateRequest() *AlibabaihomectomcasemainpicupdateAPIRequest {
	return &AlibabaihomectomcasemainpicupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaihomectomcasemainpicupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.ihome.ctom.case.mainpic.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaihomectomcasemainpicupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaihomectomcasemainpicupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTraceId is TraceId Setter
// 32位字符串
func (r *AlibabaihomectomcasemainpicupdateAPIRequest) SetTraceId(_traceId string) error {
	r._traceId = _traceId
	r.Set("trace_id", _traceId)
	return nil
}

// GetTraceId TraceId Getter
func (r AlibabaihomectomcasemainpicupdateAPIRequest) GetTraceId() string {
	return r._traceId
}

// SetCaseId is CaseId Setter
// 方案id
func (r *AlibabaihomectomcasemainpicupdateAPIRequest) SetCaseId(_caseId string) error {
	r._caseId = _caseId
	r.Set("case_id", _caseId)
	return nil
}

// GetCaseId CaseId Getter
func (r AlibabaihomectomcasemainpicupdateAPIRequest) GetCaseId() string {
	return r._caseId
}

// SetPicUrl is PicUrl Setter
// 图片的地址
func (r *AlibabaihomectomcasemainpicupdateAPIRequest) SetPicUrl(_picUrl string) error {
	r._picUrl = _picUrl
	r.Set("pic_url", _picUrl)
	return nil
}

// GetPicUrl PicUrl Getter
func (r AlibabaihomectomcasemainpicupdateAPIRequest) GetPicUrl() string {
	return r._picUrl
}

// SetPicType is PicType Setter
// 图片类型
func (r *AlibabaihomectomcasemainpicupdateAPIRequest) SetPicType(_picType string) error {
	r._picType = _picType
	r.Set("pic_type", _picType)
	return nil
}

// GetPicType PicType Getter
func (r AlibabaihomectomcasemainpicupdateAPIRequest) GetPicType() string {
	return r._picType
}
