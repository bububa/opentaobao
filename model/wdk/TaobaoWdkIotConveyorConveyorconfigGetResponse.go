package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取悬挂链基本配置信息 APIResponse
taobao.wdk.iot.conveyor.conveyorconfig.get

用于从云端WCS获取悬挂链基本配置信息
*/
type TaobaoWdkIotConveyorConveyorconfigGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoWdkIotConveyorConveyorconfigGetResponse `json:"wdk_iot_conveyor_conveyorconfig_get_response,omitempty"` 
    TaobaoWdkIotConveyorConveyorconfigGetResponse
}

/* model for simplify = false
type TaobaoWdkIotConveyorConveyorconfigGetResponse struct {

    // 返回值
    
    Result  *struct {
        HMResult  *HMResult `json:"hm_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoWdkIotConveyorConveyorconfigGetResponse struct {

    // 返回值
    Result   *HMResult `json:"result,omitempty"`

}
