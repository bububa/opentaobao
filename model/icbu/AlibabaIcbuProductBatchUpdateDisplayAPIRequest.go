package icbu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIcbuProductBatchUpdateDisplayAPIRequest
商品批量上下架接口 API请求
alibaba.icbu.product.batch.update.display

给国际站的三方服务商提供批量上下架接口 */
type AlibabaIcbuProductBatchUpdateDisplayAPIRequest struct {
	model.Params
	// on表示上架，off表示下架
	_newDisplay string
	// 用逗号分隔的混淆id字符串
	_productIdList string
}

// NewAlibabaIcbuProductBatchUpdateDisplayRequest 初始化AlibabaIcbuProductBatchUpdateDisplayAPIRequest对象
func NewAlibabaIcbuProductBatchUpdateDisplayRequest() *AlibabaIcbuProductBatchUpdateDisplayAPIRequest {
	return &AlibabaIcbuProductBatchUpdateDisplayAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIcbuProductBatchUpdateDisplayAPIRequest) GetApiMethodName() string {
	return "alibaba.icbu.product.batch.update.display"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIcbuProductBatchUpdateDisplayAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is NewDisplay Setter
// on表示上架，off表示下架
func (r *AlibabaIcbuProductBatchUpdateDisplayAPIRequest) SetNewDisplay(_newDisplay string) error {
	r._newDisplay = _newDisplay
	r.Set("new_display", _newDisplay)
	return nil
}

// Get NewDisplay Getter
func (r AlibabaIcbuProductBatchUpdateDisplayAPIRequest) GetNewDisplay() string {
	return r._newDisplay
}

// Set is ProductIdList Setter
// 用逗号分隔的混淆id字符串
func (r *AlibabaIcbuProductBatchUpdateDisplayAPIRequest) SetProductIdList(_productIdList string) error {
	r._productIdList = _productIdList
	r.Set("product_id_list", _productIdList)
	return nil
}

// Get ProductIdList Getter
func (r AlibabaIcbuProductBatchUpdateDisplayAPIRequest) GetProductIdList() string {
	return r._productIdList
}
