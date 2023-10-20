package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihouseimreceivemodelsyncAPIRequest IM承接方式同步 API请求
// alibaba.alihouse.im.receive.model.sync
//
// IM承接方式同步
type AlibabaalihouseimreceivemodelsyncAPIRequest struct {
	model.Params
	// IM承接方式DTO
	_imReceiveModelDto *ImreceiveModelReqDto
}

// NewAlibabaalihouseimreceivemodelsyncRequest 初始化AlibabaalihouseimreceivemodelsyncAPIRequest对象
func NewAlibabaalihouseimreceivemodelsyncRequest() *AlibabaalihouseimreceivemodelsyncAPIRequest {
	return &AlibabaalihouseimreceivemodelsyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihouseimreceivemodelsyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.im.receive.model.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihouseimreceivemodelsyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihouseimreceivemodelsyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetImReceiveModelDto is ImReceiveModelDto Setter
// IM承接方式DTO
func (r *AlibabaalihouseimreceivemodelsyncAPIRequest) SetImReceiveModelDto(_imReceiveModelDto *ImreceiveModelReqDto) error {
	r._imReceiveModelDto = _imReceiveModelDto
	r.Set("im_receive_model_dto", _imReceiveModelDto)
	return nil
}

// GetImReceiveModelDto ImReceiveModelDto Getter
func (r AlibabaalihouseimreceivemodelsyncAPIRequest) GetImReceiveModelDto() *ImreceiveModelReqDto {
	return r._imReceiveModelDto
}
