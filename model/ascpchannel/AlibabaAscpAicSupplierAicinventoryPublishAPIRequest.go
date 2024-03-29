package ascpchannel

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpAicSupplierAicinventoryPublishAPIRequest 商家仓操作aic库存发布服务 API请求
// alibaba.ascp.aic.supplier.aicinventory.publish
//
// 商家调用这个接口来发布增加库存数据
type AlibabaAscpAicSupplierAicinventoryPublishAPIRequest struct {
	model.Params
	// 库存发布请求参数
	_aicInventoryPublishRequest *Aicinventorypublishrequest
}

// NewAlibabaAscpAicSupplierAicinventoryPublishRequest 初始化AlibabaAscpAicSupplierAicinventoryPublishAPIRequest对象
func NewAlibabaAscpAicSupplierAicinventoryPublishRequest() *AlibabaAscpAicSupplierAicinventoryPublishAPIRequest {
	return &AlibabaAscpAicSupplierAicinventoryPublishAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAscpAicSupplierAicinventoryPublishAPIRequest) Reset() {
	r._aicInventoryPublishRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpAicSupplierAicinventoryPublishAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.aic.supplier.aicinventory.publish"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpAicSupplierAicinventoryPublishAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAscpAicSupplierAicinventoryPublishAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAicInventoryPublishRequest is AicInventoryPublishRequest Setter
// 库存发布请求参数
func (r *AlibabaAscpAicSupplierAicinventoryPublishAPIRequest) SetAicInventoryPublishRequest(_aicInventoryPublishRequest *Aicinventorypublishrequest) error {
	r._aicInventoryPublishRequest = _aicInventoryPublishRequest
	r.Set("aic_inventory_publish_request", _aicInventoryPublishRequest)
	return nil
}

// GetAicInventoryPublishRequest AicInventoryPublishRequest Getter
func (r AlibabaAscpAicSupplierAicinventoryPublishAPIRequest) GetAicInventoryPublishRequest() *Aicinventorypublishrequest {
	return r._aicInventoryPublishRequest
}

var poolAlibabaAscpAicSupplierAicinventoryPublishAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAscpAicSupplierAicinventoryPublishRequest()
	},
}

// GetAlibabaAscpAicSupplierAicinventoryPublishRequest 从 sync.Pool 获取 AlibabaAscpAicSupplierAicinventoryPublishAPIRequest
func GetAlibabaAscpAicSupplierAicinventoryPublishAPIRequest() *AlibabaAscpAicSupplierAicinventoryPublishAPIRequest {
	return poolAlibabaAscpAicSupplierAicinventoryPublishAPIRequest.Get().(*AlibabaAscpAicSupplierAicinventoryPublishAPIRequest)
}

// ReleaseAlibabaAscpAicSupplierAicinventoryPublishAPIRequest 将 AlibabaAscpAicSupplierAicinventoryPublishAPIRequest 放入 sync.Pool
func ReleaseAlibabaAscpAicSupplierAicinventoryPublishAPIRequest(v *AlibabaAscpAicSupplierAicinventoryPublishAPIRequest) {
	v.Reset()
	poolAlibabaAscpAicSupplierAicinventoryPublishAPIRequest.Put(v)
}
