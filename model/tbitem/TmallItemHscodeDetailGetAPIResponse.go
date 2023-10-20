package tbitem

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallitemhscodedetailgetAPIResponse 通过hscode获取计量单位 API返回值
// tmall.item.hscode.detail.get
//
// 通过hscode获取计量单位和销售单位
type TmallitemhscodedetailgetAPIResponse struct {
	model.CommonResponse
	TmallitemhscodedetailgetAPIResponseModel
}

// TmallitemhscodedetailgetAPIResponseModel is 通过hscode获取计量单位 成功返回结果
type TmallitemhscodedetailgetAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_item_hscode_detail_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回的计量单位和销售单位
	Results []string `json:"results,omitempty" xml:"results>string,omitempty"`
}
