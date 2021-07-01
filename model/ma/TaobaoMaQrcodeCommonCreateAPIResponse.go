package ma

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoMaQrcodeCommonCreateAPIResponse
创建码平台常用二维码 API返回值
taobao.ma.qrcode.common.create

创建码平台对外提供的常用二维码接口，不适于码平台业务类型的码创建，如不支持包裹码、媒体码等，业务类型的码需要单独提供API。 */
type TaobaoMaQrcodeCommonCreateAPIResponse struct {
	model.CommonResponse
	TaobaoMaQrcodeCommonCreateAPIResponseModel
}

// TaobaoMaQrcodeCommonCreateAPIResponseModel is 创建码平台常用二维码 成功返回结果
type TaobaoMaQrcodeCommonCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"ma_qrcode_common_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 二维码对像
	Modules []QrcodeDo `json:"modules,omitempty" xml:"modules>qrcode_do,omitempty"`
	// 执行是否成功
	Suc bool `json:"suc,omitempty" xml:"suc,omitempty"`
}
