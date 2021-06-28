package traderate

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
飞猪通用评价接口 APIResponse
taobao.fliggy.wrate.getmixratelist

飞猪评价通用接口
*/
type TaobaoFliggyWrateGetmixratelistAPIResponse struct {
    model.CommonResponse
    TaobaoFliggyWrateGetmixratelistResponse
}

type TaobaoFliggyWrateGetmixratelistResponse struct {
    XMLName xml.Name `xml:"fliggy_wrate_getmixratelist_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 接口返回model
    
    Result   *TaobaoFliggyWrateGetmixratelistResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
