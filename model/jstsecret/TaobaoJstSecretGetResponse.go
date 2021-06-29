package jstsecret

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取订单消费者的隐私号码 APIResponse
taobao.jst.secret.get

根据订单号获取消费者的隐私号
*/
type TaobaoJstSecretGetAPIResponse struct {
    model.CommonResponse
    TaobaoJstSecretGetResponse
}

type TaobaoJstSecretGetResponse struct {
    XMLName xml.Name `xml:"jst_secret_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 结果
    
    Secret   *SecretNoBindResponse `json:"secret,omitempty" xml:"secret,omitempty"`

    
}
