package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascpindustryuopsupplierconsignoderAPIRequest 商家推单 API请求
// alibaba.ascp.industry.uop.supplier.consignoder
//
// 商家推单
type AlibabaascpindustryuopsupplierconsignoderAPIRequest struct {
	model.Params
	// 发货主单信息
	_erpNormalConsignOrderRequest *Erpnormalconsignorderrequest
}

// NewAlibabaascpindustryuopsupplierconsignoderRequest 初始化AlibabaascpindustryuopsupplierconsignoderAPIRequest对象
func NewAlibabaascpindustryuopsupplierconsignoderRequest() *AlibabaascpindustryuopsupplierconsignoderAPIRequest {
	return &AlibabaascpindustryuopsupplierconsignoderAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaascpindustryuopsupplierconsignoderAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.industry.uop.supplier.consignoder"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaascpindustryuopsupplierconsignoderAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaascpindustryuopsupplierconsignoderAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetErpNormalConsignOrderRequest is ErpNormalConsignOrderRequest Setter
// 发货主单信息
func (r *AlibabaascpindustryuopsupplierconsignoderAPIRequest) SetErpNormalConsignOrderRequest(_erpNormalConsignOrderRequest *Erpnormalconsignorderrequest) error {
	r._erpNormalConsignOrderRequest = _erpNormalConsignOrderRequest
	r.Set("erp_normal_consign_order_request", _erpNormalConsignOrderRequest)
	return nil
}

// GetErpNormalConsignOrderRequest ErpNormalConsignOrderRequest Getter
func (r AlibabaascpindustryuopsupplierconsignoderAPIRequest) GetErpNormalConsignOrderRequest() *Erpnormalconsignorderrequest {
	return r._erpNormalConsignOrderRequest
}
