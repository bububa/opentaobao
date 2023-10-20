package legalsuit

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLegalSuitCourttimePushAPIRequest 开庭时间推送（带附件） API请求
// alibaba.legal.suit.courttime.push
//
// 开庭时间推送（带附件）
type AlibabaLegalSuitCourttimePushAPIRequest struct {
	model.Params
	// 开庭时间
	_courtTime string
	// 案件id
	_caseId int64
	// 委托id
	_entrustId int64
	// 文件对象
	_fileModel *FileModel
}

// NewAlibabaLegalSuitCourttimePushRequest 初始化AlibabaLegalSuitCourttimePushAPIRequest对象
func NewAlibabaLegalSuitCourttimePushRequest() *AlibabaLegalSuitCourttimePushAPIRequest {
	return &AlibabaLegalSuitCourttimePushAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaLegalSuitCourttimePushAPIRequest) Reset() {
	r._courtTime = ""
	r._caseId = 0
	r._entrustId = 0
	r._fileModel = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLegalSuitCourttimePushAPIRequest) GetApiMethodName() string {
	return "alibaba.legal.suit.courttime.push"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLegalSuitCourttimePushAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaLegalSuitCourttimePushAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCourtTime is CourtTime Setter
// 开庭时间
func (r *AlibabaLegalSuitCourttimePushAPIRequest) SetCourtTime(_courtTime string) error {
	r._courtTime = _courtTime
	r.Set("court_time", _courtTime)
	return nil
}

// GetCourtTime CourtTime Getter
func (r AlibabaLegalSuitCourttimePushAPIRequest) GetCourtTime() string {
	return r._courtTime
}

// SetCaseId is CaseId Setter
// 案件id
func (r *AlibabaLegalSuitCourttimePushAPIRequest) SetCaseId(_caseId int64) error {
	r._caseId = _caseId
	r.Set("case_id", _caseId)
	return nil
}

// GetCaseId CaseId Getter
func (r AlibabaLegalSuitCourttimePushAPIRequest) GetCaseId() int64 {
	return r._caseId
}

// SetEntrustId is EntrustId Setter
// 委托id
func (r *AlibabaLegalSuitCourttimePushAPIRequest) SetEntrustId(_entrustId int64) error {
	r._entrustId = _entrustId
	r.Set("entrust_id", _entrustId)
	return nil
}

// GetEntrustId EntrustId Getter
func (r AlibabaLegalSuitCourttimePushAPIRequest) GetEntrustId() int64 {
	return r._entrustId
}

// SetFileModel is FileModel Setter
// 文件对象
func (r *AlibabaLegalSuitCourttimePushAPIRequest) SetFileModel(_fileModel *FileModel) error {
	r._fileModel = _fileModel
	r.Set("file_model", _fileModel)
	return nil
}

// GetFileModel FileModel Getter
func (r AlibabaLegalSuitCourttimePushAPIRequest) GetFileModel() *FileModel {
	return r._fileModel
}

var poolAlibabaLegalSuitCourttimePushAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaLegalSuitCourttimePushRequest()
	},
}

// GetAlibabaLegalSuitCourttimePushRequest 从 sync.Pool 获取 AlibabaLegalSuitCourttimePushAPIRequest
func GetAlibabaLegalSuitCourttimePushAPIRequest() *AlibabaLegalSuitCourttimePushAPIRequest {
	return poolAlibabaLegalSuitCourttimePushAPIRequest.Get().(*AlibabaLegalSuitCourttimePushAPIRequest)
}

// ReleaseAlibabaLegalSuitCourttimePushAPIRequest 将 AlibabaLegalSuitCourttimePushAPIRequest 放入 sync.Pool
func ReleaseAlibabaLegalSuitCourttimePushAPIRequest(v *AlibabaLegalSuitCourttimePushAPIRequest) {
	v.Reset()
	poolAlibabaLegalSuitCourttimePushAPIRequest.Put(v)
}
