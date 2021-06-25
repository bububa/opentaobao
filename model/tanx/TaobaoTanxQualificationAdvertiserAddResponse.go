package tanx

import (
    "github.com/bububa/opentaobao/model"
)

/* 
新增广告主接口 APIResponse
taobao.tanx.qualification.advertiser.add

外部dsp调用接口时会根据广告主名称和类型在tanx系统中新增一个广告主
*/
type TaobaoTanxQualificationAdvertiserAddAPIResponse struct {
    model.CommonResponse
    Response *TaobaoTanxQualificationAdvertiserAddResponse `json:"taobao_tanx_qualification_advertiser_add_response,omitempty"`
}

type TaobaoTanxQualificationAdvertiserAddResponse struct {

    // 调用是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

    // 返回的广告主dto对象
    AdvertiserList   []AdvertiserDto `json:"advertiser_list,omitempty"`

}
