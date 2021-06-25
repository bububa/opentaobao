package shop

import (
    "github.com/bububa/opentaobao/model"
)

/* 
添加卖家自定义类目 APIResponse
taobao.sellercats.list.add

此API添加卖家店铺内自定义类目 <br/>父类目parent_cid值等于0：表示此类目为店铺下的一级类目，值不等于0：表示此类目有父类目 <br/>注：因为缓存的关系,添加的新类目需8个小时后才可以在淘宝页面上正常显示，但是不影响在该类目下商品发布
*/
type TaobaoSellercatsListAddAPIResponse struct {
    model.CommonResponse
    Response *TaobaoSellercatsListAddResponse `json:"taobao_sellercats_list_add_response,omitempty"`
}

type TaobaoSellercatsListAddResponse struct {

    // 返回seller_cat数据结构中的：cid,created
    SellerCat   *SellerCat `json:"seller_cat,omitempty"`

}
