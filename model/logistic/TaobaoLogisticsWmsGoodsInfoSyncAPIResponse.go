package logistic

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsWmsGoodsInfoSyncAPIResponse WMS回传货品长宽高图片等信息 API返回值
// taobao.logistics.wms.goods.info.sync
//
// WMS回传货品长宽高图片等信息
type TaobaoLogisticsWmsGoodsInfoSyncAPIResponse struct {
	model.CommonResponse
	TaobaoLogisticsWmsGoodsInfoSyncAPIResponseModel
}

// TaobaoLogisticsWmsGoodsInfoSyncAPIResponseModel is WMS回传货品长宽高图片等信息 成功返回结果
type TaobaoLogisticsWmsGoodsInfoSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_wms_goods_info_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	BizErrorMessage string `json:"biz_error_message,omitempty" xml:"biz_error_message,omitempty"`
	// 错误编码
	BizErrorCode string `json:"biz_error_code,omitempty" xml:"biz_error_code,omitempty"`
	// 是否成功
	Suc bool `json:"suc,omitempty" xml:"suc,omitempty"`
	// 是否支持重试
	Retry bool `json:"retry,omitempty" xml:"retry,omitempty"`
}
