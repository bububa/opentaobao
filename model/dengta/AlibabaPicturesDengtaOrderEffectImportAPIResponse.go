package dengta

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaPicturesDengtaOrderEffectImportAPIResponse
天下秀订单数据导入 API返回值
alibaba.pictures.dengta.order.effect.import

提供接口给天下秀，天下秀订单数据效果生成时回流到灯塔系统 */
type AlibabaPicturesDengtaOrderEffectImportAPIResponse struct {
	model.CommonResponse
	AlibabaPicturesDengtaOrderEffectImportAPIResponseModel
}

// AlibabaPicturesDengtaOrderEffectImportAPIResponseModel is 天下秀订单数据导入 成功返回结果
type AlibabaPicturesDengtaOrderEffectImportAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_pictures_dengta_order_effect_import_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *ApiGeneralResult `json:"result,omitempty" xml:"result,omitempty"`
}
