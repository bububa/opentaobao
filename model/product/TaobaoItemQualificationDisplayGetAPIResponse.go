package product

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoItemQualificationDisplayGetAPIResponse 资质采集配置异步获取接口 API返回值
// taobao.item.qualification.display.get
//
// 根据类目，商品，属性等参与动态获得资质采集配置
type TaobaoItemQualificationDisplayGetAPIResponse struct {
	model.CommonResponse
	TaobaoItemQualificationDisplayGetAPIResponseModel
}

// TaobaoItemQualificationDisplayGetAPIResponseModel is 资质采集配置异步获取接口 成功返回结果
type TaobaoItemQualificationDisplayGetAPIResponseModel struct {
	XMLName xml.Name `xml:"item_qualification_display_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回资质采集配置
	DisplayConf *DisplayQualifications `json:"display_conf,omitempty" xml:"display_conf,omitempty"`
}
