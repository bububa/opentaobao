package degoperation

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
激励抽奖 APIResponse
taobao.degoperation.do.luckydraw

激励平台抽奖接口。用户可以通过接口完成抽奖功能
*/
type TaobaoDegoperationDoLuckydrawAPIResponse struct {
    model.CommonResponse
    TaobaoDegoperationDoLuckydrawResponse
}

type TaobaoDegoperationDoLuckydrawResponse struct {
    XMLName xml.Name `xml:"degoperation_do_luckydraw_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *BonusResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
