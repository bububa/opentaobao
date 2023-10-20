package tbitem

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallitemcalculatehscodegetAPIResponse 算法获取hscode API返回值
// tmall.item.calculate.hscode.get
//
// 算法获取hscode
type TmallitemcalculatehscodegetAPIResponse struct {
	model.CommonResponse
	TmallitemcalculatehscodegetAPIResponseModel
}

// TmallitemcalculatehscodegetAPIResponseModel is 算法获取hscode 成功返回结果
type TmallitemcalculatehscodegetAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_item_calculate_hscode_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 算法返回预测的hscode数据
	Results []string `json:"results,omitempty" xml:"results>string,omitempty"`
}
