package tbk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
淘宝客-推广者-超级红包领取状态查询 APIResponse
taobao.tbk.dg.vegas.send.status

淘宝客传入用户标识的信息，查询该用户是否有当前阶段待核销的红包（淘客接入前需签署协议 https://pub.alimama.com/fourth/protocol/common.htm?key=hangye_laxin）
*/
type TaobaoTbkDgVegasSendStatusAPIResponse struct {
    model.CommonResponse
    TaobaoTbkDgVegasSendStatusResponse
}

type TaobaoTbkDgVegasSendStatusResponse struct {
    XMLName xml.Name `xml:"tbk_dg_vegas_send_status_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回结果描述信息
    
    ResultMsg   string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`

    
    // 返回结果封装对象
    
    Data   *TaobaoTbkDgVegasSendStatusData `json:"data,omitempty" xml:"data,omitempty"`

    
}
