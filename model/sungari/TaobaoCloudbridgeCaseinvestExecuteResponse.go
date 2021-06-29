package sungari

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
红盾云桥案件协查服务 APIResponse
taobao.cloudbridge.caseinvest.execute

通过API接口直接提供政府部门录入及查询函件服务
*/
type TaobaoCloudbridgeCaseinvestExecuteAPIResponse struct {
    model.CommonResponse
    TaobaoCloudbridgeCaseinvestExecuteResponse
}

type TaobaoCloudbridgeCaseinvestExecuteResponse struct {
    XMLName xml.Name `xml:"cloudbridge_caseinvest_execute_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *TaobaoCloudbridgeCaseinvestExecuteResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
