package alihouse

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseImReceiveModelSyncAPIRequest) Reset() {
	r._imReceiveModelDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseImReceiveModelSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.im.receive.model.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseImReceiveModelSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseImReceiveModelSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaAlihouseImReceiveModelSyncAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseImReceiveModelSyncRequest()
	},
}

// GetAlibabaAlihouseImReceiveModelSyncRequest 从 sync.Pool 获取 AlibabaAlihouseImReceiveModelSyncAPIRequest
func GetAlibabaAlihouseImReceiveModelSyncAPIRequest() *AlibabaAlihouseImReceiveModelSyncAPIRequest {
	return poolAlibabaAlihouseImReceiveModelSyncAPIRequest.Get().(*AlibabaAlihouseImReceiveModelSyncAPIRequest)
}

// ReleaseAlibabaAlihouseImReceiveModelSyncAPIRequest 将 AlibabaAlihouseImReceiveModelSyncAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseImReceiveModelSyncAPIRequest(v *AlibabaAlihouseImReceiveModelSyncAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseImReceiveModelSyncAPIRequest.Put(v)
}
