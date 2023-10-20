package idle

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidletenderbtobitemqueryAPIResponse 暗拍b2b商品查询 API返回值
// alibaba.idle.tender.btob.item.query
//
// 暗拍b2b商品查询
type AlibabaidletenderbtobitemqueryAPIResponse struct {
	model.CommonResponse
	AlibabaidletenderbtobitemqueryAPIResponseModel
}

// AlibabaidletenderbtobitemqueryAPIResponseModel is 暗拍b2b商品查询 成功返回结果
type AlibabaidletenderbtobitemqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_tender_btob_item_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *CommonPageResult `json:"result,omitempty" xml:"result,omitempty"`
}
