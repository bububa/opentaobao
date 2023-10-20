package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalicomwttopentradegetproductinfoAPIRequest 查询存送产品信息 API请求
// alibaba.alicom.wtt.opentrade.getproductinfo
//
// 话费宝查询产品信息相关配置
type AlibabaalicomwttopentradegetproductinfoAPIRequest struct {
	model.Params
	// 阿里通信产品ID
	_productId string
	// 类型
	_bizType string
}

// NewAlibabaalicomwttopentradegetproductinfoRequest 初始化AlibabaalicomwttopentradegetproductinfoAPIRequest对象
func NewAlibabaalicomwttopentradegetproductinfoRequest() *AlibabaalicomwttopentradegetproductinfoAPIRequest {
	return &AlibabaalicomwttopentradegetproductinfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalicomwttopentradegetproductinfoAPIRequest) GetApiMethodName() string {
	return "alibaba.alicom.wtt.opentrade.getproductinfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalicomwttopentradegetproductinfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalicomwttopentradegetproductinfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProductId is ProductId Setter
// 阿里通信产品ID
func (r *AlibabaalicomwttopentradegetproductinfoAPIRequest) SetProductId(_productId string) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r AlibabaalicomwttopentradegetproductinfoAPIRequest) GetProductId() string {
	return r._productId
}

// SetBizType is BizType Setter
// 类型
func (r *AlibabaalicomwttopentradegetproductinfoAPIRequest) SetBizType(_bizType string) error {
	r._bizType = _bizType
	r.Set("biz_type", _bizType)
	return nil
}

// GetBizType BizType Getter
func (r AlibabaalicomwttopentradegetproductinfoAPIRequest) GetBizType() string {
	return r._bizType
}
