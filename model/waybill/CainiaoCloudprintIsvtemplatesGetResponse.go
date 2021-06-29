package waybill

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取商家使用的标准模板 APIResponse
cainiao.cloudprint.isvtemplates.get

获取商家使用的标准模板
*/
type CainiaoCloudprintIsvtemplatesGetAPIResponse struct {
    model.CommonResponse
    CainiaoCloudprintIsvtemplatesGetResponse
}

type CainiaoCloudprintIsvtemplatesGetResponse struct {
    XMLName xml.Name `xml:"cainiao_cloudprint_isvtemplates_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *CloudPrintBaseResult `json:"result,omitempty" xml:"result,omitempty"`

    
    // result
    
    Result   *CloudPrintBaseResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
