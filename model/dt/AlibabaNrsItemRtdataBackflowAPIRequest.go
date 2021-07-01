package dt

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaNrsItemRtdataBackflowAPIRequest
RT竞价数据回流 API请求
alibaba.nrs.item.rtdata.backflow

回流竞品价格数据，用与后续OCR识别价签数据，做精确化数据纠正 */
type AlibabaNrsItemRtdataBackflowAPIRequest struct {
	model.Params
	// 入参
	_rtItemPriceTagBackParam *RtItemPriceTagBackParam
}

// New
