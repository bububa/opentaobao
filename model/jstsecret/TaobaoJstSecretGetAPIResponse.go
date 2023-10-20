package jstsecret

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJstSecretGetAPIResponse 获取订单消费者的隐私号码 API返回值
// taobao.jst.secret.get
//
// 根据订单号获取消费者的隐私号
type TaobaoJstSecretGetAPIResponse struct {
	model.CommonResponse
	TaobaoJstSecretGetAPIResponseModel
}

// TaobaoJstSecretGetAPIResponseModel is 获取订单消费者的隐私号码 成功返回结果
type TaobaoJstSecretGetAPIResponseModel struct {
	XMLName xml.Name `xml:"jst_secret_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Secret *SecretNoBindResponse `json:"secret,omitempty" xml:"secret,omitempty"`
}
