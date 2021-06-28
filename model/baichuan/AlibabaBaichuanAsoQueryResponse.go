package baichuan

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询app在设备上的安装信息 APIResponse
alibaba.baichuan.aso.query

查询app在设备上的安装信息
*/
type AlibabaBaichuanAsoQueryAPIResponse struct {
    model.CommonResponse
    AlibabaBaichuanAsoQueryResponse
}

type AlibabaBaichuanAsoQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_baichuan_aso_query_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *AsoQueryDeviceResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
