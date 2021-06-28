package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据营销积木块代码获取积木块 APIResponse
taobao.ump.mbb.getbycode

根据营销积木块代码获取积木块。接口返回该代码最新版本的积木块。如果要查询某个非最新版本的积木块，可以使用积木块id查询接口。
*/
type TaobaoUmpMbbGetbycodeAPIResponse struct {
    model.CommonResponse
    TaobaoUmpMbbGetbycodeResponse
}

type TaobaoUmpMbbGetbycodeResponse struct {
    XMLName xml.Name `xml:"ump_mbb_getbycode_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 营销积木块的内容，通过ump sdk来进行处理
    
    Mbb   string `json:"mbb,omitempty" xml:"mbb,omitempty"`

    
}
