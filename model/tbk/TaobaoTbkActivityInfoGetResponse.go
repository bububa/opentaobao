package tbk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
淘宝客-推广者-官方活动转链 APIResponse
taobao.tbk.activity.info.get

支持入参推广位和官方活动会场ID，获取活动信息和推广链接，包含推广长链接、短链接、淘口令、微信推广二维码地址等。改该API支持二方、三方类型的转链。官方活动会场ID，从淘宝客后台“我要推广-活动推广”中获取。
*/
type TaobaoTbkActivityInfoGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoTbkActivityInfoGetResponse `json:"taobao_tbk_activity_info_get_response,omitempty"`
}

type TaobaoTbkActivityInfoGetResponse struct {

    // 返回结果对象
    Data   *TaobaoTbkActivityInfoGetData `json:"data,omitempty"`

}
