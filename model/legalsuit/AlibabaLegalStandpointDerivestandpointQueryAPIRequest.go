package legalsuit

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLegalStandpointDerivestandpointQueryAPIRequest 查询衍生口径 API请求
// alibaba.legal.standpoint.derivestandpoint.query
//
// 查询衍生口径
type AlibabaLegalStandpointDerivestandpointQueryAPIRequest struct {
	model.Params
	// 系统标识
	_inputSystemCode string
	// 操作人
	_operateWorkNo string
	// 业务id
	_busiId string
	// 口径id
	_standpointId int64
	// 当前页
	_pageNum int64
	// 页大小
	_pageSize int64
}

// NewAlibabaLegalStandpointDerivestandpointQueryRequest 初始化AlibabaLegalStandpointDerivestandpointQueryAPIRequest对象
func NewAlibabaLegalStandpointDerivestandpointQueryRequest() *AlibabaLegalStandpointDerivestandpointQueryAPIRequest {
	return &AlibabaLegalStandpointDerivestandpointQueryAPIRequest{
		Params: model.NewParams(6),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaLegalStandpointDerivestandpointQueryAPIRequest) Reset() {
	r._inputSystemCode = ""
	r._operateWorkNo = ""
	r._busiId = ""
	r._standpointId = 0
	r._pageNum = 0
	r._pageSize = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLegalStandpointDerivestandpointQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.legal.standpoint.derivestandpoint.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLegalStandpointDerivestandpointQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaLegalStandpointDerivestandpointQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInputSystemCode is InputSystemCode Setter
// 系统标识
func (r *AlibabaLegalStandpointDerivestandpointQueryAPIRequest) SetInputSystemCode(_inputSystemCode string) error {
	r._inputSystemCode = _inputSystemCode
	r.Set("input_system_code", _inputSystemCode)
	return nil
}

// GetInputSystemCode InputSystemCode Getter
func (r AlibabaLegalStandpointDerivestandpointQueryAPIRequest) GetInputSystemCode() string {
	return r._inputSystemCode
}

// SetOperateWorkNo is OperateWorkNo Setter
// 操作人
func (r *AlibabaLegalStandpointDerivestandpointQueryAPIRequest) SetOperateWorkNo(_operateWorkNo string) error {
	r._operateWorkNo = _operateWorkNo
	r.Set("operate_work_no", _operateWorkNo)
	return nil
}

// GetOperateWorkNo OperateWorkNo Getter
func (r AlibabaLegalStandpointDerivestandpointQueryAPIRequest) GetOperateWorkNo() string {
	return r._operateWorkNo
}

// SetBusiId is BusiId Setter
// 业务id
func (r *AlibabaLegalStandpointDerivestandpointQueryAPIRequest) SetBusiId(_busiId string) error {
	r._busiId = _busiId
	r.Set("busi_id", _busiId)
	return nil
}

// GetBusiId BusiId Getter
func (r AlibabaLegalStandpointDerivestandpointQueryAPIRequest) GetBusiId() string {
	return r._busiId
}

// SetStandpointId is StandpointId Setter
// 口径id
func (r *AlibabaLegalStandpointDerivestandpointQueryAPIRequest) SetStandpointId(_standpointId int64) error {
	r._standpointId = _standpointId
	r.Set("standpoint_id", _standpointId)
	return nil
}

// GetStandpointId StandpointId Getter
func (r AlibabaLegalStandpointDerivestandpointQueryAPIRequest) GetStandpointId() int64 {
	return r._standpointId
}

// SetPageNum is PageNum Setter
// 当前页
func (r *AlibabaLegalStandpointDerivestandpointQueryAPIRequest) SetPageNum(_pageNum int64) error {
	r._pageNum = _pageNum
	r.Set("page_num", _pageNum)
	return nil
}

// GetPageNum PageNum Getter
func (r AlibabaLegalStandpointDerivestandpointQueryAPIRequest) GetPageNum() int64 {
	return r._pageNum
}

// SetPageSize is PageSize Setter
// 页大小
func (r *AlibabaLegalStandpointDerivestandpointQueryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlibabaLegalStandpointDerivestandpointQueryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

var poolAlibabaLegalStandpointDerivestandpointQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaLegalStandpointDerivestandpointQueryRequest()
	},
}

// GetAlibabaLegalStandpointDerivestandpointQueryRequest 从 sync.Pool 获取 AlibabaLegalStandpointDerivestandpointQueryAPIRequest
func GetAlibabaLegalStandpointDerivestandpointQueryAPIRequest() *AlibabaLegalStandpointDerivestandpointQueryAPIRequest {
	return poolAlibabaLegalStandpointDerivestandpointQueryAPIRequest.Get().(*AlibabaLegalStandpointDerivestandpointQueryAPIRequest)
}

// ReleaseAlibabaLegalStandpointDerivestandpointQueryAPIRequest 将 AlibabaLegalStandpointDerivestandpointQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaLegalStandpointDerivestandpointQueryAPIRequest(v *AlibabaLegalStandpointDerivestandpointQueryAPIRequest) {
	v.Reset()
	poolAlibabaLegalStandpointDerivestandpointQueryAPIRequest.Put(v)
}
