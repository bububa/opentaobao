package alihouse

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
删除楼盘预售证 API返回值 
alibaba.alihouse.newhome.project.presalepermit.delete

删除楼盘预售证信息
*/
type AlibabaAlihouseNewhomeProjectPresalepermitDeleteAPIResponse struct {
    model.CommonResponse
    AlibabaAlihouseNewhomeProjectPresalepermitDeleteAPIResponseModel
}

// 删除楼盘预售证 成功返回结果
type AlibabaAlihouseNewhomeProjectPresalepermitDeleteAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_alihouse_newhome_project_presalepermit_delete_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 接口返回model
    Result   *AlibabaAlihouseNewhomeProjectPresalepermitDeleteResult `json:"result,omitempty" xml:"result,omitempty"`
}
