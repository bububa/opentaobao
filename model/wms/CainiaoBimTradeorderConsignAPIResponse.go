package wms

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// CainiaobimtradeorderconsignAPIResponse 驱动保税交易订单发货 API返回值
// cainiao.bim.tradeorder.consign
//
// 驱动保税交易订单发货
type CainiaobimtradeorderconsignAPIResponse struct {
	model.CommonResponse
	CainiaobimtradeorderconsignAPIResponseModel
}

// CainiaobimtradeorderconsignAPIResponseModel is 驱动保税交易订单发货 成功返回结果
type CainiaobimtradeorderconsignAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_bim_tradeorder_consign_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 菜鸟仓库作业单据号
	StoreOrderCode string `json:"store_order_code,omitempty" xml:"store_order_code,omitempty"`
	// 菜鸟物流订单号,格式为LP开头
	LgOrderCode string `json:"lg_order_code,omitempty" xml:"lg_order_code,omitempty"`
}
