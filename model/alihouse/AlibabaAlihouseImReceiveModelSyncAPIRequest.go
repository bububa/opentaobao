package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseImReceiveModelSyncAPIRequest IM承接方式同步 API请求
// alibaba.alihouse.im.receive.model.sync
//
// IM承接方式同步
type AlibabaAlihouseImReceiveModelSyncAPIRequest struct {
	model.Params
	// IM承接方式DTO
	_imReceiveModelDto *IMReceiveModelReqDto
}

// NewAlibabaAlihouseImReceiveModelSyncRequest 初始化AlibabaAlihouseImReceiveModelSyncAPIRequest对象
func NewAlibabaAlihouseImReceiveModelSyncRequest() *AlibabaAlihouseImReceiveModelSyncAPIRequest {
	return &AlibabaAlihouseImReceiveModelSyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseImReceiveModelSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.im.receive.model.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseImReceiveModelSyncAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetImReceiveModelDto is ImReceiveModelDto Setter
// IM承接方式DTO
func (r *AlibabaAlihouseImReceiveModelSyncAPIRequest) SetImReceiveModelDto(_imReceiveModelDto *IMReceiveModelReqDto) error {
	r._imReceiveModelDto = _imReceiveModelDto
	r.Set("im_receive_model_dto", _imReceiveModelDto)
	return nil
}

// GetImReceiveModelDto ImReceiveModelDto Getter
func (r AlibabaAlihouseImReceiveModelSyncAPIRequest) GetImReceiveModelDto() *IMReceiveModelReqDto {
	return r._imReceiveModelDto
}
