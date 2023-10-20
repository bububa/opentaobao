package tbitem

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaitemeditschemagetAPIResponse 商品编辑获取schema信息 API返回值
// alibaba.item.edit.schema.get
//
// 商品编辑时，获取商品规则信息
type AlibabaitemeditschemagetAPIResponse struct {
	model.CommonResponse
	AlibabaitemeditschemagetAPIResponseModel
}

// AlibabaitemeditschemagetAPIResponseModel is 商品编辑获取schema信息 成功返回结果
type AlibabaitemeditschemagetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_item_edit_schema_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 商品已有规则信息，XML格式.
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}
