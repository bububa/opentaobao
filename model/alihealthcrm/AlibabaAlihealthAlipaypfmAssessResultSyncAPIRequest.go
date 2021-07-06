package alihealthcrm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthAlipaypfmAssessResultSyncAPIRequest 用户测评结果回传接口 API请求
// alibaba.alihealth.alipaypfm.assess.result.sync
//
// 用户测评结果回传接口
type AlibabaAlihealthAlipaypfmAssessResultSyncAPIRequest struct {
	model.Params
	// 测评类型
	_assessType string
	// 测评结果
	_assessResult string
	// 测评结果冗余字段
	_refrenceResult string
	// userId
	_userId int64
}

// NewAlibabaAlihealthAlipaypfmAssessResultSyncRequest 初始化AlibabaAlihealthAlipaypfmAssessResultSyncAPIRequest对象
func NewAlibabaAlihealthAlipaypfmAssessResultSyncRequest() *AlibabaAlihealthAlipaypfmAssessResultSyncAPIRequest {
	return &AlibabaAlihealthAlipaypfmAssessResultSyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthAlipaypfmAssessResultSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.alipaypfm.assess.result.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthAlipaypfmAssessResultSyncAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetAssessType is AssessType Setter
// 测评类型
func (r *AlibabaAlihealthAlipaypfmAssessResultSyncAPIRequest) SetAssessType(_assessType string) error {
	r._assessType = _assessType
	r.Set("assess_type", _assessType)
	return nil
}

// GetAssessType AssessType Getter
func (r AlibabaAlihealthAlipaypfmAssessResultSyncAPIRequest) GetAssessType() string {
	return r._assessType
}

// SetAssessResult is AssessResult Setter
// 测评结果
func (r *AlibabaAlihealthAlipaypfmAssessResultSyncAPIRequest) SetAssessResult(_assessResult string) error {
	r._assessResult = _assessResult
	r.Set("assess_result", _assessResult)
	return nil
}

// GetAssessResult AssessResult Getter
func (r AlibabaAlihealthAlipaypfmAssessResultSyncAPIRequest) GetAssessResult() string {
	return r._assessResult
}

// SetRefrenceResult is RefrenceResult Setter
// 测评结果冗余字段
func (r *AlibabaAlihealthAlipaypfmAssessResultSyncAPIRequest) SetRefrenceResult(_refrenceResult string) error {
	r._refrenceResult = _refrenceResult
	r.Set("refrence_result", _refrenceResult)
	return nil
}

// GetRefrenceResult RefrenceResult Getter
func (r AlibabaAlihealthAlipaypfmAssessResultSyncAPIRequest) GetRefrenceResult() string {
	return r._refrenceResult
}

// SetUserId is UserId Setter
// userId
func (r *AlibabaAlihealthAlipaypfmAssessResultSyncAPIRequest) SetUserId(_userId int64) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r AlibabaAlihealthAlipaypfmAssessResultSyncAPIRequest) GetUserId() int64 {
	return r._userId
}
