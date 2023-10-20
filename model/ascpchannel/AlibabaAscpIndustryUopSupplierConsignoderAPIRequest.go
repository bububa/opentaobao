package ascpchannel

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpIndustryUopSupplierConsignoderAPIRequest 商家推单 API请求
// alibaba.ascp.industry.uop.supplier.consignoder
//
// 商家推单
type AlibabaAscpIndustryUopSupplierConsignoderAPIRequest struct {
	model.Params
	// 发货主单信息
	_erpNormalConsignOrderRequest *Erpnormalconsignorderrequest
}

// NewAlibabaAscpIndustryUopSupplierConsignoderRequest 初始化AlibabaAscpIndustryUopSupplierConsignoderAPIRequest对象
func NewAlibabaAscpIndustryUopSupplierConsignoderRequest() *AlibabaAscpIndustryUopSupplierConsignoderAPIRequest {
	return &AlibabaAscpIndustryUopSupplierConsignoderAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAscpIndustryUopSupplierConsignoderAPIRequest) Reset() {
	r._erpNormalConsignOrderRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpIndustryUopSupplierConsignoderAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.industry.uop.supplier.consignoder"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpIndustryUopSupplierConsignoderAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAscpIndustryUopSupplierConsignoderAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetErpNormalConsignOrderRequest is ErpNormalConsignOrderRequest Setter
// 发货主单信息
func (r *AlibabaAscpIndustryUopSupplierConsignoderAPIRequest) SetErpNormalConsignOrderRequest(_erpNormalConsignOrderRequest *Erpnormalconsignorderrequest) error {
	r._erpNormalConsignOrderRequest = _erpNormalConsignOrderRequest
	r.Set("erp_normal_consign_order_request", _erpNormalConsignOrderRequest)
	return nil
}

// GetErpNormalConsignOrderRequest ErpNormalConsignOrderRequest Getter
func (r AlibabaAscpIndustryUopSupplierConsignoderAPIRequest) GetErpNormalConsignOrderRequest() *Erpnormalconsignorderrequest {
	return r._erpNormalConsignOrderRequest
}

var poolAlibabaAscpIndustryUopSupplierConsignoderAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAscpIndustryUopSupplierConsignoderRequest()
	},
}

// GetAlibabaAscpIndustryUopSupplierConsignoderRequest 从 sync.Pool 获取 AlibabaAscpIndustryUopSupplierConsignoderAPIRequest
func GetAlibabaAscpIndustryUopSupplierConsignoderAPIRequest() *AlibabaAscpIndustryUopSupplierConsignoderAPIRequest {
	return poolAlibabaAscpIndustryUopSupplierConsignoderAPIRequest.Get().(*AlibabaAscpIndustryUopSupplierConsignoderAPIRequest)
}

// ReleaseAlibabaAscpIndustryUopSupplierConsignoderAPIRequest 将 AlibabaAscpIndustryUopSupplierConsignoderAPIRequest 放入 sync.Pool
func ReleaseAlibabaAscpIndustryUopSupplierConsignoderAPIRequest(v *AlibabaAscpIndustryUopSupplierConsignoderAPIRequest) {
	v.Reset()
	poolAlibabaAscpIndustryUopSupplierConsignoderAPIRequest.Put(v)
}
