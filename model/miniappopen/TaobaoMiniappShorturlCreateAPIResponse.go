package miniappopen

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoMiniappShorturlCreateAPIResponse
生成淘宝小程序短链接 API返回值
taobao.miniapp.shorturl.create

提供淘宝小程序短链接生成的能力，只允许对淘宝小程序对应的域名：https://m.duanqu.com/ 生成对应的短链接，其他域名无效
【特别注意：短链接有效期为30天，超过时效短链接将无法正常跳转到原始链接地址，请勿将短链接投放或装修到长期存在的入口】 */
type TaobaoMiniappShorturlCreateAPIResponse struct {
	model.CommonResponse
	TaobaoMiniappShorturlCreateAPIResponseModel
}

// TaobaoMiniappShorturlCreateAPIResponseModel is 生成淘宝小程序短链接 成功返回结果
type TaobaoMiniappShorturlCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"miniapp_shorturl_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoMiniappShorturlCreateResult `json:"result,omitempty" xml:"result,omitempty"`
}
