package koubeimall

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
分页查询已授权的商圈列表信息 APIResponse
taobao.koubei.mall.common.mall.auth.page

分页查询口碑已授权商圈的列表信息
*/
type TaobaoKoubeiMallCommonMallAuthPageAPIResponse struct {
    model.CommonResponse
    TaobaoKoubeiMallCommonMallAuthPageResponse
}

type TaobaoKoubeiMallCommonMallAuthPageResponse struct {
    XMLName xml.Name `xml:"koubei_mall_common_mall_auth_page_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // API接口返回的result模型
    
    Result   *TaobaoKoubeiMallCommonMallAuthPageResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
