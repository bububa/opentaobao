package legalsuit

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalegalstandpointstandpointcollectionAPIRequest 收藏|取消收藏 API请求
// alibaba.legal.standpoint.standpoint.collection
//
// 收藏|取消收藏
type AlibabalegalstandpointstandpointcollectionAPIRequest struct {
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

// NewAlibabalegalstandpointstandpointcollectionRequest 初始化AlibabalegalstandpointstandpointcollectionAPIRequest对象
func NewAlibabalegalstandpointstandpointcollectionRequest() *AlibabalegalstandpointstandpointcollectionAPIRequest {
	return &AlibabalegalstandpointstandpointcollectionAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalegalstandpointstandpointcollectionAPIRequest) GetApiMethodName() string {
	return "alibaba.legal.standpoint.standpoint.collection"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalegalstandpointstandpointcollectionAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalegalstandpointstandpointcollectionAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInputSystemCode is InputSystemCode Setter
// 系统标识
func (r *AlibabalegalstandpointstandpointcollectionAPIRequest) SetInputSystemCode(_inputSystemCode string) error {
	r._inputSystemCode = _inputSystemCode
	r.Set("input_system_code", _inputSystemCode)
	return nil
}

// GetInputSystemCode InputSystemCode Getter
func (r AlibabalegalstandpointstandpointcollectionAPIRequest) GetInputSystemCode() string {
	return r._inputSystemCode
}

// SetOperateWorkNo is OperateWorkNo Setter
// 操作人
func (r *AlibabalegalstandpointstandpointcollectionAPIRequest) SetOperateWorkNo(_operateWorkNo string) error {
	r._operateWorkNo = _operateWorkNo
	r.Set("operate_work_no", _operateWorkNo)
	return nil
}

// GetOperateWorkNo OperateWorkNo Getter
func (r AlibabalegalstandpointstandpointcollectionAPIRequest) GetOperateWorkNo() string {
	return r._operateWorkNo
}

// SetOperateName is OperateName Setter
// 操作人名称
func (r *AlibabalegalstandpointstandpointcollectionAPIRequest) SetOperateName(_operateName string) error {
	r._operateName = _operateName
	r.Set("operate_name", _operateName)
	return nil
}

// GetOperateName OperateName Getter
func (r AlibabalegalstandpointstandpointcollectionAPIRequest) GetOperateName() string {
	return r._operateName
}

// SetCollectionFlag is CollectionFlag Setter
// 收藏
func (r *AlibabalegalstandpointstandpointcollectionAPIRequest) SetCollectionFlag(_collectionFlag string) error {
	r._collectionFlag = _collectionFlag
	r.Set("collection_flag", _collectionFlag)
	return nil
}

// GetCollectionFlag CollectionFlag Getter
func (r AlibabalegalstandpointstandpointcollectionAPIRequest) GetCollectionFlag() string {
	return r._collectionFlag
}

// SetStandpointId is StandpointId Setter
// 口径id
func (r *AlibabalegalstandpointstandpointcollectionAPIRequest) SetStandpointId(_standpointId int64) error {
	r._standpointId = _standpointId
	r.Set("standpoint_id", _standpointId)
	return nil
}

// GetStandpointId StandpointId Getter
func (r AlibabalegalstandpointstandpointcollectionAPIRequest) GetStandpointId() int64 {
	return r._standpointId
}
