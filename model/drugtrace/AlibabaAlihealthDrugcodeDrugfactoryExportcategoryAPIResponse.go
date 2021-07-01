package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
导出临床药品目录 API返回值 
alibaba.alihealth.drugcode.drugfactory.exportcategory

导出临床药品目录
*/
type AlibabaAlihealthDrugcodeDrugfactoryExportcategoryAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDrugcodeDrugfactoryExportcategoryAPIResponseModel
}

// 导出临床药品目录 成功返回结果
type AlibabaAlihealthDrugcodeDrugfactoryExportcategoryAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_alihealth_drugcode_drugfactory_exportcategory_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 导出的药品目录
    Model   []ProductDto `json:"model,omitempty" xml:"model>product_dto,omitempty"`
    // 操作码
    MsgCode   string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
    // 操作说明
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
}
