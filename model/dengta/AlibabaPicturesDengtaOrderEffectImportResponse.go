package dengta

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
天下秀订单数据导入 APIResponse
alibaba.pictures.dengta.order.effect.import

提供接口给天下秀，天下秀订单数据效果生成时回流到灯塔系统
*/
type AlibabaPicturesDengtaOrderEffectImportAPIResponse struct {
    model.CommonResponse
    AlibabaPicturesDengtaOrderEffectImportResponse
}

type AlibabaPicturesDengtaOrderEffectImportResponse struct {
    XMLName xml.Name `xml:"alibaba_pictures_dengta_order_effect_import_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 结果
    
    Result   *ApiGeneralResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
