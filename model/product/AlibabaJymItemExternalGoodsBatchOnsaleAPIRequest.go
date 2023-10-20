package product

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaJymItemExternalGoodsBatchOnsaleAPIRequest 交易猫外部商家批量上架商品接口 API请求
// alibaba.jym.item.external.goods.batch.onsale
//
// 供外部B端商家接入，提交批量上架商品请求，返回批量上架任务结果
type AlibabaJymItemExternalGoodsBatchOnsaleAPIRequest struct {
	model.Params
	// 商品批量上架接口入参
	_goodsOnSaleCommand *GoodsOnSaleCommandDto
}

// NewAlibabaJymItemExternalGoodsBatchOnsaleRequest 初始化AlibabaJymItemExternalGoodsBatchOnsaleAPIRequest对象
func NewAlibabaJymItemExternalGoodsBatchOnsaleRequest() *AlibabaJymItemExternalGoodsBatchOnsaleAPIRequest {
	return &AlibabaJymItemExternalGoodsBatchOnsaleAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaJymItemExternalGoodsBatchOnsaleAPIRequest) Reset() {
	r._goodsOnSaleCommand = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaJymItemExternalGoodsBatchOnsaleAPIRequest) GetApiMethodName() string {
	return "alibaba.jym.item.external.goods.batch.onsale"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaJymItemExternalGoodsBatchOnsaleAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaJymItemExternalGoodsBatchOnsaleAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetGoodsOnSaleCommand is GoodsOnSaleCommand Setter
// 商品批量上架接口入参
func (r *AlibabaJymItemExternalGoodsBatchOnsaleAPIRequest) SetGoodsOnSaleCommand(_goodsOnSaleCommand *GoodsOnSaleCommandDto) error {
	r._goodsOnSaleCommand = _goodsOnSaleCommand
	r.Set("goods_on_sale_command", _goodsOnSaleCommand)
	return nil
}

// GetGoodsOnSaleCommand GoodsOnSaleCommand Getter
func (r AlibabaJymItemExternalGoodsBatchOnsaleAPIRequest) GetGoodsOnSaleCommand() *GoodsOnSaleCommandDto {
	return r._goodsOnSaleCommand
}

var poolAlibabaJymItemExternalGoodsBatchOnsaleAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaJymItemExternalGoodsBatchOnsaleRequest()
	},
}

// GetAlibabaJymItemExternalGoodsBatchOnsaleRequest 从 sync.Pool 获取 AlibabaJymItemExternalGoodsBatchOnsaleAPIRequest
func GetAlibabaJymItemExternalGoodsBatchOnsaleAPIRequest() *AlibabaJymItemExternalGoodsBatchOnsaleAPIRequest {
	return poolAlibabaJymItemExternalGoodsBatchOnsaleAPIRequest.Get().(*AlibabaJymItemExternalGoodsBatchOnsaleAPIRequest)
}

// ReleaseAlibabaJymItemExternalGoodsBatchOnsaleAPIRequest 将 AlibabaJymItemExternalGoodsBatchOnsaleAPIRequest 放入 sync.Pool
func ReleaseAlibabaJymItemExternalGoodsBatchOnsaleAPIRequest(v *AlibabaJymItemExternalGoodsBatchOnsaleAPIRequest) {
	v.Reset()
	poolAlibabaJymItemExternalGoodsBatchOnsaleAPIRequest.Put(v)
}
