package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlicomWttOpentradeGetproductinfoAPIRequest
查询存送产品信息 API请求
alibaba.alicom.wtt.opentrade.getproductinfo

话费宝查询产品信息相关配置 */
type AlibabaAlicomWttOpentradeGetproductinfoAPIRequest struct {
	model.Params
	// 阿里通信产品ID
	_productId string
	// 类型
	_bizType string
}

// NewAlibabaAlicomWttOpentradeGetproductinfoRequest 初始化AlibabaAlicomWttOpentradeGetproductinfoAPIRequest对象
func NewAlibabaAlicomWttOpentradeGetproductinfoRequest() *AlibabaAlicomWttOpentradeGetproductinfoAPIRequest {
	return &AlibabaAlicomWttOpentradeGetproductinfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlicomWttOpentradeGetproductinfoAPIRequest) GetApiMethodName() string {
	return "alibaba.alicom.wtt.opentrade.getproductinfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlicomWttOpentradeGetproductinfoAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ProductId Setter
// 阿里通信产品ID
func (r *AlibabaAlicomWttOpentradeGetproductinfoAPIRequest) SetProductId(_productId string) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// Get ProductId Getter
func (r AlibabaAlicomWttOpentradeGetproductinfoAPIRequest) GetProductId() string {
	return r._productId
}

// Set is BizType Setter
// 类型
func (r *AlibabaAlicomWttOpentradeGetproductinfoAPIRequest) SetBizType(_bizType string) error {
	r._bizType = _bizType
	r.Set("biz_type", _bizType)
	return nil
}

// Get BizType Getter
func (r AlibabaAlicomWttOpentradeGetproductinfoAPIRequest) GetBizType() string {
	return r._bizType
}
