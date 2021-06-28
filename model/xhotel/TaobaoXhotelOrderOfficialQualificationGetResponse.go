package xhotel

import (
    "github.com/bububa/opentaobao/model"
)

/* 
官网信用住用户资质校验 APIResponse
taobao.xhotel.order.official.qualification.get

官网信用住在下单前对用户进行资质校验，资质校验通过才能进行信用支付
*/
type TaobaoXhotelOrderOfficialQualificationGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoXhotelOrderOfficialQualificationGetResponse `json:"xhotel_order_official_qualification_get_response,omitempty"` 
    TaobaoXhotelOrderOfficialQualificationGetResponse
}

/* model for simplify = false
type TaobaoXhotelOrderOfficialQualificationGetResponse struct {

    // 阿里旅行中间调用URL
    
    TransferUrl   string `json:"transfer_url,omitempty"`
    

    // 无资格原因提示信息
    
    Reason   string `json:"reason,omitempty"`
    

    // 资质校验是否成功,有资格返回true, 无资格返回false
    
    MatchCondition   bool `json:"match_condition,omitempty"`
    

    // 入参信息回传, 用于校验的证件号码
    
    IdNumber   string `json:"id_number,omitempty"`
    

    // 入参信息回传，用于校验的外部会员账号
    
    OutMemeberAccount   string `json:"out_memeber_account,omitempty"`
    

    // * 外部请求序列表号回传\流水号（如果外部订单号唯一，建议外部订单号，这样容易核对问题），用于区别一次请求和排查问题，阿里旅行不做任何格式校验(必须)
    
    OutUUID   string `json:"out_u_u_i_d,omitempty"`
    

    // 扩展字段，json串，后续留用
    
    ExtendAttrs   string `json:"extend_attrs,omitempty"`
    

}
*/

type TaobaoXhotelOrderOfficialQualificationGetResponse struct {

    // 阿里旅行中间调用URL
    TransferUrl   string `json:"transfer_url,omitempty"`

    // 无资格原因提示信息
    Reason   string `json:"reason,omitempty"`

    // 资质校验是否成功,有资格返回true, 无资格返回false
    MatchCondition   bool `json:"match_condition,omitempty"`

    // 入参信息回传, 用于校验的证件号码
    IdNumber   string `json:"id_number,omitempty"`

    // 入参信息回传，用于校验的外部会员账号
    OutMemeberAccount   string `json:"out_memeber_account,omitempty"`

    // * 外部请求序列表号回传\流水号（如果外部订单号唯一，建议外部订单号，这样容易核对问题），用于区别一次请求和排查问题，阿里旅行不做任何格式校验(必须)
    OutUUID   string `json:"out_u_u_i_d,omitempty"`

    // 扩展字段，json串，后续留用
    ExtendAttrs   string `json:"extend_attrs,omitempty"`

}
