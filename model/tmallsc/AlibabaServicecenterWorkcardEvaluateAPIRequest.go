package tmallsc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaServicecenterWorkcardEvaluateAPIRequest 服务商售后鉴定服务 API请求
// alibaba.servicecenter.workcard.evaluate
//
// 服务商售后鉴定服务,提供给服务商针对售后场景上门鉴定服务，鉴定成功则服务商完成履约，鉴定失败则取消工单
type AlibabaServicecenterWorkcardEvaluateAPIRequest struct {
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

// NewAlibabaServicecenterWorkcardEvaluateRequest 初始化AlibabaServicecenterWorkcardEvaluateAPIRequest对象
func NewAlibabaServicecenterWorkcardEvaluateRequest() *AlibabaServicecenterWorkcardEvaluateAPIRequest {
	return &AlibabaServicecenterWorkcardEvaluateAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaServicecenterWorkcardEvaluateAPIRequest) Reset() {
	r._picUrlList = ""
	r._failReason = ""
	r._extendInfo = ""
	r._workcardId = 0
	r._passEvaluation = false
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaServicecenterWorkcardEvaluateAPIRequest) GetApiMethodName() string {
	return "alibaba.servicecenter.workcard.evaluate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaServicecenterWorkcardEvaluateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaServicecenterWorkcardEvaluateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPicUrlList is PicUrlList Setter
// 鉴定结果图片列表
func (r *AlibabaServicecenterWorkcardEvaluateAPIRequest) SetPicUrlList(_picUrlList string) error {
	r._picUrlList = _picUrlList
	r.Set("pic_url_list", _picUrlList)
	return nil
}

// GetPicUrlList PicUrlList Getter
func (r AlibabaServicecenterWorkcardEvaluateAPIRequest) GetPicUrlList() string {
	return r._picUrlList
}

// SetFailReason is FailReason Setter
// 鉴定不通过时的原因编码
func (r *AlibabaServicecenterWorkcardEvaluateAPIRequest) SetFailReason(_failReason string) error {
	r._failReason = _failReason
	r.Set("fail_reason", _failReason)
	return nil
}

// GetFailReason FailReason Getter
func (r AlibabaServicecenterWorkcardEvaluateAPIRequest) GetFailReason() string {
	return r._failReason
}

// SetExtendInfo is ExtendInfo Setter
// 扩展参数，品牌回传鉴定附加信息
func (r *AlibabaServicecenterWorkcardEvaluateAPIRequest) SetExtendInfo(_extendInfo string) error {
	r._extendInfo = _extendInfo
	r.Set("extend_info", _extendInfo)
	return nil
}

// GetExtendInfo ExtendInfo Getter
func (r AlibabaServicecenterWorkcardEvaluateAPIRequest) GetExtendInfo() string {
	return r._extendInfo
}

// SetWorkcardId is WorkcardId Setter
// 工单ID
func (r *AlibabaServicecenterWorkcardEvaluateAPIRequest) SetWorkcardId(_workcardId int64) error {
	r._workcardId = _workcardId
	r.Set("workcard_id", _workcardId)
	return nil
}

// GetWorkcardId WorkcardId Getter
func (r AlibabaServicecenterWorkcardEvaluateAPIRequest) GetWorkcardId() int64 {
	return r._workcardId
}

// SetPassEvaluation is PassEvaluation Setter
// 是否鉴定通过
func (r *AlibabaServicecenterWorkcardEvaluateAPIRequest) SetPassEvaluation(_passEvaluation bool) error {
	r._passEvaluation = _passEvaluation
	r.Set("pass_evaluation", _passEvaluation)
	return nil
}

// GetPassEvaluation PassEvaluation Getter
func (r AlibabaServicecenterWorkcardEvaluateAPIRequest) GetPassEvaluation() bool {
	return r._passEvaluation
}

var poolAlibabaServicecenterWorkcardEvaluateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaServicecenterWorkcardEvaluateRequest()
	},
}

// GetAlibabaServicecenterWorkcardEvaluateRequest 从 sync.Pool 获取 AlibabaServicecenterWorkcardEvaluateAPIRequest
func GetAlibabaServicecenterWorkcardEvaluateAPIRequest() *AlibabaServicecenterWorkcardEvaluateAPIRequest {
	return poolAlibabaServicecenterWorkcardEvaluateAPIRequest.Get().(*AlibabaServicecenterWorkcardEvaluateAPIRequest)
}

// ReleaseAlibabaServicecenterWorkcardEvaluateAPIRequest 将 AlibabaServicecenterWorkcardEvaluateAPIRequest 放入 sync.Pool
func ReleaseAlibabaServicecenterWorkcardEvaluateAPIRequest(v *AlibabaServicecenterWorkcardEvaluateAPIRequest) {
	v.Reset()
	poolAlibabaServicecenterWorkcardEvaluateAPIRequest.Put(v)
}
