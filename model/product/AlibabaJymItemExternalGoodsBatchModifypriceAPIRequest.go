package product

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaJymItemExternalGoodsBatchModifypriceAPIRequest 交易猫外部商家批量商品改价接口 API请求
// alibaba.jym.item.external.goods.batch.modifyprice
//
// 供外部B端商家接入，提交批量商品改价请求，返回批量改价任务结果
type AlibabaJymItemExternalGoodsBatchModifypriceAPIRequest struct {
	model.Params
	// 商品批量改价接口入参
	_goodsPriceModifyCommand *GoodsPriceModifyCommandDto
}

// NewAlibabaJymItemExternalGoodsBatchModifypriceRequest 初始化AlibabaJymItemExternalGoodsBatchModifypriceAPIRequest对象
func NewAlibabaJymItemExternalGoodsBatchModifypriceRequest() *AlibabaJymItemExternalGoodsBatchModifypriceAPIRequest {
	return &AlibabaJymItemExternalGoodsBatchModifypriceAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaJymItemExternalGoodsBatchModifypriceAPIRequest) Reset() {
	r._goodsPriceModifyCommand = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaJymItemExternalGoodsBatchModifypriceAPIRequest) GetApiMethodName() string {
	return "alibaba.jym.item.external.goods.batch.modifyprice"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaJymItemExternalGoodsBatchModifypriceAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaJymItemExternalGoodsBatchModifypriceAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetGoodsPriceModifyCommand is GoodsPriceModifyCommand Setter
// 商品批量改价接口入参
func (r *AlibabaJymItemExternalGoodsBatchModifypriceAPIRequest) SetGoodsPriceModifyCommand(_goodsPriceModifyCommand *GoodsPriceModifyCommandDto) error {
	r._goodsPriceModifyCommand = _goodsPriceModifyCommand
	r.Set("goods_price_modify_command", _goodsPriceModifyCommand)
	return nil
}

// GetGoodsPriceModifyCommand GoodsPriceModifyCommand Getter
func (r AlibabaJymItemExternalGoodsBatchModifypriceAPIRequest) GetGoodsPriceModifyCommand() *GoodsPriceModifyCommandDto {
	return r._goodsPriceModifyCommand
}

var poolAlibabaJymItemExternalGoodsBatchModifypriceAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaJymItemExternalGoodsBatchModifypriceRequest()
	},
}

// GetAlibabaJymItemExternalGoodsBatchModifypriceRequest 从 sync.Pool 获取 AlibabaJymItemExternalGoodsBatchModifypriceAPIRequest
func GetAlibabaJymItemExternalGoodsBatchModifypriceAPIRequest() *AlibabaJymItemExternalGoodsBatchModifypriceAPIRequest {
	return poolAlibabaJymItemExternalGoodsBatchModifypriceAPIRequest.Get().(*AlibabaJymItemExternalGoodsBatchModifypriceAPIRequest)
}

// ReleaseAlibabaJymItemExternalGoodsBatchModifypriceAPIRequest 将 AlibabaJymItemExternalGoodsBatchModifypriceAPIRequest 放入 sync.Pool
func ReleaseAlibabaJymItemExternalGoodsBatchModifypriceAPIRequest(v *AlibabaJymItemExternalGoodsBatchModifypriceAPIRequest) {
	v.Reset()
	poolAlibabaJymItemExternalGoodsBatchModifypriceAPIRequest.Put(v)
}
