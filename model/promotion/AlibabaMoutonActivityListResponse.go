package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商家自运营活动列表 API返回值 
alibaba.mouton.activity.list

商家查询自己配置的活动列表
*/
type AlibabaMoutonActivityListAPIResponse struct {
    model.CommonResponse
    AlibabaMoutonActivityListResponse
}

// 商家自运营活动列表 成功返回结果
type AlibabaMoutonActivityListResponse struct {
    XMLName xml.Name `xml:"alibaba_mouton_activity_list_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回结果
    Result   *Page `json:"result,omitempty" xml:"result,omitempty"`
}
