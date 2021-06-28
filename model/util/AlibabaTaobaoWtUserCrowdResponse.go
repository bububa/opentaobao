package util

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
校验用户是否为新人 APIResponse
alibaba.taobao.wt.user.crowd

增加isv接口
根据入参手机号判断是否为新人类型
*/
type AlibabaTaobaoWtUserCrowdAPIResponse struct {
    model.CommonResponse
    AlibabaTaobaoWtUserCrowdResponse
}

type AlibabaTaobaoWtUserCrowdResponse struct {
    XMLName xml.Name `xml:"alibaba_taobao_wt_user_crowd_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 接口结果
    
    Result   *CommonResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
