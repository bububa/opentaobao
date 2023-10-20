package legalcase

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalegalcasecommonnoticeAPIRequest 消息通知 API请求
// alibaba.legal.case.common.notice
//
// 同步通知给诉讼系统
type AlibabalegalcasecommonnoticeAPIRequest struct {
	model.Params
	// adminicular_evidence（补充证据）risk_early_warning（风险预警）
	_type string
	// 案件id
	_caseId int64
	// 委托id
	_entrustId int64
	// 消息模型
	_noticeModel *NoticeModel
}

// NewAlibabalegalcasecommonnoticeRequest 初始化AlibabalegalcasecommonnoticeAPIRequest对象
func NewAlibabalegalcasecommonnoticeRequest() *AlibabalegalcasecommonnoticeAPIRequest {
	return &AlibabalegalcasecommonnoticeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalegalcasecommonnoticeAPIRequest) GetApiMethodName() string {
	return "alibaba.legal.case.common.notice"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalegalcasecommonnoticeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalegalcasecommonnoticeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetType is Type Setter
// adminicular_evidence（补充证据）risk_early_warning（风险预警）
func (r *AlibabalegalcasecommonnoticeAPIRequest) SetType(_type string) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r AlibabalegalcasecommonnoticeAPIRequest) GetType() string {
	return r._type
}

// SetCaseId is CaseId Setter
// 案件id
func (r *AlibabalegalcasecommonnoticeAPIRequest) SetCaseId(_caseId int64) error {
	r._caseId = _caseId
	r.Set("case_id", _caseId)
	return nil
}

// GetCaseId CaseId Getter
func (r AlibabalegalcasecommonnoticeAPIRequest) GetCaseId() int64 {
	return r._caseId
}

// SetEntrustId is EntrustId Setter
// 委托id
func (r *AlibabalegalcasecommonnoticeAPIRequest) SetEntrustId(_entrustId int64) error {
	r._entrustId = _entrustId
	r.Set("entrust_id", _entrustId)
	return nil
}

// GetEntrustId EntrustId Getter
func (r AlibabalegalcasecommonnoticeAPIRequest) GetEntrustId() int64 {
	return r._entrustId
}

// SetNoticeModel is NoticeModel Setter
// 消息模型
func (r *AlibabalegalcasecommonnoticeAPIRequest) SetNoticeModel(_noticeModel *NoticeModel) error {
	r._noticeModel = _noticeModel
	r.Set("notice_model", _noticeModel)
	return nil
}

// GetNoticeModel NoticeModel Getter
func (r AlibabalegalcasecommonnoticeAPIRequest) GetNoticeModel() *NoticeModel {
	return r._noticeModel
}
