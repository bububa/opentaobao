package beehive

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
分佣链接生成接口 APIResponse
taobao.beehive.item.cps.url

传入包括itemId,accountId,bizType在内的参数，对应参数返回分佣链接
*/
type TaobaoBeehiveItemCpsUrlAPIResponse struct {
    model.CommonResponse
    TaobaoBeehiveItemCpsUrlResponse
}

type TaobaoBeehiveItemCpsUrlResponse struct {
    XMLName xml.Name `xml:"beehive_item_cps_url_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 结果对象
    
    Result   *TaobaoBeehiveItemCpsUrlResultDo `json:"result,omitempty" xml:"result,omitempty"`

    
}
