package jym

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaJymIndustryVmosCommonCallbackAPIRequest vmos游戏信息采集结果回调通知 API请求
// alibaba.jym.industry.vmos.common.callback
//
// vmos游戏信息采集结果回调通知
type AlibabaJymIndustryVmosCommonCallbackAPIRequest struct {
	model.Params
	// 通用回调参数
	_param0 *CommonCallbackDto
}

// NewAlibabaJymIndustryVmosCommonCallbackRequest 初始化AlibabaJymIndustryVmosCommonCallbackAPIRequest对象
func NewAlibabaJymIndustryVmosCommonCallbackRequest() *AlibabaJymIndustryVmosCommonCallbackAPIRequest {
	return &AlibabaJymIndustryVmosCommonCallbackAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaJymIndustryVmosCommonCallbackAPIRequest) Reset() {
	r._param0 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaJymIndustryVmosCommonCallbackAPIRequest) GetApiMethodName() string {
	return "alibaba.jym.industry.vmos.common.callback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaJymIndustryVmosCommonCallbackAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaJymIndustryVmosCommonCallbackAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 通用回调参数
func (r *AlibabaJymIndustryVmosCommonCallbackAPIRequest) SetParam0(_param0 *CommonCallbackDto) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaJymIndustryVmosCommonCallbackAPIRequest) GetParam0() *CommonCallbackDto {
	return r._param0
}

var poolAlibabaJymIndustryVmosCommonCallbackAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaJymIndustryVmosCommonCallbackRequest()
	},
}

// GetAlibabaJymIndustryVmosCommonCallbackRequest 从 sync.Pool 获取 AlibabaJymIndustryVmosCommonCallbackAPIRequest
func GetAlibabaJymIndustryVmosCommonCallbackAPIRequest() *AlibabaJymIndustryVmosCommonCallbackAPIRequest {
	return poolAlibabaJymIndustryVmosCommonCallbackAPIRequest.Get().(*AlibabaJymIndustryVmosCommonCallbackAPIRequest)
}

// ReleaseAlibabaJymIndustryVmosCommonCallbackAPIRequest 将 AlibabaJymIndustryVmosCommonCallbackAPIRequest 放入 sync.Pool
func ReleaseAlibabaJymIndustryVmosCommonCallbackAPIRequest(v *AlibabaJymIndustryVmosCommonCallbackAPIRequest) {
	v.Reset()
	poolAlibabaJymIndustryVmosCommonCallbackAPIRequest.Put(v)
}
