package icbu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIcbuPhotobankGroupOperateAPIRequest 图片银行分组操作接口 API请求
// alibaba.icbu.photobank.group.operate
//
// 修改用户图片银行的分组信息，包括 新增分组，删除分组，分组重命名
type AlibabaIcbuPhotobankGroupOperateAPIRequest struct {
	model.Params
	// 图片分组操作请求对象
	_photoGroupOperationRequest *PhotoGroupOperationRequest
}

// NewAlibabaIcbuPhotobankGroupOperateRequest 初始化AlibabaIcbuPhotobankGroupOperateAPIRequest对象
func NewAlibabaIcbuPhotobankGroupOperateRequest() *AlibabaIcbuPhotobankGroupOperateAPIRequest {
	return &AlibabaIcbuPhotobankGroupOperateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIcbuPhotobankGroupOperateAPIRequest) GetApiMethodName() string {
	return "alibaba.icbu.photobank.group.operate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIcbuPhotobankGroupOperateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIcbuPhotobankGroupOperateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPhotoGroupOperationRequest is PhotoGroupOperationRequest Setter
// 图片分组操作请求对象
func (r *AlibabaIcbuPhotobankGroupOperateAPIRequest) SetPhotoGroupOperationRequest(_photoGroupOperationRequest *PhotoGroupOperationRequest) error {
	r._photoGroupOperationRequest = _photoGroupOperationRequest
	r.Set("photo_group_operation_request", _photoGroupOperationRequest)
	return nil
}

// GetPhotoGroupOperationRequest PhotoGroupOperationRequest Getter
func (r AlibabaIcbuPhotobankGroupOperateAPIRequest) GetPhotoGroupOperationRequest() *PhotoGroupOperationRequest {
	return r._photoGroupOperationRequest
}
