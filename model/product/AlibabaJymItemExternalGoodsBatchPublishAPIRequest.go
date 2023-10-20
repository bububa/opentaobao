package product

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaJymItemExternalGoodsBatchPublishAPIRequest 交易猫外部商家批量发布商品接口 API请求
// alibaba.jym.item.external.goods.batch.publish
//
// 供外部B端商家接入，提交批量发布商品请求，返回批量发布任务结果
type AlibabaJymItemExternalGoodsBatchPublishAPIRequest struct {
	model.Params
	// 商品批量发布接口入参
	_goodsPublishCommand *GoodsPublishCommandDto
}

// NewAlibabaJymItemExternalGoodsBatchPublishRequest 初始化AlibabaJymItemExternalGoodsBatchPublishAPIRequest对象
func NewAlibabaJymItemExternalGoodsBatchPublishRequest() *AlibabaJymItemExternalGoodsBatchPublishAPIRequest {
	return &AlibabaJymItemExternalGoodsBatchPublishAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaJymItemExternalGoodsBatchPublishAPIRequest) Reset() {
	r._goodsPublishCommand = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaJymItemExternalGoodsBatchPublishAPIRequest) GetApiMethodName() string {
	return "alibaba.jym.item.external.goods.batch.publish"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaJymItemExternalGoodsBatchPublishAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaJymItemExternalGoodsBatchPublishAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetGoodsPublishCommand is GoodsPublishCommand Setter
// 商品批量发布接口入参
func (r *AlibabaJymItemExternalGoodsBatchPublishAPIRequest) SetGoodsPublishCommand(_goodsPublishCommand *GoodsPublishCommandDto) error {
	r._goodsPublishCommand = _goodsPublishCommand
	r.Set("goods_publish_command", _goodsPublishCommand)
	return nil
}

// GetGoodsPublishCommand GoodsPublishCommand Getter
func (r AlibabaJymItemExternalGoodsBatchPublishAPIRequest) GetGoodsPublishCommand() *GoodsPublishCommandDto {
	return r._goodsPublishCommand
}

var poolAlibabaJymItemExternalGoodsBatchPublishAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaJymItemExternalGoodsBatchPublishRequest()
	},
}

// GetAlibabaJymItemExternalGoodsBatchPublishRequest 从 sync.Pool 获取 AlibabaJymItemExternalGoodsBatchPublishAPIRequest
func GetAlibabaJymItemExternalGoodsBatchPublishAPIRequest() *AlibabaJymItemExternalGoodsBatchPublishAPIRequest {
	return poolAlibabaJymItemExternalGoodsBatchPublishAPIRequest.Get().(*AlibabaJymItemExternalGoodsBatchPublishAPIRequest)
}

// ReleaseAlibabaJymItemExternalGoodsBatchPublishAPIRequest 将 AlibabaJymItemExternalGoodsBatchPublishAPIRequest 放入 sync.Pool
func ReleaseAlibabaJymItemExternalGoodsBatchPublishAPIRequest(v *AlibabaJymItemExternalGoodsBatchPublishAPIRequest) {
	v.Reset()
	poolAlibabaJymItemExternalGoodsBatchPublishAPIRequest.Put(v)
}
