package qimen

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenPresalespackageConsignAPIResponse 预售预包尾款推单发货 API返回值
// taobao.qimen.presalespackage.consign
//
// 预售预包尾款推单发货
type TaobaoQimenPresalespackageConsignAPIResponse struct {
	model.CommonResponse
	TaobaoQimenPresalespackageConsignAPIResponseModel
}

// TaobaoQimenPresalespackageConsignAPIResponseModel is 预售预包尾款推单发货 成功返回结果
type TaobaoQimenPresalespackageConsignAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_presalespackage_consign_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应
	Response *PresalesPackageConsignResponse `json:"response,omitempty" xml:"response,omitempty"`
}
