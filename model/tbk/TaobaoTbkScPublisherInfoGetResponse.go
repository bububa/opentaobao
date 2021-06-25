package tbk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
淘宝客-公用-私域用户备案信息查询 APIResponse
taobao.tbk.sc.publisher.info.get

查询已生成的渠道id或会员运营id的相关信息。
*/
type TaobaoTbkScPublisherInfoGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoTbkScPublisherInfoGetResponse `json:"taobao_tbk_sc_publisher_info_get_response,omitempty"`
}

type TaobaoTbkScPublisherInfoGetResponse struct {

    // data
    Data   *TaobaoTbkScPublisherInfoGetData `json:"data,omitempty"`

}
