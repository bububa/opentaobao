package pur

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabapurcmallgoodssyncAPIResponse 第三方商家接入采购商城-商品同步 API返回值
// alibaba.pur.cmall.goods.sync
//
// 第三方商家接入采购商城-商品同步
type AlibabapurcmallgoodssyncAPIResponse struct {
	model.CommonResponse
	AlibabapurcmallgoodssyncAPIResponseModel
}

// AlibabapurcmallgoodssyncAPIResponseModel is 第三方商家接入采购商城-商品同步 成功返回结果
type AlibabapurcmallgoodssyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_pur_cmall_goods_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象
	Result *ActionResult `json:"result,omitempty" xml:"result,omitempty"`
}
