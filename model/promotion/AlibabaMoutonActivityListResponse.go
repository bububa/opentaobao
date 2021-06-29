package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商家自运营活动列表 APIResponse
alibaba.mouton.activity.list

商家查询自己配置的活动列表
*/
type AlibabaMoutonActivityListAPIResponse struct {
    model.CommonResponse
    AlibabaMoutonActivityListResponse
}

type AlibabaMoutonActivityListResponse struct {
    XMLName xml.Name `xml:"alibaba_mouton_activity_list_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *Page `json:"result,omitempty" xml:"result,omitempty"`

    
}
