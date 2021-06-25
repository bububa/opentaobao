package caipiao

import (
    "github.com/bububa/opentaobao/model"
)

/* 
根据卖家id与appkey获取商品信息 APIResponse
taobao.caipiao.goods.info.get

根据卖家id与appkey获取商品信息。
*/
type TaobaoCaipiaoGoodsInfoGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoCaipiaoGoodsInfoGetResponse `json:"taobao_caipiao_goods_info_get_response,omitempty"`
}

type TaobaoCaipiaoGoodsInfoGetResponse struct {

    // 返回列表的大小
    TotalResults   int64 `json:"total_results,omitempty"`

    // 查询的结果列表
    Results   []LotteryWangcaiSellerGoodsInfo `json:"results,omitempty"`

}
