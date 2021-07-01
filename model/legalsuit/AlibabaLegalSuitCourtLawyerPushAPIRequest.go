package legalsuit

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLegalSuitCourtLawyerPushAPIRequest
推荐律师接口 API请求
alibaba.legal.suit.court.lawyer.push

为诉讼系统推荐律师 */
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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLegalSuitCourtLawyerPushAPIRequest) GetApiMethodName() string {
	return "alibaba.legal.suit.court.lawyer.push"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLegalSuitCourtLawyerPushAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is EntrustId Setter
// 委托ID
func (r *AlibabaLegalSuitCourtLawyerPushAPIRequest) SetEntrustId(_entrustId int64) error {
	r._entrustId = _entrustId
	r.Set("entrust_id", _entrustId)
	return nil
}

// Get EntrustId Getter
func (r AlibabaLegalSuitCourtLawyerPushAPIRequest) GetEntrustId() int64 {
	return r._entrustId
}

// Set is SuitId Setter
// 案件ID
func (r *AlibabaLegalSuitCourtLawyerPushAPIRequest) SetSuitId(_suitId int64) error {
	r._suitId = _suitId
	r.Set("suit_id", _suitId)
	return nil
}

// Get SuitId Getter
func (r AlibabaLegalSuitCourtLawyerPushAPIRequest) GetSuitId() int64 {
	return r._suitId
}

// Set is LawyersModel Setter
// 推荐律师模型
func (r *AlibabaLegalSuitCourtLawyerPushAPIRequest) SetLawyersModel(_lawyersModel *LawyersModel) error {
	r._lawyersModel = _lawyersModel
	r.Set("lawyers_model", _lawyersModel)
	return nil
}

// Get LawyersModel Getter
func (r AlibabaLegalSuitCourtLawyerPushAPIRequest) GetLawyersModel() *LawyersModel {
	return r._lawyersModel
}
