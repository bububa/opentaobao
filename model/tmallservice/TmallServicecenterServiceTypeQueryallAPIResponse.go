package tmallservice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
服务供应链服务类型 API返回值 
tmall.servicecenter.service.type.queryall

查询天猫服务类型列表
*/
type TmallServicecenterServiceTypeQueryallAPIResponse struct {
    model.CommonResponse
    TmallServicecenterServiceTypeQueryallAPIResponseModel
}

// 服务供应链服务类型 成功返回结果
type TmallServicecenterServiceTypeQueryallAPIResponseModel struct {
    XMLName xml.Name `xml:"tmall_servicecenter_service_type_queryall_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *ResultBase `json:"result,omitempty" xml:"result,omitempty"`
}
