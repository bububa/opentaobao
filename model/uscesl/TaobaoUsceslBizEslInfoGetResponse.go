package uscesl

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
价签设备信息查询接口 APIResponse
taobao.uscesl.biz.esl.info.get

价签设备信息查询接口
*/
type TaobaoUsceslBizEslInfoGetAPIResponse struct {
    model.CommonResponse
    TaobaoUsceslBizEslInfoGetResponse
}

type TaobaoUsceslBizEslInfoGetResponse struct {
    XMLName xml.Name `xml:"uscesl_biz_esl_info_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result.sucess表示是否成功，target是设备信息
    
    Result   string `json:"result,omitempty" xml:"result,omitempty"`

    
}
