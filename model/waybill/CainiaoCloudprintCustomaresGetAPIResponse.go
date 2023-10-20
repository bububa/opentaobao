package waybill

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// CainiaocloudprintcustomaresgetAPIResponse 获取商家的自定义区模板信息 API返回值
// cainiao.cloudprint.customares.get
//
// 供isv使用，获取商家的自定义区的模板信息
type CainiaocloudprintcustomaresgetAPIResponse struct {
	model.CommonResponse
	CainiaocloudprintcustomaresgetAPIResponseModel
}

// CainiaocloudprintcustomaresgetAPIResponseModel is 获取商家的自定义区模板信息 成功返回结果
type CainiaocloudprintcustomaresgetAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_cloudprint_customares_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *CloudPrintBaseResult `json:"result,omitempty" xml:"result,omitempty"`
}
