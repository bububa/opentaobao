package alsc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取特价菜单 APIResponse
alibaba.alsc.crm.menu.list

获取特价菜单
*/
type AlibabaAlscCrmMenuListAPIResponse struct {
    model.CommonResponse
    Response *AlibabaAlscCrmMenuListResponse `json:"alibaba_alsc_crm_menu_list_response,omitempty"`
}

type AlibabaAlscCrmMenuListResponse struct {

    // 分页返回模型
    Result   *CommonPageResult `json:"result,omitempty"`

}
