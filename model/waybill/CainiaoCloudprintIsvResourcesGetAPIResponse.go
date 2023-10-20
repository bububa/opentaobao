package waybill

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// CainiaocloudprintisvresourcesgetAPIResponse isv资源查询 API返回值
// cainiao.cloudprint.isv.resources.get
//
// isv资源查询，包括isv模板、打印项、预设的自定义区等
type CainiaocloudprintisvresourcesgetAPIResponse struct {
	model.CommonResponse
	CainiaocloudprintisvresourcesgetAPIResponseModel
}

// CainiaocloudprintisvresourcesgetAPIResponseModel is isv资源查询 成功返回结果
type CainiaocloudprintisvresourcesgetAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_cloudprint_isv_resources_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *CloudPrintBaseResult `json:"result,omitempty" xml:"result,omitempty"`
}
