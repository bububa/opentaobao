package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取悬挂链基本配置信息 APIResponse
taobao.wdk.iot.conveyor.conveyorconfig.get

用于从云端WCS获取悬挂链基本配置信息
*/
type TaobaoWdkIotConveyorConveyorconfigGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"wdk_iot_conveyor_conveyorconfig_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回值
    
    Result   *HMResult `json:"result,omitempty" xml:"