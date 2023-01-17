package legalsuit

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLegalStandpointGetrefAPIRequest 查询业务实体关联口径2 API请求
// alibaba.legal.standpoint.getref
//
// 口径查询
type AlibabaLegalStandpointGetrefAPIRequest struct {
	model.Params
	// 业务系统实体id
	_busiId string
	// 系统标识
	_inputSystemCode string
	// 页号
	_pageNum int64
	// 页面大小
	_pageSize int64
}

// NewAlibabaLegalStandpointGetrefRequest 初始化AlibabaLegalStandpointGetrefAPIRequest对象
func NewAlibabaLegalStandpointGetrefRequest() *AlibabaLegalStandpointGetrefAPIRequest {
	return &AlibabaLegalStandpointGetrefAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLegalStandpointGetrefAPIRequest) GetApiMethodName() string {
	return "alibaba.legal.standpoint.getref"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLegalStandpointGetrefAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaLegalStandpointGetrefAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBusiId is BusiId Setter
// 业务系统实体id
func (r *AlibabaLegalStandpointGetrefAPIRequest) SetBusiId(_busiId string) error {
	r._busiId = _busiId
	r.Set("busi_id", _busiId)
	return nil
}

// GetBusiId BusiId Getter
func (r AlibabaLegalStandpointGetrefAPIRequest) GetBusiId() string {
	return r._busiId
}

// SetInputSystemCode is InputSystemCode Setter
// 系统标识
func (r *AlibabaLegalStandpointGetrefAPIRequest) SetInputSystemCode(_inputSystemCode string) error {
	r._inputSystemCode = _inputSystemCode
	r.Set("input_system_code", _inputSystemCode)
	return nil
}

// GetInputSystemCode InputSystemCode Getter
func (r AlibabaLegalStandpointGetrefAPIRequest) GetInputSystemCode() string {
	return r._inputSystemCode
}

// SetPageNum is PageNum Setter
// 页号
func (r *AlibabaLegalStandpointGetrefAPIRequest) SetPageNum(_pageNum int64) error {
	r._pageNum = _pageNum
	r.Set("page_num", _pageNum)
	return nil
}

// GetPageNum PageNum Getter
func (r AlibabaLegalStandpointGetrefAPIRequest) GetPageNum() int64 {
	return r._pageNum
}

// SetPageSize is PageSize Setter
// 页面大小
func (r *AlibabaLegalStandpointGetrefAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlibabaLegalStandpointGetrefAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
