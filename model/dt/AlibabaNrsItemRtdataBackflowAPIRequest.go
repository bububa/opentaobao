package dt

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabanrsitemrtdatabackflowAPIRequest RT竞价数据回流 API请求
// alibaba.nrs.item.rtdata.backflow
//
// 回流竞品价格数据，用与后续OCR识别价签数据，做精确化数据纠正
type AlibabanrsitemrtdatabackflowAPIRequest struct {
	model.Params
	// 入参
	_rtItemPriceTagBackParam *RtItemPriceTagBackParam
}

// NewAlibabanrsitemrtdatabackflowRequest 初始化AlibabanrsitemrtdatabackflowAPIRequest对象
func NewAlibabanrsitemrtdatabackflowRequest() *AlibabanrsitemrtdatabackflowAPIRequest {
	return &AlibabanrsitemrtdatabackflowAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabanrsitemrtdatabackflowAPIRequest) GetApiMethodName() string {
	return "alibaba.nrs.item.rtdata.backflow"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabanrsitemrtdatabackflowAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabanrsitemrtdatabackflowAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRtItemPriceTagBackParam is RtItemPriceTagBackParam Setter
// 入参
func (r *AlibabanrsitemrtdatabackflowAPIRequest) SetRtItemPriceTagBackParam(_rtItemPriceTagBackParam *RtItemPriceTagBackParam) error {
	r._rtItemPriceTagBackParam = _rtItemPriceTagBackParam
	r.Set("rt_item_price_tag_back_param", _rtItemPriceTagBackParam)
	return nil
}

// GetRtItemPriceTagBackParam RtItemPriceTagBackParam Getter
func (r AlibabanrsitemrtdatabackflowAPIRequest) GetRtItemPriceTagBackParam() *RtItemPriceTagBackParam {
	return r._rtItemPriceTagBackParam
}
