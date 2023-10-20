package mc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallmcrecordordersyncAPIResponse 订单信息同步 API返回值
// tmall.mc.record.order.sync
//
// 订单信息同步(零售云接口)
type TmallmcrecordordersyncAPIResponse struct {
	model.CommonResponse
	TmallmcrecordordersyncAPIResponseModel
}

// TmallmcrecordordersyncAPIResponseModel is 订单信息同步 成功返回结果
type TmallmcrecordordersyncAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_mc_record_order_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 同步成功
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
}
