package tbk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkDgVegasTljTemporaryCreateAPIResponse 淘宝客-推广者-淘礼金创建测试营销ID API返回值
// taobao.tbk.dg.vegas.tlj.temporary.create
//
// 淘宝客-推广者-淘礼金创建测试营销ID
type TaobaoTbkDgVegasTljTemporaryCreateAPIResponse struct {
	model.CommonResponse
	TaobaoTbkDgVegasTljTemporaryCreateAPIResponseModel
}

// TaobaoTbkDgVegasTljTemporaryCreateAPIResponseModel is 淘宝客-推广者-淘礼金创建测试营销ID 成功返回结果
type TaobaoTbkDgVegasTljTemporaryCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"tbk_dg_vegas_tlj_temporary_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoTbkDgVegasTljTemporaryCreateResult `json:"result,omitempty" xml:"result,omitempty"`
}
