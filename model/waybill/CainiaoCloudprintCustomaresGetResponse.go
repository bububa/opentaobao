package waybill

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取商家的自定义区模板信息 APIResponse
cainiao.cloudprint.customares.get

供isv使用，获取商家的自定义区的模板信息
*/
type CainiaoCloudprintCustomaresGetAPIResponse struct {
    model.CommonResponse
    CainiaoCloudprintCustomaresGetResponse
}

type CainiaoCloudprintCustomaresGetResponse struct {
    XMLName xml.Name `xml:"cainiao_cloudprint_customares_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 结果
    
    Result   *CloudPrintBaseResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
