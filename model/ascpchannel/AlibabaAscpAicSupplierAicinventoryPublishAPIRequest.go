package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascpaicsupplieraicinventorypublishAPIRequest 商家仓操作aic库存发布服务 API请求
// alibaba.ascp.aic.supplier.aicinventory.publish
//
// 商家调用这个接口来发布增加库存数据
type AlibabaascpaicsupplieraicinventorypublishAPIRequest struct {
	model.Params
	// 库存发布请求参数
	_aicInventoryPublishRequest *Aicinventorypublishrequest
}

// NewAlibabaascpaicsupplieraicinventorypublishRequest 初始化AlibabaascpaicsupplieraicinventorypublishAPIRequest对象
func NewAlibabaascpaicsupplieraicinventorypublishRequest() *AlibabaascpaicsupplieraicinventorypublishAPIRequest {
	return &AlibabaascpaicsupplieraicinventorypublishAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaascpaicsupplieraicinventorypublishAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.aic.supplier.aicinventory.publish"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaascpaicsupplieraicinventorypublishAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaascpaicsupplieraicinventorypublishAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAicInventoryPublishRequest is AicInventoryPublishRequest Setter
// 库存发布请求参数
func (r *AlibabaascpaicsupplieraicinventorypublishAPIRequest) SetAicInventoryPublishRequest(_aicInventoryPublishRequest *Aicinventorypublishrequest) error {
	r._aicInventoryPublishRequest = _aicInventoryPublishRequest
	r.Set("aic_inventory_publish_request", _aicInventoryPublishRequest)
	return nil
}

// GetAicInventoryPublishRequest AicInventoryPublishRequest Getter
func (r AlibabaascpaicsupplieraicinventorypublishAPIRequest) GetAicInventoryPublishRequest() *Aicinventorypublishrequest {
	return r._aicInventoryPublishRequest
}
