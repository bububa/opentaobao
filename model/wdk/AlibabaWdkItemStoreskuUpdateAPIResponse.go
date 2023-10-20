package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkitemstoreskuupdateAPIResponse 五道口商品中心门店商品修改 API返回值
// alibaba.wdk.item.storesku.update
//
// 五道口商品中心门店商品修改
type AlibabawdkitemstoreskuupdateAPIResponse struct {
	model.CommonResponse
	AlibabawdkitemstoreskuupdateAPIResponseModel
}

// AlibabawdkitemstoreskuupdateAPIResponseModel is 五道口商品中心门店商品修改 成功返回结果
type AlibabawdkitemstoreskuupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_item_storesku_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabawdkitemstoreskuupdateResult `json:"result,omitempty" xml:"result,omitempty"`
}
