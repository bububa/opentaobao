package baichuan

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
百川H5登录二次验证 APIResponse
taobao.baichuan.user.logindoublecheck

百川H5登录二次验证
*/
type TaobaoBaichuanUserLogindoublecheckAPIResponse struct {
    model.CommonResponse
    TaobaoBaichuanUserLogindoublecheckResponse
}

type TaobaoBaichuanUserLogindoublecheckResponse struct {
    XMLName xml.Name `xml:"baichuan_user_logindoublecheck_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // name
    
    Name   string `json:"name,omitempty" xml:"name,omitempty"`

    
}
