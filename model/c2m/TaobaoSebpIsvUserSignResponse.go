package c2m

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
淘小铺三方签约同步 APIResponse
taobao.sebp.isv.user.sign

同步淘小铺三方服务签约信息
*/
type TaobaoSebpIsvUserSignAPIResponse struct {
    model.CommonResponse
    TaobaoSebpIsvUserSignResponse
}

type TaobaoSebpIsvUserSignResponse struct {
    XMLName xml.Name `xml:"sebp_isv_user_sign_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 是否成功
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`

    
}
