package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkUmsReturnitemsGetAPIResponse
退货库位商品查询（退货出库辅助）-回流单 API返回值
alibaba.wdk.ums.returnitems.get

退货库位商品查询（退货出库辅助）-回流单 */
type AlibabaWdkUmsReturnitemsGetAPIResponse struct {
	model.CommonResponse
	AlibabaWdkUmsReturnitemsGetAPIResponseModel
}

// AlibabaWdkUmsReturnitemsGetAPIResponseModel is 退货库位商品查询（退货出库辅助）-回流单 成功返回结果
type AlibabaWdkUmsReturnitemsGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_ums_returnitems_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// reslut
	Result *UtmsResult `json:"result,omitempty" xml:"result,omitempty"`
}
