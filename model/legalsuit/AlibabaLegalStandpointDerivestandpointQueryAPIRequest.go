package legalsuit

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalegalstandpointderivestandpointqueryAPIRequest 查询衍生口径 API请求
// alibaba.legal.standpoint.derivestandpoint.query
//
// 查询衍生口径
type AlibabalegalstandpointderivestandpointqueryAPIRequest struct {
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

// NewAlibabalegalstandpointderivestandpointqueryRequest 初始化AlibabalegalstandpointderivestandpointqueryAPIRequest对象
func NewAlibabalegalstandpointderivestandpointqueryRequest() *AlibabalegalstandpointderivestandpointqueryAPIRequest {
	return &AlibabalegalstandpointderivestandpointqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalegalstandpointderivestandpointqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.legal.standpoint.derivestandpoint.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalegalstandpointderivestandpointqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalegalstandpointderivestandpointqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInputSystemCode is InputSystemCode Setter
// 系统标识
func (r *AlibabalegalstandpointderivestandpointqueryAPIRequest) SetInputSystemCode(_inputSystemCode string) error {
	r._inputSystemCode = _inputSystemCode
	r.Set("input_system_code", _inputSystemCode)
	return nil
}

// GetInputSystemCode InputSystemCode Getter
func (r AlibabalegalstandpointderivestandpointqueryAPIRequest) GetInputSystemCode() string {
	return r._inputSystemCode
}

// SetOperateWorkNo is OperateWorkNo Setter
// 操作人
func (r *AlibabalegalstandpointderivestandpointqueryAPIRequest) SetOperateWorkNo(_operateWorkNo string) error {
	r._operateWorkNo = _operateWorkNo
	r.Set("operate_work_no", _operateWorkNo)
	return nil
}

// GetOperateWorkNo OperateWorkNo Getter
func (r AlibabalegalstandpointderivestandpointqueryAPIRequest) GetOperateWorkNo() string {
	return r._operateWorkNo
}

// SetBusiId is BusiId Setter
// 业务id
func (r *AlibabalegalstandpointderivestandpointqueryAPIRequest) SetBusiId(_busiId string) error {
	r._busiId = _busiId
	r.Set("busi_id", _busiId)
	return nil
}

// GetBusiId BusiId Getter
func (r AlibabalegalstandpointderivestandpointqueryAPIRequest) GetBusiId() string {
	return r._busiId
}

// SetStandpointId is StandpointId Setter
// 口径id
func (r *AlibabalegalstandpointderivestandpointqueryAPIRequest) SetStandpointId(_standpointId int64) error {
	r._standpointId = _standpointId
	r.Set("standpoint_id", _standpointId)
	return nil
}

// GetStandpointId StandpointId Getter
func (r AlibabalegalstandpointderivestandpointqueryAPIRequest) GetStandpointId() int64 {
	return r._standpointId
}

// SetPageNum is PageNum Setter
// 当前页
func (r *AlibabalegalstandpointderivestandpointqueryAPIRequest) SetPageNum(_pageNum int64) error {
	r._pageNum = _pageNum
	r.Set("page_num", _pageNum)
	return nil
}

// GetPageNum PageNum Getter
func (r AlibabalegalstandpointderivestandpointqueryAPIRequest) GetPageNum() int64 {
	return r._pageNum
}

// SetPageSize is PageSize Setter
// 页大小
func (r *AlibabalegalstandpointderivestandpointqueryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlibabalegalstandpointderivestandpointqueryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
