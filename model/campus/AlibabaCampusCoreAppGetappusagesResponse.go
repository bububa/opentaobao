package campus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据应用ID获得应用所在的园区 API返回值 
alibaba.campus.core.app.getappusages

传入应用的id,  获得用户授权的园区
*/
type AlibabaCampusCoreAppGetappusagesAPIResponse struct {
    model.CommonResponse
    AlibabaCampusCoreAppGetappusagesResponse
}

// 根据应用ID获得应用所在的园区 成功返回结果
type AlibabaCampusCoreAppGetappusagesResponse struct {
    XMLName xml.Name `xml:"alibaba_campus_core_app_getappusages_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *CollectionResult `json:"result,omitempty" xml:"result,omitempty"`
}
