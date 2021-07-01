package tbk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
淘宝客-推广者-淘礼金创建 API返回值 
taobao.tbk.dg.vegas.tlj.create

创建淘礼金
*/
type TaobaoTbkDgVegasTljCreateAPIResponse struct {
    model.CommonResponse
    TaobaoTbkDgVegasTljCreateAPIResponseModel
}

// 淘宝客-推广者-淘礼金创建 成功返回结果
type TaobaoTbkDgVegasTljCreateAPIResponseModel struct {
    XMLName xml.Name `xml:"tbk_dg_vegas_tlj_create_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *TaobaoTbkDgVegasTljCreateResult `json:"result,omitempty" xml:"result,omitempty"`
}
