package waybill

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取商家单一自定义区 APIResponse
cainiao.cloudprint.single.customarea.get

商家所有快递公司模板只有一个自定义区
*/
type CainiaoCloudprintSingleCustomareaGetAPIResponse struct {
    model.CommonResponse
    CainiaoCloudprintSingleCustomareaGetResponse
}

type CainiaoCloudprintSingleCustomareaGetResponse struct {
    XMLName xml.Name `xml:"cainiao_cloudprint_single_customarea_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *CloudPrintBaseResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
