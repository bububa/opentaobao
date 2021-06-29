package feedflow

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
信息流查看商品列表 APIResponse
taobao.feedflow.item.item.page

信息流查看商品列表
*/
type TaobaoFeedflowItemItemPageAPIResponse struct {
    model.CommonResponse
    TaobaoFeedflowItemItemPageResponse
}

type TaobaoFeedflowItemItemPageResponse struct {
    XMLName xml.Name `xml:"feedflow_item_item_page_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果对象
    
    Result   *TaobaoFeedflowItemItemPageResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
