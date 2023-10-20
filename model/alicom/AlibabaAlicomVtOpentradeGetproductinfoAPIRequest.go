package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalicomvtopentradegetproductinfoAPIRequest 查询新虚拟产品配置信息 API请求
// alibaba.alicom.vt.opentrade.getproductinfo
//
// 话费宝查询产品信息相关配置
type AlibabaalicomvtopentradegetproductinfoAPIRequest struct {
	model.Params
	// 阿里通信产品ID
	_productId string
	// 类型
	_bizType string
}

// NewAlibabaalicomvtopentradegetproductinfoRequest 初始化AlibabaalicomvtopentradegetproductinfoAPIRequest对象
func NewAlibabaalicomvtopentradegetproductinfoRequest() *AlibabaalicomvtopentradegetproductinfoAPIRequest {
	return &AlibabaalicomvtopentradegetproductinfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalicomvtopentradegetproductinfoAPIRequest) GetApiMethodName() string {
	return "alibaba.alicom.vt.opentrade.getproductinfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalicomvtopentradegetproductinfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalicomvtopentradegetproductinfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProductId is ProductId Setter
// 阿里通信产品ID
func (r *AlibabaalicomvtopentradegetproductinfoAPIRequest) SetProductId(_productId string) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r AlibabaalicomvtopentradegetproductinfoAPIRequest) GetProductId() string {
	return r._productId
}

// SetBizType is BizType Setter
// 类型
func (r *AlibabaalicomvtopentradegetproductinfoAPIRequest) SetBizType(_bizType string) error {
	r._bizType = _bizType
	r.Set("biz_type", _bizType)
	return nil
}

// GetBizType BizType Getter
func (r AlibabaalicomvtopentradegetproductinfoAPIRequest) GetBizType() string {
	return r._bizType
}
