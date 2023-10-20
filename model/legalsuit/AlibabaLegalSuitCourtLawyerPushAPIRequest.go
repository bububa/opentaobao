package legalsuit

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalegalsuitcourtlawyerpushAPIRequest 推荐律师接口 API请求
// alibaba.legal.suit.court.lawyer.push
//
// 为诉讼系统推荐律师
type AlibabalegalsuitcourtlawyerpushAPIRequest struct {
	model.Params
	// 委托ID
	_entrustId int64
	// 案件ID
	_suitId int64
	// 推荐律师模型
	_lawyersModel *LawyersModel
}

// NewAlibabalegalsuitcourtlawyerpushRequest 初始化AlibabalegalsuitcourtlawyerpushAPIRequest对象
func NewAlibabalegalsuitcourtlawyerpushRequest() *AlibabalegalsuitcourtlawyerpushAPIRequest {
	return &AlibabalegalsuitcourtlawyerpushAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalegalsuitcourtlawyerpushAPIRequest) GetApiMethodName() string {
	return "alibaba.legal.suit.court.lawyer.push"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalegalsuitcourtlawyerpushAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalegalsuitcourtlawyerpushAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEntrustId is EntrustId Setter
// 委托ID
func (r *AlibabalegalsuitcourtlawyerpushAPIRequest) SetEntrustId(_entrustId int64) error {
	r._entrustId = _entrustId
	r.Set("entrust_id", _entrustId)
	return nil
}

// GetEntrustId EntrustId Getter
func (r AlibabalegalsuitcourtlawyerpushAPIRequest) GetEntrustId() int64 {
	return r._entrustId
}

// SetSuitId is SuitId Setter
// 案件ID
func (r *AlibabalegalsuitcourtlawyerpushAPIRequest) SetSuitId(_suitId int64) error {
	r._suitId = _suitId
	r.Set("suit_id", _suitId)
	return nil
}

// GetSuitId SuitId Getter
func (r AlibabalegalsuitcourtlawyerpushAPIRequest) GetSuitId() int64 {
	return r._suitId
}

// SetLawyersModel is LawyersModel Setter
// 推荐律师模型
func (r *AlibabalegalsuitcourtlawyerpushAPIRequest) SetLawyersModel(_lawyersModel *LawyersModel) error {
	r._lawyersModel = _lawyersModel
	r.Set("lawyers_model", _lawyersModel)
	return nil
}

// GetLawyersModel LawyersModel Getter
func (r AlibabalegalsuitcourtlawyerpushAPIRequest) GetLawyersModel() *LawyersModel {
	return r._lawyersModel
}
