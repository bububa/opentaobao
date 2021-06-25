package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取搜索人群TOP用户可添加人群信息 APIResponse
taobao.simba.searchtagtemplate.get

获取搜索人群用户可添加人群信息
*/
type TaobaoSimbaSearchtagtemplateGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoSimbaSearchtagtemplateGetResponse `json:"taobao_simba_searchtagtemplate_get_response,omitempty"`
}

type TaobaoSimbaSearchtagtemplateGetResponse struct {

    // result
    TemplateList   []TaobaoSimbaSearchtagtemplateGetResult `json:"template_list,omitempty"`

}
