package waybill

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取用户使用的菜鸟电子面单模板信息 APIResponse
cainiao.cloudprint.mystdtemplates.get

获取用户使用的菜鸟电子面单
*/
type CainiaoCloudprintMystdtemplatesGetAPIResponse struct {
    model.CommonResponse
    CainiaoCloudprintMystdtemplatesGetResponse
}

type CainiaoCloudprintMystdtemplatesGetResponse struct {
    XMLName xml.Name `xml:"cainiao_cloudprint_mystdtemplates_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *CloudPrintBaseResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
