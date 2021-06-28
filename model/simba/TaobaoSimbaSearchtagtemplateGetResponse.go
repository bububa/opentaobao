package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取搜索人群TOP用户可添加人群信息 APIResponse
taobao.simba.searchtagtemplate.get

获取搜索人群用户可添加人群信息
*/
type TaobaoSimbaSearchtagtemplateGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"simba_searchtagtemplate_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    TemplateList   []TaobaoSimbaSearchtagtemplateGetResult `json:"template_list,omitempty" xml:"