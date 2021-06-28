package baichuan

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询app在设备上的安装信息 APIResponse
alibaba.baichuan.aso.query

查询app在设备上的安装信息
*/
type AlibabaBaichuanAsoQueryAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaBaichuanAsoQueryResponse `json:"alibaba_baichuan_aso_query_response,omitempty"` 
    AlibabaBaichuanAsoQueryResponse
}

/* model for simplify = false
type AlibabaBaichuanAsoQueryResponse struct {

    // result
    
    Result  *struct {
        AsoQueryDeviceResult  *AsoQueryDeviceResult `json:"aso_query_device_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaBaichuanAsoQueryResponse struct {

    // result
    Result   *AsoQueryDeviceResult `json:"result,omitempty"`

}
