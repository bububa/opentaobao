package tblogistics

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaologisticswmspackageexceptionreportAPIResponse 无主件回告 API返回值
// taobao.logistics.wms.packageexception.report
//
// 无主件回告
type TaobaologisticswmspackageexceptionreportAPIResponse struct {
	model.CommonResponse
	TaobaologisticswmspackageexceptionreportAPIResponseModel
}

// TaobaologisticswmspackageexceptionreportAPIResponseModel is 无主件回告 成功返回结果
type TaobaologisticswmspackageexceptionreportAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_wms_packageexception_report_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应信息
	Result *ResultWrapper `json:"result,omitempty" xml:"result,omitempty"`
}
