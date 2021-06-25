package tbk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
淘宝客-公用-私域用户备案 APIResponse
taobao.tbk.sc.publisher.info.save

通过入参渠道管理或会员运营管理的邀请码，生成渠道id或会员运营id，完成渠道或会员的备案。
*/
type TaobaoTbkScPublisherInfoSaveAPIResponse struct {
    model.CommonResponse
    Response *TaobaoTbkScPublisherInfoSaveResponse `json:"taobao_tbk_sc_publisher_info_save_response,omitempty"`
}

type TaobaoTbkScPublisherInfoSaveResponse struct {

    // data
    Data   *TaobaoTbkScPublisherInfoSaveData `json:"data,omitempty"`

}
