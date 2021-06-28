package tanx

import (
    "github.com/bububa/opentaobao/model"
)

/* 
提交资质接口 APIResponse
taobao.tanx.qualification.add

dsp客户提交客户资质和行业资质
*/
type TaobaoTanxQualificationAddAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoTanxQualificationAddResponse `json:"tanx_qualification_add_response,omitempty"` 
    TaobaoTanxQualificationAddResponse
}

/* model for simplify = false
type TaobaoTanxQualificationAddResponse struct {

    // 是否成功
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

    // 返回id对应列表
    
    IdList  struct {
        Number  []int64 `json:"int64,omitempty"`
    } `json:"id_list,omitempty"`
    

}
*/

type TaobaoTanxQualificationAddResponse struct {

    // 是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

    // 返回id对应列表
    IdList   []int64 `json:"id_list,omitempty"`

}
