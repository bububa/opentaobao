package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCuntaoInteractRequisitionUpdateAPIRequest 更新物料制作状态 API请求
// alibaba.cuntao.interact.requisition.update
//
// 村淘物料下沉，更新物料制作状态
type AlibabaCuntaoInteractRequisitionUpdateAPIRequest struct {
	model.Params
	// 申请单id列表
	_uuidList []string
	// 物料制作状态
	_status string
}

// NewAlibabaCuntaoInteractRequisitionUpdateRequest 初始化AlibabaCuntaoInteractRequisitionUpdateAPIRequest对象
func NewAlibabaCuntaoInteractRequisitionUpdateRequest() *AlibabaCuntaoInteractRequisitionUpdateAPIRequest {
	return &AlibabaCuntaoInteractRequisitionUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCuntaoInteractRequisitionUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.cuntao.interact.requisition.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCuntaoInteractRequisitionUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetUuidList is UuidList Setter
// 申请单id列表
func (r *AlibabaCuntaoInteractRequisitionUpdateAPIRequest) SetUuidList(_uuidList []string) error {
	r._uuidList = _uuidList
	r.Set("uuid_list", _uuidList)
	return nil
}

// GetUuidList UuidList Getter
func (r AlibabaCuntaoInteractRequisitionUpdateAPIRequest) GetUuidList() []string {
	return r._uuidList
}

// SetStatus is Status Setter
// 物料制作状态
func (r *AlibabaCuntaoInteractRequisitionUpdateAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r AlibabaCuntaoInteractRequisitionUpdateAPIRequest) GetStatus() string {
	return r._status
}
