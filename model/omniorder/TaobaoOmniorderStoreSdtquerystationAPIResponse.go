package omniorder

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOmniorderStoreSdtquerystationAPIResponse 速店通查询站点信息 API返回值
// taobao.omniorder.store.sdtquerystation
//
// 速店通查询站点信息
type TaobaoOmniorderStoreSdtquerystationAPIResponse struct {
	model.CommonResponse
	TaobaoOmniorderStoreSdtquerystationAPIResponseModel
}

// TaobaoOmniorderStoreSdtquerystationAPIResponseModel is 速店通查询站点信息 成功返回结果
type TaobaoOmniorderStoreSdtquerystationAPIResponseModel struct {
	XMLName xml.Name `xml:"omniorder_store_sdtquerystation_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoOmniorderStoreSdtquerystationResult `json:"result,omitempty" xml:"result,omitempty"`
}
