package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkUmsInventoryPublishAPIRequest 初始化覆盖实物库存 API请求
// alibaba.wdk.ums.inventory.publish
//
// 先去库存这边查询当前实物库存有多少的量，然后算出来需要增加的量。接下来调用ums原来的入库语义的接口进行库存的增量补充
type AlibabaWdkUmsInventoryPublishAPIRequest struct {
	model.Params
	// 1
	_wdkErpArrivalNotice *WdkErpArrivalNoticeDto
}

// NewAlibabaWdkUmsInventoryPublishRequest 初始化AlibabaWdkUmsInventoryPublishAPIRequest对象
func NewAlibabaWdkUmsInventoryPublishRequest() *AlibabaWdkUmsInventoryPublishAPIRequest {
	return &AlibabaWdkUmsInventoryPublishAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkUmsInventoryPublishAPIRequest) Reset() {
	r._wdkErpArrivalNotice = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkUmsInventoryPublishAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.ums.inventory.publish"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkUmsInventoryPublishAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkUmsInventoryPublishAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWdkErpArrivalNotice is WdkErpArrivalNotice Setter
// 1
func (r *AlibabaWdkUmsInventoryPublishAPIRequest) SetWdkErpArrivalNotice(_wdkErpArrivalNotice *WdkErpArrivalNoticeDto) error {
	r._wdkErpArrivalNotice = _wdkErpArrivalNotice
	r.Set("wdk_erp_arrival_notice", _wdkErpArrivalNotice)
	return nil
}

// GetWdkErpArrivalNotice WdkErpArrivalNotice Getter
func (r AlibabaWdkUmsInventoryPublishAPIRequest) GetWdkErpArrivalNotice() *WdkErpArrivalNoticeDto {
	return r._wdkErpArrivalNotice
}

var poolAlibabaWdkUmsInventoryPublishAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkUmsInventoryPublishRequest()
	},
}

// GetAlibabaWdkUmsInventoryPublishRequest 从 sync.Pool 获取 AlibabaWdkUmsInventoryPublishAPIRequest
func GetAlibabaWdkUmsInventoryPublishAPIRequest() *AlibabaWdkUmsInventoryPublishAPIRequest {
	return poolAlibabaWdkUmsInventoryPublishAPIRequest.Get().(*AlibabaWdkUmsInventoryPublishAPIRequest)
}

// ReleaseAlibabaWdkUmsInventoryPublishAPIRequest 将 AlibabaWdkUmsInventoryPublishAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkUmsInventoryPublishAPIRequest(v *AlibabaWdkUmsInventoryPublishAPIRequest) {
	v.Reset()
	poolAlibabaWdkUmsInventoryPublishAPIRequest.Put(v)
}
