package legalsuit

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLegalStandpointCollectionstandpointQueryAPIRequest 查询收藏口径 API请求
// alibaba.legal.standpoint.collectionstandpoint.query
//
// 查询收藏口径
type AlibabaLegalStandpointCollectionstandpointQueryAPIRequest struct {
	model.Params
	// 系统标识
	_inputSystemCode string
	// 操作人
	_operateWorkNo string
	// 实体id
	_busiId string
	// 页数
	_pageNum int64
	// 页大小
	_pageSize int64
}

// NewAlibabaLegalStandpointCollectionstandpointQueryRequest 初始化AlibabaLegalStandpointCollectionstandpointQueryAPIRequest对象
func NewAlibabaLegalStandpointCollectionstandpointQueryRequest() *AlibabaLegalStandpointCollectionstandpointQueryAPIRequest {
	return &AlibabaLegalStandpointCollectionstandpointQueryAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaLegalStandpointCollectionstandpointQueryAPIRequest) Reset() {
	r._inputSystemCode = ""
	r._operateWorkNo = ""
	r._busiId = ""
	r._pageNum = 0
	r._pageSize = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLegalStandpointCollectionstandpointQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.legal.standpoint.collectionstandpoint.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLegalStandpointCollectionstandpointQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaLegalStandpointCollectionstandpointQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInputSystemCode is InputSystemCode Setter
// 系统标识
func (r *AlibabaLegalStandpointCollectionstandpointQueryAPIRequest) SetInputSystemCode(_inputSystemCode string) error {
	r._inputSystemCode = _inputSystemCode
	r.Set("input_system_code", _inputSystemCode)
	return nil
}

// GetInputSystemCode InputSystemCode Getter
func (r AlibabaLegalStandpointCollectionstandpointQueryAPIRequest) GetInputSystemCode() string {
	return r._inputSystemCode
}

// SetOperateWorkNo is OperateWorkNo Setter
// 操作人
func (r *AlibabaLegalStandpointCollectionstandpointQueryAPIRequest) SetOperateWorkNo(_operateWorkNo string) error {
	r._operateWorkNo = _operateWorkNo
	r.Set("operate_work_no", _operateWorkNo)
	return nil
}

// GetOperateWorkNo OperateWorkNo Getter
func (r AlibabaLegalStandpointCollectionstandpointQueryAPIRequest) GetOperateWorkNo() string {
	return r._operateWorkNo
}

// SetBusiId is BusiId Setter
// 实体id
func (r *AlibabaLegalStandpointCollectionstandpointQueryAPIRequest) SetBusiId(_busiId string) error {
	r._busiId = _busiId
	r.Set("busi_id", _busiId)
	return nil
}

// GetBusiId BusiId Getter
func (r AlibabaLegalStandpointCollectionstandpointQueryAPIRequest) GetBusiId() string {
	return r._busiId
}

// SetPageNum is PageNum Setter
// 页数
func (r *AlibabaLegalStandpointCollectionstandpointQueryAPIRequest) SetPageNum(_pageNum int64) error {
	r._pageNum = _pageNum
	r.Set("page_num", _pageNum)
	return nil
}

// GetPageNum PageNum Getter
func (r AlibabaLegalStandpointCollectionstandpointQueryAPIRequest) GetPageNum() int64 {
	return r._pageNum
}

// SetPageSize is PageSize Setter
// 页大小
func (r *AlibabaLegalStandpointCollectionstandpointQueryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlibabaLegalStandpointCollectionstandpointQueryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

var poolAlibabaLegalStandpointCollectionstandpointQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaLegalStandpointCollectionstandpointQueryRequest()
	},
}

// GetAlibabaLegalStandpointCollectionstandpointQueryRequest 从 sync.Pool 获取 AlibabaLegalStandpointCollectionstandpointQueryAPIRequest
func GetAlibabaLegalStandpointCollectionstandpointQueryAPIRequest() *AlibabaLegalStandpointCollectionstandpointQueryAPIRequest {
	return poolAlibabaLegalStandpointCollectionstandpointQueryAPIRequest.Get().(*AlibabaLegalStandpointCollectionstandpointQueryAPIRequest)
}

// ReleaseAlibabaLegalStandpointCollectionstandpointQueryAPIRequest 将 AlibabaLegalStandpointCollectionstandpointQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaLegalStandpointCollectionstandpointQueryAPIRequest(v *AlibabaLegalStandpointCollectionstandpointQueryAPIRequest) {
	v.Reset()
	poolAlibabaLegalStandpointCollectionstandpointQueryAPIRequest.Put(v)
}
