package ascpchannel

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascppurchasepricecreateAPIResponse ascp采购价写入接口 API返回值
// alibaba.ascp.purchase.price.create
//
// 供应链平台采购价创建或修改接口
type AlibabaascppurchasepricecreateAPIResponse struct {
	model.CommonResponse
	AlibabaascppurchasepricecreateAPIResponseModel
}

// AlibabaascppurchasepricecreateAPIResponseModel is ascp采购价写入接口 成功返回结果
type AlibabaascppurchasepricecreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_purchase_price_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *ResultWrapper `json:"result,omitempty" xml:"result,omitempty"`
}
