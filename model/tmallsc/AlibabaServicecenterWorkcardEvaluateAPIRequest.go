package tmallsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaservicecenterworkcardevaluateAPIRequest 服务商售后鉴定服务 API请求
// alibaba.servicecenter.workcard.evaluate
//
// 服务商售后鉴定服务,提供给服务商针对售后场景上门鉴定服务，鉴定成功则服务商完成履约，鉴定失败则取消工单
type AlibabaservicecenterworkcardevaluateAPIRequest struct {
	model.Params
	// 鉴定结果图片列表
	_picUrlList string
	// 鉴定不通过时的原因编码
	_failReason string
	// 扩展参数，品牌回传鉴定附加信息
	_extendInfo string
	// 工单ID
	_workcardId int64
	// 是否鉴定通过
	_passEvaluation bool
}

// NewAlibabaservicecenterworkcardevaluateRequest 初始化AlibabaservicecenterworkcardevaluateAPIRequest对象
func NewAlibabaservicecenterworkcardevaluateRequest() *AlibabaservicecenterworkcardevaluateAPIRequest {
	return &AlibabaservicecenterworkcardevaluateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaservicecenterworkcardevaluateAPIRequest) GetApiMethodName() string {
	return "alibaba.servicecenter.workcard.evaluate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaservicecenterworkcardevaluateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaservicecenterworkcardevaluateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPicUrlList is PicUrlList Setter
// 鉴定结果图片列表
func (r *AlibabaservicecenterworkcardevaluateAPIRequest) SetPicUrlList(_picUrlList string) error {
	r._picUrlList = _picUrlList
	r.Set("pic_url_list", _picUrlList)
	return nil
}

// GetPicUrlList PicUrlList Getter
func (r AlibabaservicecenterworkcardevaluateAPIRequest) GetPicUrlList() string {
	return r._picUrlList
}

// SetFailReason is FailReason Setter
// 鉴定不通过时的原因编码
func (r *AlibabaservicecenterworkcardevaluateAPIRequest) SetFailReason(_failReason string) error {
	r._failReason = _failReason
	r.Set("fail_reason", _failReason)
	return nil
}

// GetFailReason FailReason Getter
func (r AlibabaservicecenterworkcardevaluateAPIRequest) GetFailReason() string {
	return r._failReason
}

// SetExtendInfo is ExtendInfo Setter
// 扩展参数，品牌回传鉴定附加信息
func (r *AlibabaservicecenterworkcardevaluateAPIRequest) SetExtendInfo(_extendInfo string) error {
	r._extendInfo = _extendInfo
	r.Set("extend_info", _extendInfo)
	return nil
}

// GetExtendInfo ExtendInfo Getter
func (r AlibabaservicecenterworkcardevaluateAPIRequest) GetExtendInfo() string {
	return r._extendInfo
}

// SetWorkcardId is WorkcardId Setter
// 工单ID
func (r *AlibabaservicecenterworkcardevaluateAPIRequest) SetWorkcardId(_workcardId int64) error {
	r._workcardId = _workcardId
	r.Set("workcard_id", _workcardId)
	return nil
}

// GetWorkcardId WorkcardId Getter
func (r AlibabaservicecenterworkcardevaluateAPIRequest) GetWorkcardId() int64 {
	return r._workcardId
}

// SetPassEvaluation is PassEvaluation Setter
// 是否鉴定通过
func (r *AlibabaservicecenterworkcardevaluateAPIRequest) SetPassEvaluation(_passEvaluation bool) error {
	r._passEvaluation = _passEvaluation
	r.Set("pass_evaluation", _passEvaluation)
	return nil
}

// GetPassEvaluation PassEvaluation Getter
func (r AlibabaservicecenterworkcardevaluateAPIRequest) GetPassEvaluation() bool {
	return r._passEvaluation
}
