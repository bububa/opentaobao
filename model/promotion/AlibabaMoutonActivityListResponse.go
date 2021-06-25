package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
商家自运营活动列表 APIResponse
alibaba.mouton.activity.list

商家查询自己配置的活动列表
*/
type AlibabaMoutonActivityListAPIResponse struct {
    model.CommonResponse
    Response *AlibabaMoutonActivityListResponse `json:"alibaba_mouton_activity_list_response,omitempty"`
}

type AlibabaMoutonActivityListResponse struct {

    // 返回结果
    Result   *Page `json:"result,omitempty"`

}
