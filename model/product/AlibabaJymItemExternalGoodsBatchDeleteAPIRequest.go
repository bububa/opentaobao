package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaJymItemExternalGoodsBatchDeleteAPIRequest 交易猫外部商家批量删除商品接口 API请求
// alibaba.jym.item.external.goods.batch.delete
//
// 交易猫外部商家批量删除商品接口
type AlibabaJymItemExternalGoodsBatchDeleteAPIRequest struct {
	model.Params
	// 商品批量删除接口入参
	_goodsDeleteCommandDto *GoodsDeleteCommandDto
}

// NewAlibabaJymItemExternalGoodsBatchDeleteRequest 初始化AlibabaJymItemExternalGoodsBatchDeleteAPIRequest对象
func NewAlibabaJymItemExternalGoodsBatchDeleteRequest() *AlibabaJymItemExternalGoodsBatchDeleteAPIRequest {
	return &AlibabaJymItemExternalGoodsBatchDeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaJymItemExternalGoodsBatchDeleteAPIRequest) GetApiMethodName() string {
	return "alibaba.jym.item.external.goods.batch.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaJymItemExternalGoodsBatchDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaJymItemExternalGoodsBatchDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetGoodsDeleteCommandDto is GoodsDeleteCommandDto Setter
// 商品批量删除接口入参
func (r *AlibabaJymItemExternalGoodsBatchDeleteAPIRequest) SetGoodsDeleteCommandDto(_goodsDeleteCommandDto *GoodsDeleteCommandDto) error {
	r._goodsDeleteCommandDto = _goodsDeleteCommandDto
	r.Set("goods_delete_command_dto", _goodsDeleteCommandDto)
	return nil
}

// GetGoodsDeleteCommandDto GoodsDeleteCommandDto Getter
func (r AlibabaJymItemExternalGoodsBatchDeleteAPIRequest) GetGoodsDeleteCommandDto() *GoodsDeleteCommandDto {
	return r._goodsDeleteCommandDto
}
