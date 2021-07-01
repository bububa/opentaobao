package baichuan

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaBaichuanTaopasswordConfigAPIResponse
淘口令配置数据 API返回值
alibaba.baichuan.taopassword.config

百川淘口令规则配置接口 */
type AlibabaBaichuanTaopasswordConfigAPIResponse struct {
	model.CommonResponse
	AlibabaBaichuanTaopasswordConfigAPIResponseModel
}

// AlibabaBaichuanTaopasswordConfigAPIResponseModel is 淘口令配置数据 成功返回结果
type AlibabaBaichuanTaopasswordConfigAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_baichuan_taopassword_config_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ShareResult `json:"result,omitempty" xml:"result,omitempty"`
}
