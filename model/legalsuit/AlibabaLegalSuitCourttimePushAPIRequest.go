package legalsuit

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalegalsuitcourttimepushAPIRequest 开庭时间推送（带附件） API请求
// alibaba.legal.suit.courttime.push
//
// 开庭时间推送（带附件）
type AlibabalegalsuitcourttimepushAPIRequest struct {
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

// NewAlibabalegalsuitcourttimepushRequest 初始化AlibabalegalsuitcourttimepushAPIRequest对象
func NewAlibabalegalsuitcourttimepushRequest() *AlibabalegalsuitcourttimepushAPIRequest {
	return &AlibabalegalsuitcourttimepushAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalegalsuitcourttimepushAPIRequest) GetApiMethodName() string {
	return "alibaba.legal.suit.courttime.push"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalegalsuitcourttimepushAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalegalsuitcourttimepushAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCourtTime is CourtTime Setter
// 开庭时间
func (r *AlibabalegalsuitcourttimepushAPIRequest) SetCourtTime(_courtTime string) error {
	r._courtTime = _courtTime
	r.Set("court_time", _courtTime)
	return nil
}

// GetCourtTime CourtTime Getter
func (r AlibabalegalsuitcourttimepushAPIRequest) GetCourtTime() string {
	return r._courtTime
}

// SetCaseId is CaseId Setter
// 案件id
func (r *AlibabalegalsuitcourttimepushAPIRequest) SetCaseId(_caseId int64) error {
	r._caseId = _caseId
	r.Set("case_id", _caseId)
	return nil
}

// GetCaseId CaseId Getter
func (r AlibabalegalsuitcourttimepushAPIRequest) GetCaseId() int64 {
	return r._caseId
}

// SetEntrustId is EntrustId Setter
// 委托id
func (r *AlibabalegalsuitcourttimepushAPIRequest) SetEntrustId(_entrustId int64) error {
	r._entrustId = _entrustId
	r.Set("entrust_id", _entrustId)
	return nil
}

// GetEntrustId EntrustId Getter
func (r AlibabalegalsuitcourttimepushAPIRequest) GetEntrustId() int64 {
	return r._entrustId
}

// SetFileModel is FileModel Setter
// 文件对象
func (r *AlibabalegalsuitcourttimepushAPIRequest) SetFileModel(_fileModel *FileModel) error {
	r._fileModel = _fileModel
	r.Set("file_model", _fileModel)
	return nil
}

// GetFileModel FileModel Getter
func (r AlibabalegalsuitcourttimepushAPIRequest) GetFileModel() *FileModel {
	return r._fileModel
}
