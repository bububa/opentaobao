package legalsuit

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalegalstandpointcollectionstandpointqueryAPIRequest 查询收藏口径 API请求
// alibaba.legal.standpoint.collectionstandpoint.query
//
// 查询收藏口径
type AlibabalegalstandpointcollectionstandpointqueryAPIRequest struct {
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

// NewAlibabalegalstandpointcollectionstandpointqueryRequest 初始化AlibabalegalstandpointcollectionstandpointqueryAPIRequest对象
func NewAlibabalegalstandpointcollectionstandpointqueryRequest() *AlibabalegalstandpointcollectionstandpointqueryAPIRequest {
	return &AlibabalegalstandpointcollectionstandpointqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalegalstandpointcollectionstandpointqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.legal.standpoint.collectionstandpoint.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalegalstandpointcollectionstandpointqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalegalstandpointcollectionstandpointqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInputSystemCode is InputSystemCode Setter
// 系统标识
func (r *AlibabalegalstandpointcollectionstandpointqueryAPIRequest) SetInputSystemCode(_inputSystemCode string) error {
	r._inputSystemCode = _inputSystemCode
	r.Set("input_system_code", _inputSystemCode)
	return nil
}

// GetInputSystemCode InputSystemCode Getter
func (r AlibabalegalstandpointcollectionstandpointqueryAPIRequest) GetInputSystemCode() string {
	return r._inputSystemCode
}

// SetOperateWorkNo is OperateWorkNo Setter
// 操作人
func (r *AlibabalegalstandpointcollectionstandpointqueryAPIRequest) SetOperateWorkNo(_operateWorkNo string) error {
	r._operateWorkNo = _operateWorkNo
	r.Set("operate_work_no", _operateWorkNo)
	return nil
}

// GetOperateWorkNo OperateWorkNo Getter
func (r AlibabalegalstandpointcollectionstandpointqueryAPIRequest) GetOperateWorkNo() string {
	return r._operateWorkNo
}

// SetBusiId is BusiId Setter
// 实体id
func (r *AlibabalegalstandpointcollectionstandpointqueryAPIRequest) SetBusiId(_busiId string) error {
	r._busiId = _busiId
	r.Set("busi_id", _busiId)
	return nil
}

// GetBusiId BusiId Getter
func (r AlibabalegalstandpointcollectionstandpointqueryAPIRequest) GetBusiId() string {
	return r._busiId
}

// SetPageNum is PageNum Setter
// 页数
func (r *AlibabalegalstandpointcollectionstandpointqueryAPIRequest) SetPageNum(_pageNum int64) error {
	r._pageNum = _pageNum
	r.Set("page_num", _pageNum)
	return nil
}

// GetPageNum PageNum Getter
func (r AlibabalegalstandpointcollectionstandpointqueryAPIRequest) GetPageNum() int64 {
	return r._pageNum
}

// SetPageSize is PageSize Setter
// 页大小
func (r *AlibabalegalstandpointcollectionstandpointqueryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlibabalegalstandpointcollectionstandpointqueryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
