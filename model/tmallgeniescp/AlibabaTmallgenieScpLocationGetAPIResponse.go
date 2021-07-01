package tmallgeniescp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
2-IBP查询CDC和RDC数据接口 API返回值 
alibaba.tmallgenie.scp.location.get

天猫精灵供应链-计划域-IBP查询CDC和RDC数据接口
*/
type AlibabaTmallgenieScpLocationGetAPIResponse struct {
    model.CommonResponse
    AlibabaTmallgenieScpLocationGetAPIResponseModel
}

// 2-IBP查询CDC和RDC数据接口 成功返回结果
type AlibabaTmallgenieScpLocationGetAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_tmallgenie_scp_location_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回描述
    ResultMsg   string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
    // 返回码
    ResultCode   string `json:"result_code,omitempty" xml:"result_code,omitempty"`
    // 数据集合
    DataList   []AlibabaTmallgenieScpLocationGetData `json:"data_list,omitempty" xml:"data_list>alibaba_tmallgenie_scp_location_get_data,omitempty"`
    // 请求序列号
    TraceId   string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
}
