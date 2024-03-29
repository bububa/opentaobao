package legalsuit

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLegalSuitCourtLawyerPushAPIRequest 推荐律师接口 API请求
// alibaba.legal.suit.court.lawyer.push
//
// 为诉讼系统推荐律师
type AlibabaLegalSuitCourtLawyerPushAPIRequest struct {
	model.Params
	// 委托ID
	_entrustId int64
	// 案件ID
	_suitId int64
	// 推荐律师模型
	_lawyersModel *LawyersModel
}

// NewAlibabaLegalSuitCourtLawyerPushRequest 初始化AlibabaLegalSuitCourtLawyerPushAPIRequest对象
func NewAlibabaLegalSuitCourtLawyerPushRequest() *AlibabaLegalSuitCourtLawyerPushAPIRequest {
	return &AlibabaLegalSuitCourtLawyerPushAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaLegalSuitCourtLawyerPushAPIRequest) Reset() {
	r._entrustId = 0
	r._suitId = 0
	r._lawyersModel = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLegalSuitCourtLawyerPushAPIRequest) GetApiMethodName() string {
	return "alibaba.legal.suit.court.lawyer.push"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLegalSuitCourtLawyerPushAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaLegalSuitCourtLawyerPushAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEntrustId is EntrustId Setter
// 委托ID
func (r *AlibabaLegalSuitCourtLawyerPushAPIRequest) SetEntrustId(_entrustId int64) error {
	r._entrustId = _entrustId
	r.Set("entrust_id", _entrustId)
	return nil
}

// GetEntrustId EntrustId Getter
func (r AlibabaLegalSuitCourtLawyerPushAPIRequest) GetEntrustId() int64 {
	return r._entrustId
}

// SetSuitId is SuitId Setter
// 案件ID
func (r *AlibabaLegalSuitCourtLawyerPushAPIRequest) SetSuitId(_suitId int64) error {
	r._suitId = _suitId
	r.Set("suit_id", _suitId)
	return nil
}

// GetSuitId SuitId Getter
func (r AlibabaLegalSuitCourtLawyerPushAPIRequest) GetSuitId() int64 {
	return r._suitId
}

// SetLawyersModel is LawyersModel Setter
// 推荐律师模型
func (r *AlibabaLegalSuitCourtLawyerPushAPIRequest) SetLawyersModel(_lawyersModel *LawyersModel) error {
	r._lawyersModel = _lawyersModel
	r.Set("lawyers_model", _lawyersModel)
	return nil
}

// GetLawyersModel LawyersModel Getter
func (r AlibabaLegalSuitCourtLawyerPushAPIRequest) GetLawyersModel() *LawyersModel {
	return r._lawyersModel
}

var poolAlibabaLegalSuitCourtLawyerPushAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaLegalSuitCourtLawyerPushRequest()
	},
}

// GetAlibabaLegalSuitCourtLawyerPushRequest 从 sync.Pool 获取 AlibabaLegalSuitCourtLawyerPushAPIRequest
func GetAlibabaLegalSuitCourtLawyerPushAPIRequest() *AlibabaLegalSuitCourtLawyerPushAPIRequest {
	return poolAlibabaLegalSuitCourtLawyerPushAPIRequest.Get().(*AlibabaLegalSuitCourtLawyerPushAPIRequest)
}

// ReleaseAlibabaLegalSuitCourtLawyerPushAPIRequest 将 AlibabaLegalSuitCourtLawyerPushAPIRequest 放入 sync.Pool
func ReleaseAlibabaLegalSuitCourtLawyerPushAPIRequest(v *AlibabaLegalSuitCourtLawyerPushAPIRequest) {
	v.Reset()
	poolAlibabaLegalSuitCourtLawyerPushAPIRequest.Put(v)
}
