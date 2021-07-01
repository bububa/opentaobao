package waybill

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取商家使用的标准模板 API返回值 
cainiao.cloudprint.isvtemplates.get

获取商家使用的标准模板
*/
type CainiaoCloudprintIsvtemplatesGetAPIResponse struct {
    model.CommonResponse
    CainiaoCloudprintIsvtemplatesGetAPIResponseModel
}

// 获取商家使用的标准模板 成功返回结果
type CainiaoCloudprintIsvtemplatesGetAPIResponseModel struct {
    XMLName xml.Name `xml:"cainiao_cloudprint_isvtemplates_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *CloudPrintBaseResult `json:"result,omitempty" xml:"result,omitempty"`
}
