package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkumsinventorypublishAPIRequest 初始化覆盖实物库存 API请求
// alibaba.wdk.ums.inventory.publish
//
// 先去库存这边查询当前实物库存有多少的量，然后算出来需要增加的量。接下来调用ums原来的入库语义的接口进行库存的增量补充
type AlibabawdkumsinventorypublishAPIRequest struct {
	model.Params
	// 1
	_wdkErpArrivalNotice *WdkErpArrivalNoticeDto
}

// NewAlibabawdkumsinventorypublishRequest 初始化AlibabawdkumsinventorypublishAPIRequest对象
func NewAlibabawdkumsinventorypublishRequest() *AlibabawdkumsinventorypublishAPIRequest {
	return &AlibabawdkumsinventorypublishAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkumsinventorypublishAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.ums.inventory.publish"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkumsinventorypublishAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkumsinventorypublishAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWdkErpArrivalNotice is WdkErpArrivalNotice Setter
// 1
func (r *AlibabawdkumsinventorypublishAPIRequest) SetWdkErpArrivalNotice(_wdkErpArrivalNotice *WdkErpArrivalNoticeDto) error {
	r._wdkErpArrivalNotice = _wdkErpArrivalNotice
	r.Set("wdk_erp_arrival_notice", _wdkErpArrivalNotice)
	return nil
}

// GetWdkErpArrivalNotice WdkErpArrivalNotice Getter
func (r AlibabawdkumsinventorypublishAPIRequest) GetWdkErpArrivalNotice() *WdkErpArrivalNoticeDto {
	return r._wdkErpArrivalNotice
}
