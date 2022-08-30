package ascp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangItemBatchUpdateAsyncAPIResponse 货品新建/更新接口 API返回值
// alibaba.dchain.aoxiang.item.batch.update.async
//
// 货品新建/更新接口
type AlibabaDchainAoxiangItemBatchUpdateAsyncAPIResponse struct {
	model.CommonResponse
	AlibabaDchainAoxiangItemBatchUpdateAsyncAPIResponseModel
}

// AlibabaDchainAoxiangItemBatchUpdateAsyncAPIResponseModel is 货品新建/更新接口 成功返回结果
type AlibabaDchainAoxiangItemBatchUpdateAsyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_dchain_aoxiang_item_batch_update_async_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求出参
	ItemUpdateResponse *ItemUpdateAsnycResponse `json:"item_update_response,omitempty" xml:"item_update_response,omitempty"`
}
