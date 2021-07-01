package baichuan

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询app在设备上的安装信息 API返回值 
alibaba.baichuan.aso.query

查询app在设备上的安装信息
*/
type AlibabaBaichuanAsoQueryAPIResponse struct {
    model.CommonResponse
    AlibabaBaichuanAsoQueryAPIResponseModel
}

// 查询app在设备上的安装信息 成功返回结果
type AlibabaBaichuanAsoQueryAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_baichuan_aso_query_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *AsoQueryDeviceResult `json:"result,omitempty" xml:"result,omitempty"`
}
