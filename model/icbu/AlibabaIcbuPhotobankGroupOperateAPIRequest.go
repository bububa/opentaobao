package icbu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaicbuphotobankgroupoperateAPIRequest 图片银行分组操作接口 API请求
// alibaba.icbu.photobank.group.operate
//
// 修改用户图片银行的分组信息，包括 新增分组，删除分组，分组重命名
type AlibabaicbuphotobankgroupoperateAPIRequest struct {
	model.Params
	// 图片分组操作请求对象
	_photoGroupOperationRequest *PhotoGroupOperationRequest
}

// NewAlibabaicbuphotobankgroupoperateRequest 初始化AlibabaicbuphotobankgroupoperateAPIRequest对象
func NewAlibabaicbuphotobankgroupoperateRequest() *AlibabaicbuphotobankgroupoperateAPIRequest {
	return &AlibabaicbuphotobankgroupoperateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaicbuphotobankgroupoperateAPIRequest) GetApiMethodName() string {
	return "alibaba.icbu.photobank.group.operate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaicbuphotobankgroupoperateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaicbuphotobankgroupoperateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPhotoGroupOperationRequest is PhotoGroupOperationRequest Setter
// 图片分组操作请求对象
func (r *AlibabaicbuphotobankgroupoperateAPIRequest) SetPhotoGroupOperationRequest(_photoGroupOperationRequest *PhotoGroupOperationRequest) error {
	r._photoGroupOperationRequest = _photoGroupOperationRequest
	r.Set("photo_group_operation_request", _photoGroupOperationRequest)
	return nil
}

// GetPhotoGroupOperationRequest PhotoGroupOperationRequest Getter
func (r AlibabaicbuphotobankgroupoperateAPIRequest) GetPhotoGroupOperationRequest() *PhotoGroupOperationRequest {
	return r._photoGroupOperationRequest
}
