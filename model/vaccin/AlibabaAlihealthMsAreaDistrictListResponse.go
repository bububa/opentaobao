package vaccin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
疫苗预约地市信息查询 API返回值 
alibaba.alihealth.ms.area.district.list

微信小程序疫苗预约地市信息查询
*/
type AlibabaAlihealthMsAreaDistrictListAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthMsAreaDistrictListResponse
}

// 疫苗预约地市信息查询 成功返回结果
type AlibabaAlihealthMsAreaDistrictListResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_ms_area_district_list_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}
