package alicom

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlicomWttOpentradeGetproductinfoAPIRequest 查询存送产品信息 API请求
// alibaba.alicom.wtt.opentrade.getproductinfo
//
// 话费宝查询产品信息相关配置
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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlicomWttOpentradeGetproductinfoAPIRequest) Reset() {
	r._productId = ""
	r._bizType = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlicomWttOpentradeGetproductinfoAPIRequest) GetApiMethodName() string {
	return "alibaba.alicom.wtt.opentrade.getproductinfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlicomWttOpentradeGetproductinfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlicomWttOpentradeGetproductinfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProductId is ProductId Setter
// 阿里通信产品ID
func (r *AlibabaAlicomWttOpentradeGetproductinfoAPIRequest) SetProductId(_productId string) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r AlibabaAlicomWttOpentradeGetproductinfoAPIRequest) GetProductId() string {
	return r._productId
}

// SetBizType is BizType Setter
// 类型
func (r *AlibabaAlicomWttOpentradeGetproductinfoAPIRequest) SetBizType(_bizType string) error {
	r._bizType = _bizType
	r.Set("biz_type", _bizType)
	return nil
}

// GetBizType BizType Getter
func (r AlibabaAlicomWttOpentradeGetproductinfoAPIRequest) GetBizType() string {
	return r._bizType
}

var poolAlibabaAlicomWttOpentradeGetproductinfoAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlicomWttOpentradeGetproductinfoRequest()
	},
}

// GetAlibabaAlicomWttOpentradeGetproductinfoRequest 从 sync.Pool 获取 AlibabaAlicomWttOpentradeGetproductinfoAPIRequest
func GetAlibabaAlicomWttOpentradeGetproductinfoAPIRequest() *AlibabaAlicomWttOpentradeGetproductinfoAPIRequest {
	return poolAlibabaAlicomWttOpentradeGetproductinfoAPIRequest.Get().(*AlibabaAlicomWttOpentradeGetproductinfoAPIRequest)
}

// ReleaseAlibabaAlicomWttOpentradeGetproductinfoAPIRequest 将 AlibabaAlicomWttOpentradeGetproductinfoAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlicomWttOpentradeGetproductinfoAPIRequest(v *AlibabaAlicomWttOpentradeGetproductinfoAPIRequest) {
	v.Reset()
	poolAlibabaAlicomWttOpentradeGetproductinfoAPIRequest.Put(v)
}
