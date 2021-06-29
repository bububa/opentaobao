package baichuan

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
百川登录二次验证 APIResponse
taobao.baichuan.openaccount.logindoublecheck

百川登录二次验证
*/
type TaobaoBaichuanOpenaccountLogindoublecheckAPIResponse struct {
    model.CommonResponse
    TaobaoBaichuanOpenaccountLogindoublecheckResponse
}

type TaobaoBaichuanOpenaccountLogindoublecheckResponse struct {
    XMLName xml.Name `xml:"baichuan_openaccount_logindoublecheck_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // name
    
    Name   string `json:"name,omitempty" xml:"name,omitempty"`

    
}
