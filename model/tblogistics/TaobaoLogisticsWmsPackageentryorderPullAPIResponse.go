package tblogistics

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsWmsPackageentryorderPullAPIResponse 包裹入库单拉单 API返回值
// taobao.logistics.wms.packageentryorder.pull
//
// 包裹入库单拉单
type TaobaoLogisticsWmsPackageentryorderPullAPIResponse struct {
	model.CommonResponse
	TaobaoLogisticsWmsPackageentryorderPullAPIResponseModel
}

// TaobaoLogisticsWmsPackageentryorderPullAPIResponseModel is 包裹入库单拉单 成功返回结果
type TaobaoLogisticsWmsPackageentryorderPullAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_wms_packageentryorder_pull_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *ResultWrapper `json:"result,omitempty" xml:"result,omitempty"`
}
