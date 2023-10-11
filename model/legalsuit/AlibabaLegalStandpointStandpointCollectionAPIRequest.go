package legalsuit

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLegalStandpointStandpointCollectionAPIRequest 收藏|取消收藏 API请求
// alibaba.legal.standpoint.standpoint.collection
//
// 收藏|取消收藏
type AlibabaLegalStandpointStandpointCollectionAPIRequest struct {
	model.Params
	// 系统标识
	_inputSystemCode string
	// 操作人
	_operateWorkNo string
	// 操作人名称
	_operateName string
	// 收藏
	_collectionFlag string
	// 口径id
	_standpointId int64
}

// NewAlibabaLegalStandpointStandpointCollectionRequest 初始化AlibabaLegalStandpointStandpointCollectionAPIRequest对象
func NewAlibabaLegalStandpointStandpointCollectionRequest() *AlibabaLegalStandpointStandpointCollectionAPIRequest {
	return &AlibabaLegalStandpointStandpointCollectionAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLegalStandpointStandpointCollectionAPIRequest) GetApiMethodName() string {
	return "alibaba.legal.standpoint.standpoint.collection"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLegalStandpointStandpointCollectionAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaLegalStandpointStandpointCollectionAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInputSystemCode is InputSystemCode Setter
// 系统标识
func (r *AlibabaLegalStandpointStandpointCollectionAPIRequest) SetInputSystemCode(_inputSystemCode string) error {
	r._inputSystemCode = _inputSystemCode
	r.Set("input_system_code", _inputSystemCode)
	return nil
}

// GetInputSystemCode InputSystemCode Getter
func (r AlibabaLegalStandpointStandpointCollectionAPIRequest) GetInputSystemCode() string {
	return r._inputSystemCode
}

// SetOperateWorkNo is OperateWorkNo Setter
// 操作人
func (r *AlibabaLegalStandpointStandpointCollectionAPIRequest) SetOperateWorkNo(_operateWorkNo string) error {
	r._operateWorkNo = _operateWorkNo
	r.Set("operate_work_no", _operateWorkNo)
	return nil
}

// GetOperateWorkNo OperateWorkNo Getter
func (r AlibabaLegalStandpointStandpointCollectionAPIRequest) GetOperateWorkNo() string {
	return r._operateWorkNo
}

// SetOperateName is OperateName Setter
// 操作人名称
func (r *AlibabaLegalStandpointStandpointCollectionAPIRequest) SetOperateName(_operateName string) error {
	r._operateName = _operateName
	r.Set("operate_name", _operateName)
	return nil
}

// GetOperateName OperateName Getter
func (r AlibabaLegalStandpointStandpointCollectionAPIRequest) GetOperateName() string {
	return r._operateName
}

// SetCollectionFlag is CollectionFlag Setter
// 收藏
func (r *AlibabaLegalStandpointStandpointCollectionAPIRequest) SetCollectionFlag(_collectionFlag string) error {
	r._collectionFlag = _collectionFlag
	r.Set("collection_flag", _collectionFlag)
	return nil
}

// GetCollectionFlag CollectionFlag Getter
func (r AlibabaLegalStandpointStandpointCollectionAPIRequest) GetCollectionFlag() string {
	return r._collectionFlag
}

// SetStandpointId is StandpointId Setter
// 口径id
func (r *AlibabaLegalStandpointStandpointCollectionAPIRequest) SetStandpointId(_standpointId int64) error {
	r._standpointId = _standpointId
	r.Set("standpoint_id", _standpointId)
	return nil
}

// GetStandpointId StandpointId Getter
func (r AlibabaLegalStandpointStandpointCollectionAPIRequest) GetStandpointId() int64 {
	return r._standpointId
}
