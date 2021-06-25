package fenxiao

import (
    "github.com/bububa/opentaobao/model"
)

/* 
产品铺货状态查询 APIResponse
tmall.supplychain.channel.product.release.status.get

巴拿马战役渠道产品状态查询
*/
type TmallSupplychainChannelProductReleaseStatusGetAPIResponse struct {
    model.CommonResponse
    Response *TmallSupplychainChannelProductReleaseStatusGetResponse `json:"tmall_supplychain_channel_product_release_status_get_response,omitempty"`
}

type TmallSupplychainChannelProductReleaseStatusGetResponse struct {

    // 异步获取历史数据接口返回结果
    Result   *TmallSupplychainChannelProductReleaseStatusGetResultDto `json:"result,omitempty"`

}
