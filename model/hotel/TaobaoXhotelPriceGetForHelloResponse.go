package hotel

import (
    "github.com/bububa/opentaobao/model"
)

/* 
哈罗合作方获取酒店库存报价 APIResponse
taobao.xhotel.price.get.for.hello

哈罗合作项目，供哈罗合作方按需查询已开通城市下的标准酒店下指定时间段内的库存报价信息；在用户登录方面，返回结果不涉及用户个人信息，不涉及商家信息；仅根据不同用户，查询对应会员等级后，返回不同报价；
*/
type TaobaoXhotelPriceGetForHelloAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoXhotelPriceGetForHelloResponse `json:"xhotel_price_get_for_hello_response,omitempty"` 
    TaobaoXhotelPriceGetForHelloResponse
}

/* model for simplify = false
type TaobaoXhotelPriceGetForHelloResponse struct {

    // 库价结果封装
    
    Result  *struct {
        HotelPriceResultSet  *HotelPriceResultSet `json:"hotel_price_result_set,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoXhotelPriceGetForHelloResponse struct {

    // 库价结果封装
    Result   *HotelPriceResultSet `json:"result,omitempty"`

}
