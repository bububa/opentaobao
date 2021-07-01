package tbk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
淘宝客-推广者-新用户订单明细查询 API返回值 
taobao.tbk.dg.newuser.order.get

拉新API
*/
type TaobaoTbkDgNewuserOrderGetAPIResponse struct {
    model.CommonResponse
    TaobaoTbkDgNewuserOrderGetAPIResponseModel
}

// 淘宝客-推广者-新用户订单明细查询 成功返回结果
type TaobaoTbkDgNewuserOrderGetAPIResponseModel struct {
    XMLName xml.Name `xml:"tbk_dg_newuser_order_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // data
    Results   *TaobaoTbkDgNewuserOrderGetResults `json:"results,omitempty" xml:"results,omitempty"`
}
