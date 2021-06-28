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
    TaobaoSimbaSearchtagtemplateGetResponse
}

type TaobaoSimbaSearchtagtemplateGetResponse struct {
    XMLName xml.Name `xml:"simba_searchtagtemplate_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    TemplateList   []TaobaoSimbaSearchtagtemplateGetResult `json:"template_list,omitempty" xml:"template_list>taobao_simba_searchtagtemplate_get_result,omitempty"`
    
    
}
