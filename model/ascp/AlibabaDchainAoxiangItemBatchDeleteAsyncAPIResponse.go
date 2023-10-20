package ascp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabadchainaoxiangitembatchdeleteasyncAPIResponse 货品与组合货品删除 API返回值
// alibaba.dchain.aoxiang.item.batch.delete.async
//
// 货品与组合货品删除
type AlibabadchainaoxiangitembatchdeleteasyncAPIResponse struct {
	model.CommonResponse
	AlibabadchainaoxiangitembatchdeleteasyncAPIResponseModel
}

// AlibabadchainaoxiangitembatchdeleteasyncAPIResponseModel is 货品与组合货品删除 成功返回结果
type AlibabadchainaoxiangitembatchdeleteasyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_dchain_aoxiang_item_batch_delete_async_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求出参
	ItemDeleteResponse *ItemDeleteAsyncResponse `json:"item_delete_response,omitempty" xml:"item_delete_response,omitempty"`
}
