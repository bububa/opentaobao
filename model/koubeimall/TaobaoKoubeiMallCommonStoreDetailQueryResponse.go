package koubeimall

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询综合体内的门店详细信息 APIResponse
taobao.koubei.mall.common.store.detail.query

查询口碑综合体内的门店详情信息
*/
type TaobaoKoubeiMallCommonStoreDetailQueryAPIResponse struct {
    model.CommonResponse
    TaobaoKoubeiMallCommonStoreDetailQueryResponse
}

type TaobaoKoubeiMallCommonStoreDetailQueryResponse struct {
    XMLName xml.Name `xml:"koubei_mall_common_store_detail_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *TaobaoKoubeiMallCommonStoreDetailQueryResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
