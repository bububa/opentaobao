package sungari

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
红盾云桥案件协查服务 API返回值 
taobao.cloudbridge.caseinvest.execute

通过API接口直接提供政府部门录入及查询函件服务
*/
type TaobaoCloudbridgeCaseinvestExecuteAPIResponse struct {
    model.CommonResponse
    TaobaoCloudbridgeCaseinvestExecuteAPIResponseModel
}

// 红盾云桥案件协查服务 成功返回结果
type TaobaoCloudbridgeCaseinvestExecuteAPIResponseModel struct {
    XMLName xml.Name `xml:"cloudbridge_caseinvest_execute_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *TaobaoCloudbridgeCaseinvestExecuteResult `json:"result,omitempty" xml:"result,omitempty"`
}
