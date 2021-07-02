package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWholesaleGoodsGetAPIRequest 查询阿里巴巴批发市场商品详情 API请求
// alibaba.wholesale.goods.get
//
// 查询阿里巴巴批发市场商品详情
type AlibabaWholesaleGoodsGetAPIRequest struct {
	model.Params
	// country_code
	_countryCode string
	// id
	_id string
}

// NewAlibabaWholesaleGoodsGetRequest 初始化AlibabaWholesaleGoodsGetAPIRequest对象
func NewAlibabaWholesaleGoodsGetRequest() *AlibabaWholesaleGoodsGetAPIRequest {
	return &AlibabaWholesaleGoodsGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWholesaleGoodsGetAPIRequest) GetApiMethodName() string {
	return "alibaba.wholesale.goods.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWholesaleGoodsGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetCountryCode is CountryCode Setter
// country_code
func (r *AlibabaWholesaleGoodsGetAPIRequest) SetCountryCode(_countryCode string) error {
	r._countryCode = _countryCode
	r.Set("country_code", _countryCode)
	return nil
}

// GetCountryCode CountryCode Getter
func (r AlibabaWholesaleGoodsGetAPIRequest) GetCountryCode() string {
	return r._countryCode
}

// SetId is Id Setter
// id
func (r *AlibabaWholesaleGoodsGetAPIRequest) SetId(_id string) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r AlibabaWholesaleGoodsGetAPIRequest) GetId() string {
	return r._id
}
