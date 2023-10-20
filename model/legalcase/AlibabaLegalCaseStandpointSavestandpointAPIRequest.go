package legalcase

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLegalCaseStandpointSavestandpointAPIRequest 新增反馈口径 API请求
// alibaba.legal.case.standpoint.savestandpoint
//
// 新增反馈口径 ,从外部接受反馈的口径
type AlibabaLegalCaseStandpointSavestandpointAPIRequest struct {
	model.Params
	// 答辩口径
	_defenseCaliber string
	// 口径描述
	_standpointDesc string
	// 提交人
	_submitPeople string
	// 案件id
	_suitId int64
	// 委托id
	_entrustId int64
}

// NewAlibabaLegalCaseStandpointSavestandpointRequest 初始化AlibabaLegalCaseStandpointSavestandpointAPIRequest对象
func NewAlibabaLegalCaseStandpointSavestandpointRequest() *AlibabaLegalCaseStandpointSavestandpointAPIRequest {
	return &AlibabaLegalCaseStandpointSavestandpointAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaLegalCaseStandpointSavestandpointAPIRequest) Reset() {
	r._defenseCaliber = ""
	r._standpointDesc = ""
	r._submitPeople = ""
	r._suitId = 0
	r._entrustId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLegalCaseStandpointSavestandpointAPIRequest) GetApiMethodName() string {
	return "alibaba.legal.case.standpoint.savestandpoint"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLegalCaseStandpointSavestandpointAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaLegalCaseStandpointSavestandpointAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDefenseCaliber is DefenseCaliber Setter
// 答辩口径
func (r *AlibabaLegalCaseStandpointSavestandpointAPIRequest) SetDefenseCaliber(_defenseCaliber string) error {
	r._defenseCaliber = _defenseCaliber
	r.Set("defense_caliber", _defenseCaliber)
	return nil
}

// GetDefenseCaliber DefenseCaliber Getter
func (r AlibabaLegalCaseStandpointSavestandpointAPIRequest) GetDefenseCaliber() string {
	return r._defenseCaliber
}

// SetStandpointDesc is StandpointDesc Setter
// 口径描述
func (r *AlibabaLegalCaseStandpointSavestandpointAPIRequest) SetStandpointDesc(_standpointDesc string) error {
	r._standpointDesc = _standpointDesc
	r.Set("standpoint_desc", _standpointDesc)
	return nil
}

// GetStandpointDesc StandpointDesc Getter
func (r AlibabaLegalCaseStandpointSavestandpointAPIRequest) GetStandpointDesc() string {
	return r._standpointDesc
}

// SetSubmitPeople is SubmitPeople Setter
// 提交人
func (r *AlibabaLegalCaseStandpointSavestandpointAPIRequest) SetSubmitPeople(_submitPeople string) error {
	r._submitPeople = _submitPeople
	r.Set("submit_people", _submitPeople)
	return nil
}

// GetSubmitPeople SubmitPeople Getter
func (r AlibabaLegalCaseStandpointSavestandpointAPIRequest) GetSubmitPeople() string {
	return r._submitPeople
}

// SetSuitId is SuitId Setter
// 案件id
func (r *AlibabaLegalCaseStandpointSavestandpointAPIRequest) SetSuitId(_suitId int64) error {
	r._suitId = _suitId
	r.Set("suit_id", _suitId)
	return nil
}

// GetSuitId SuitId Getter
func (r AlibabaLegalCaseStandpointSavestandpointAPIRequest) GetSuitId() int64 {
	return r._suitId
}

// SetEntrustId is EntrustId Setter
// 委托id
func (r *AlibabaLegalCaseStandpointSavestandpointAPIRequest) SetEntrustId(_entrustId int64) error {
	r._entrustId = _entrustId
	r.Set("entrust_id", _entrustId)
	return nil
}

// GetEntrustId EntrustId Getter
func (r AlibabaLegalCaseStandpointSavestandpointAPIRequest) GetEntrustId() int64 {
	return r._entrustId
}

var poolAlibabaLegalCaseStandpointSavestandpointAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaLegalCaseStandpointSavestandpointRequest()
	},
}

// GetAlibabaLegalCaseStandpointSavestandpointRequest 从 sync.Pool 获取 AlibabaLegalCaseStandpointSavestandpointAPIRequest
func GetAlibabaLegalCaseStandpointSavestandpointAPIRequest() *AlibabaLegalCaseStandpointSavestandpointAPIRequest {
	return poolAlibabaLegalCaseStandpointSavestandpointAPIRequest.Get().(*AlibabaLegalCaseStandpointSavestandpointAPIRequest)
}

// ReleaseAlibabaLegalCaseStandpointSavestandpointAPIRequest 将 AlibabaLegalCaseStandpointSavestandpointAPIRequest 放入 sync.Pool
func ReleaseAlibabaLegalCaseStandpointSavestandpointAPIRequest(v *AlibabaLegalCaseStandpointSavestandpointAPIRequest) {
	v.Reset()
	poolAlibabaLegalCaseStandpointSavestandpointAPIRequest.Put(v)
}
