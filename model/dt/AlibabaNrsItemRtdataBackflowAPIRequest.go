package dt

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaNrsItemRtdataBackflowAPIRequest RT竞价数据回流 API请求
// alibaba.nrs.item.rtdata.backflow
//
// 回流竞品价格数据，用与后续OCR识别价签数据，做精确化数据纠正
type AlibabaNrsItemRtdataBackflowAPIRequest struct {
	model.Params
	// 入参
	_rtItemPriceTagBackParam *RtItemPriceTagBackParam
}

// NewAlibabaNrsItemRtdataBackflowRequest 初始化AlibabaNrsItemRtdataBackflowAPIRequest对象
func NewAlibabaNrsItemRtdataBackflowRequest() *AlibabaNrsItemRtdataBackflowAPIRequest {
	return &AlibabaNrsItemRtdataBackflowAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaNrsItemRtdataBackflowAPIRequest) GetApiMethodName() string {
	return "alibaba.nrs.item.rtdata.backflow"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaNrsItemRtdataBackflowAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetRtItemPriceTagBackParam is RtItemPriceTagBackParam Setter
// 入参
func (r *AlibabaNrsItemRtdataBackflowAPIRequest) SetRtItemPriceTagBackParam(_rtItemPriceTagBackParam *RtItemPriceTagBackParam) error {
	r._rtItemPriceTagBackParam = _rtItemPriceTagBackParam
	r.Set("rt_item_price_tag_back_param", _rtItemPriceTagBackParam)
	return nil
}

// GetRtItemPriceTagBackParam RtItemPriceTagBackParam Getter
func (r AlibabaNrsItemRtdataBackflowAPIRequest) GetRtItemPriceTagBackParam() *RtItemPriceTagBackParam {
	return r._rtItemPriceTagBackParam
}
