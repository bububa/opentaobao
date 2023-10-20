package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaowdkiotconveyorconveyorconfiggetAPIResponse 获取悬挂链基本配置信息 API返回值
// taobao.wdk.iot.conveyor.conveyorconfig.get
//
// 用于从云端WCS获取悬挂链基本配置信息
type TaobaowdkiotconveyorconveyorconfiggetAPIResponse struct {
	model.CommonResponse
	TaobaowdkiotconveyorconveyorconfiggetAPIResponseModel
}

// TaobaowdkiotconveyorconveyorconfiggetAPIResponseModel is 获取悬挂链基本配置信息 成功返回结果
type TaobaowdkiotconveyorconveyorconfiggetAPIResponseModel struct {
	XMLName xml.Name `xml:"wdk_iot_conveyor_conveyorconfig_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result *TaobaowdkiotconveyorconveyorconfiggetHmresult `json:"result,omitempty" xml:"result,omitempty"`
}
