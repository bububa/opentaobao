package fenxiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
渠道无仓库存更新接口 APIResponse
tmall.supplychain.channel.product.quantity.update

渠道无仓库存更新接口
*/
type TmallSupplychainChannelProductQuantityUpdateAPIResponse struct {
    model.CommonResponse
    TmallSupplychainChannelProductQuantityUpdateResponse
}

type TmallSupplychainChannelProductQuantityUpdateResponse struct {
    XMLName xml.Name `xml:"tmall_supplychain_channel_product_quantity_update_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 异步获取历史数据接口返回结果
    
    Result   *TmallSupplychainChannelProductQuantityUpdateResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
