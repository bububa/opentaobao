package idle

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidletenderbtobitemdeleteAPIResponse 暗拍b2b商品下架/删除 API返回值
// alibaba.idle.tender.btob.item.delete
//
// 暗拍b2b商品下架/删除
type AlibabaidletenderbtobitemdeleteAPIResponse struct {
	model.CommonResponse
	AlibabaidletenderbtobitemdeleteAPIResponseModel
}

// AlibabaidletenderbtobitemdeleteAPIResponseModel is 暗拍b2b商品下架/删除 成功返回结果
type AlibabaidletenderbtobitemdeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_tender_btob_item_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}
