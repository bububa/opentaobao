package pur

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaPurCmallGoodsStatusSyncAPIResponse 第三方商城接入采购商城-商品状态同步 API返回值
// alibaba.pur.cmall.goods.status.sync
//
// 第三方商城接入采购商城-商品状态同步
type AlibabaPurCmallGoodsStatusSyncAPIResponse struct {
	model.CommonResponse
	AlibabaPurCmallGoodsStatusSyncAPIResponseModel
}

// AlibabaPurCmallGoodsStatusSyncAPIResponseModel is 第三方商城接入采购商城-商品状态同步 成功返回结果
type AlibabaPurCmallGoodsStatusSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_pur_cmall_goods_status_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *ActionResult `json:"result,omitempty" xml:"result,omitempty"`
}
