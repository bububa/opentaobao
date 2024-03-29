package alihealth2

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHealthNrCepOutorderUploadAPIRequest 线上订单收货验收单、出入库单据生成接口 API请求
// alibaba.health.nr.cep.outorder.upload
//
// 线上订单收货验收单、出入库单据生成接口
type AlibabaHealthNrCepOutorderUploadAPIRequest struct {
	model.Params
	// 出库单对象
	_topWarOutDto *TopWarOutDto
}

// NewAlibabaHealthNrCepOutorderUploadRequest 初始化AlibabaHealthNrCepOutorderUploadAPIRequest对象
func NewAlibabaHealthNrCepOutorderUploadRequest() *AlibabaHealthNrCepOutorderUploadAPIRequest {
	return &AlibabaHealthNrCepOutorderUploadAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaHealthNrCepOutorderUploadAPIRequest) Reset() {
	r._topWarOutDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHealthNrCepOutorderUploadAPIRequest) GetApiMethodName() string {
	return "alibaba.health.nr.cep.outorder.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHealthNrCepOutorderUploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaHealthNrCepOutorderUploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopWarOutDto is TopWarOutDto Setter
// 出库单对象
func (r *AlibabaHealthNrCepOutorderUploadAPIRequest) SetTopWarOutDto(_topWarOutDto *TopWarOutDto) error {
	r._topWarOutDto = _topWarOutDto
	r.Set("top_war_out_dto", _topWarOutDto)
	return nil
}

// GetTopWarOutDto TopWarOutDto Getter
func (r AlibabaHealthNrCepOutorderUploadAPIRequest) GetTopWarOutDto() *TopWarOutDto {
	return r._topWarOutDto
}

var poolAlibabaHealthNrCepOutorderUploadAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaHealthNrCepOutorderUploadRequest()
	},
}

// GetAlibabaHealthNrCepOutorderUploadRequest 从 sync.Pool 获取 AlibabaHealthNrCepOutorderUploadAPIRequest
func GetAlibabaHealthNrCepOutorderUploadAPIRequest() *AlibabaHealthNrCepOutorderUploadAPIRequest {
	return poolAlibabaHealthNrCepOutorderUploadAPIRequest.Get().(*AlibabaHealthNrCepOutorderUploadAPIRequest)
}

// ReleaseAlibabaHealthNrCepOutorderUploadAPIRequest 将 AlibabaHealthNrCepOutorderUploadAPIRequest 放入 sync.Pool
func ReleaseAlibabaHealthNrCepOutorderUploadAPIRequest(v *AlibabaHealthNrCepOutorderUploadAPIRequest) {
	v.Reset()
	poolAlibabaHealthNrCepOutorderUploadAPIRequest.Put(v)
}
