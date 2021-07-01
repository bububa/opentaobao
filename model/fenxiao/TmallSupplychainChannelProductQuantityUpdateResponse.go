package fenxiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
渠道无仓库存更新接口 API返回值 
tmall.supplychain.channel.product.quantity.update

渠道无仓库存更新接口
*/
type TmallSupplychainChannelProductQuantityUpdateAPIResponse struct {
    model.CommonResponse
    TmallSupplychainChannelProductQuantityUpdateResponse
}

// 渠道无仓库存更新接口 成功返回结果
type TmallSupplychainChannelProductQuantityUpdateResponse struct {
    XMLName xml.Name `xml:"tmall_supplychain_channel_product_quantity_update_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 异步获取历史数据接口返回结果
    Result   *ResultDTO `json:"result,omitempty" xml:"result,omitempty"`
}
