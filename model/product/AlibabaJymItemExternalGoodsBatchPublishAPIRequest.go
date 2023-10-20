package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabajymitemexternalgoodsbatchpublishAPIRequest 交易猫外部商家批量发布商品接口 API请求
// alibaba.jym.item.external.goods.batch.publish
//
// 供外部B端商家接入，提交批量发布商品请求，返回批量发布任务结果
type AlibabajymitemexternalgoodsbatchpublishAPIRequest struct {
	model.Params
	// 商品批量发布接口入参
	_goodsPublishCommand *GoodsPublishCommandDto
}

// NewAlibabajymitemexternalgoodsbatchpublishRequest 初始化AlibabajymitemexternalgoodsbatchpublishAPIRequest对象
func NewAlibabajymitemexternalgoodsbatchpublishRequest() *AlibabajymitemexternalgoodsbatchpublishAPIRequest {
	return &AlibabajymitemexternalgoodsbatchpublishAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabajymitemexternalgoodsbatchpublishAPIRequest) GetApiMethodName() string {
	return "alibaba.jym.item.external.goods.batch.publish"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabajymitemexternalgoodsbatchpublishAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabajymitemexternalgoodsbatchpublishAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetGoodsPublishCommand is GoodsPublishCommand Setter
// 商品批量发布接口入参
func (r *AlibabajymitemexternalgoodsbatchpublishAPIRequest) SetGoodsPublishCommand(_goodsPublishCommand *GoodsPublishCommandDto) error {
	r._goodsPublishCommand = _goodsPublishCommand
	r.Set("goods_publish_command", _goodsPublishCommand)
	return nil
}

// GetGoodsPublishCommand GoodsPublishCommand Getter
func (r AlibabajymitemexternalgoodsbatchpublishAPIRequest) GetGoodsPublishCommand() *GoodsPublishCommandDto {
	return r._goodsPublishCommand
}
