package tmallservice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商绑定工人 APIResponse
alibaba.ssc.supplyplatform.serviceworker.save

服务商将上传工人与服务商自己建立关系，需要将工人的服务区域和住址回传
*/
type AlibabaSscSupplyplatformServiceworkerSaveAPIResponse struct {
    model.CommonResponse
    AlibabaSscSupplyplatformServiceworkerSaveResponse
}

type AlibabaSscSupplyplatformServiceworkerSaveResponse struct {
    XMLName xml.Name `xml:"alibaba_ssc_supplyplatform_serviceworker_save_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *AlibabaSscSupplyplatformServiceworkerSaveResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
