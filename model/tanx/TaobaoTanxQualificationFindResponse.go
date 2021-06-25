package tanx

import (
    "github.com/bububa/opentaobao/model"
)

/* 
资质查询接口 APIResponse
taobao.tanx.qualification.find

资质查询接口
*/
type TaobaoTanxQualificationFindAPIResponse struct {
    model.CommonResponse
    Response *TaobaoTanxQualificationFindResponse `json:"taobao_tanx_qualification_find_response,omitempty"`
}

type TaobaoTanxQualificationFindResponse struct {

    // 是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

    // 返回的资质内容dto
    QualificationList   []QualificationDto `json:"qualification_list,omitempty"`

    // 查询返回总条数
    Count   string `json:"count,omitempty"`

}
