package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkUmsInventoryPublishAPIRequest
初始化覆盖实物库存 API请求
alibaba.wdk.ums.inventory.publish

先去库存这边查询当前实物库存有多少的量，然后算出来需要增加的量。接下来调用ums原来的入库语义的接口进行库存的增量补充 */
type AlibabaWdkUmsInventoryPublishAPIRequest struct {
	model.Params
	// 1
	_wdkErpArrivalNotice *WdkErpArrivalNoticeDto
}

// NewAlibabaWdkUmsInventoryPublishRequest 初始化AlibabaWdkUmsInventoryPublishAPIRequest对象
func NewAlibabaWdkUmsInventoryPublishRequest() *AlibabaWdkUmsInventoryPublishAPIRequest {
	return &AlibabaWdkUmsInventoryPublishAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkUmsInventoryPublishAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.ums.inventory.publish"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkUmsInventoryPublishAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is WdkErpArrivalNotice Setter
// 1
func (r *AlibabaWdkUmsInventoryPublishAPIRequest) SetWdkErpArrivalNotice(_wdkErpArrivalNotice *WdkErpArrivalNoticeDto) error {
	r._wdkErpArrivalNotice = _wdkErpArrivalNotice
	r.Set("wdk_erp_arrival_notice", _wdkErpArrivalNotice)
	return nil
}

// Get WdkErpArrivalNotice Getter
func (r AlibabaWdkUmsInventoryPublishAPIRequest) GetWdkErpArrivalNotice() *WdkErpArrivalNoticeDto {
	return r._wdkErpArrivalNotice
}
