package moscm

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMosGoodsSetpriceAPIRequest 价格变更接口 API请求
// alibaba.mos.goods.setprice
//
// 价格变更接口，供供应商修改价格时使用。
type AlibabaMosGoodsSetpriceAPIRequest struct {
	model.Params
	// 价格变更对象列表
	_priceDtoList []PriceDto
}

// NewAlibabaMosGoodsSetpriceRequest 初始化AlibabaMosGoodsSetpriceAPIRequest对象
func NewAlibabaMosGoodsSetpriceRequest() *AlibabaMosGoodsSetpriceAPIRequest {
	return &AlibabaMosGoodsSetpriceAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaMosGoodsSetpriceAPIRequest) Reset() {
	r._priceDtoList = r._priceDtoList[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMosGoodsSetpriceAPIRequest) GetApiMethodName() string {
	return "alibaba.mos.goods.setprice"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMosGoodsSetpriceAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMosGoodsSetpriceAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPriceDtoList is PriceDtoList Setter
// 价格变更对象列表
func (r *AlibabaMosGoodsSetpriceAPIRequest) SetPriceDtoList(_priceDtoList []PriceDto) error {
	r._priceDtoList = _priceDtoList
	r.Set("price_dto_list", _priceDtoList)
	return nil
}

// GetPriceDtoList PriceDtoList Getter
func (r AlibabaMosGoodsSetpriceAPIRequest) GetPriceDtoList() []PriceDto {
	return r._priceDtoList
}

var poolAlibabaMosGoodsSetpriceAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaMosGoodsSetpriceRequest()
	},
}

// GetAlibabaMosGoodsSetpriceRequest 从 sync.Pool 获取 AlibabaMosGoodsSetpriceAPIRequest
func GetAlibabaMosGoodsSetpriceAPIRequest() *AlibabaMosGoodsSetpriceAPIRequest {
	return poolAlibabaMosGoodsSetpriceAPIRequest.Get().(*AlibabaMosGoodsSetpriceAPIRequest)
}

// ReleaseAlibabaMosGoodsSetpriceAPIRequest 将 AlibabaMosGoodsSetpriceAPIRequest 放入 sync.Pool
func ReleaseAlibabaMosGoodsSetpriceAPIRequest(v *AlibabaMosGoodsSetpriceAPIRequest) {
	v.Reset()
	poolAlibabaMosGoodsSetpriceAPIRequest.Put(v)
}
