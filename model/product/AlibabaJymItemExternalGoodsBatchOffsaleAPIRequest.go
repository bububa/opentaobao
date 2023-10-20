package product

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaJymItemExternalGoodsBatchOffsaleAPIRequest 交易猫外部商家批量下架商品接口 API请求
// alibaba.jym.item.external.goods.batch.offsale
//
// 供外部B端商家接入，提交批量下架商品请求，返回批量下架任务结果
type AlibabaJymItemExternalGoodsBatchOffsaleAPIRequest struct {
	model.Params
	// 商品批量下架接口入参
	_goodsOffSaleCommand *GoodsOffSaleCommandDto
}

// NewAlibabaJymItemExternalGoodsBatchOffsaleRequest 初始化AlibabaJymItemExternalGoodsBatchOffsaleAPIRequest对象
func NewAlibabaJymItemExternalGoodsBatchOffsaleRequest() *AlibabaJymItemExternalGoodsBatchOffsaleAPIRequest {
	return &AlibabaJymItemExternalGoodsBatchOffsaleAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaJymItemExternalGoodsBatchOffsaleAPIRequest) Reset() {
	r._goodsOffSaleCommand = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaJymItemExternalGoodsBatchOffsaleAPIRequest) GetApiMethodName() string {
	return "alibaba.jym.item.external.goods.batch.offsale"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaJymItemExternalGoodsBatchOffsaleAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaJymItemExternalGoodsBatchOffsaleAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetGoodsOffSaleCommand is GoodsOffSaleCommand Setter
// 商品批量下架接口入参
func (r *AlibabaJymItemExternalGoodsBatchOffsaleAPIRequest) SetGoodsOffSaleCommand(_goodsOffSaleCommand *GoodsOffSaleCommandDto) error {
	r._goodsOffSaleCommand = _goodsOffSaleCommand
	r.Set("goods_off_sale_command", _goodsOffSaleCommand)
	return nil
}

// GetGoodsOffSaleCommand GoodsOffSaleCommand Getter
func (r AlibabaJymItemExternalGoodsBatchOffsaleAPIRequest) GetGoodsOffSaleCommand() *GoodsOffSaleCommandDto {
	return r._goodsOffSaleCommand
}

var poolAlibabaJymItemExternalGoodsBatchOffsaleAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaJymItemExternalGoodsBatchOffsaleRequest()
	},
}

// GetAlibabaJymItemExternalGoodsBatchOffsaleRequest 从 sync.Pool 获取 AlibabaJymItemExternalGoodsBatchOffsaleAPIRequest
func GetAlibabaJymItemExternalGoodsBatchOffsaleAPIRequest() *AlibabaJymItemExternalGoodsBatchOffsaleAPIRequest {
	return poolAlibabaJymItemExternalGoodsBatchOffsaleAPIRequest.Get().(*AlibabaJymItemExternalGoodsBatchOffsaleAPIRequest)
}

// ReleaseAlibabaJymItemExternalGoodsBatchOffsaleAPIRequest 将 AlibabaJymItemExternalGoodsBatchOffsaleAPIRequest 放入 sync.Pool
func ReleaseAlibabaJymItemExternalGoodsBatchOffsaleAPIRequest(v *AlibabaJymItemExternalGoodsBatchOffsaleAPIRequest) {
	v.Reset()
	poolAlibabaJymItemExternalGoodsBatchOffsaleAPIRequest.Put(v)
}
