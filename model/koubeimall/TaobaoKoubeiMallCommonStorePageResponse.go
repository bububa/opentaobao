package koubeimall

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
分页查询综合体内的门店列表信息 APIResponse
taobao.koubei.mall.common.store.page

分页查询综合体内的门店列表信息
*/
type TaobaoKoubeiMallCommonStorePageAPIResponse struct {
    model.CommonResponse
    TaobaoKoubeiMallCommonStorePageResponse
}

type TaobaoKoubeiMallCommonStorePageResponse struct {
    XMLName xml.Name `xml:"koubei_mall_common_store_page_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // API接口返回的result模型
    
    Result   *TaobaoKoubeiMallCommonStorePageResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
