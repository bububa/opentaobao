package servicecenter

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
内购服务确认单明细上传接口 APIResponse
taobao.fuwu.sp.billreord.add

isv能通过该接口上传确认单明细数据
*/
type TaobaoFuwuSpBillreordAddAPIResponse struct {
    model.CommonResponse
    TaobaoFuwuSpBillreordAddResponse
}

type TaobaoFuwuSpBillreordAddResponse struct {
    XMLName xml.Name `xml:"fuwu_sp_billreord_add_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回调用结果
    
    AddResult   bool `json:"add_result,omitempty" xml:"add_result,omitempty"`

    
}
