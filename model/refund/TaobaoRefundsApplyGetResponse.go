package refund

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询买家申请的退款列表 APIResponse
taobao.refunds.apply.get

查询买家申请的退款列表，且查询外店的退款列表时需要指定交易类型
*/
type TaobaoRefundsApplyGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"refunds_apply_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 搜索到的退款信息列表
    
    Refunds   []Refund `json:"refunds,omitempty" xml:"