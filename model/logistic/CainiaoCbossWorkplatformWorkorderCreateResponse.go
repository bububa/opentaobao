package logistic

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
菜鸟工单创建接口 API返回值 
cainiao.cboss.workplatform.workorder.create

菜鸟工单创建接口，目前调用者ISV
*/
type CainiaoCbossWorkplatformWorkorderCreateAPIResponse struct {
    model.CommonResponse
    CainiaoCbossWorkplatformWorkorderCreateResponse
}

// 菜鸟工单创建接口 成功返回结果
type CainiaoCbossWorkplatformWorkorderCreateResponse struct {
    XMLName xml.Name `xml:"cainiao_cboss_workplatform_workorder_create_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 接口返回model
    Result   *CainiaoCbossWorkplatformWorkorderCreateResult `json:"result,omitempty" xml:"result,omitempty"`
}
