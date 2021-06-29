package hotel

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
哈罗合作方获取酒店库存报价 API返回值 
taobao.xhotel.price.get.for.hello

哈罗合作项目，供哈罗合作方按需查询已开通城市下的标准酒店下指定时间段内的库存报价信息；在用户登录方面，返回结果不涉及用户个人信息，不涉及商家信息；仅根据不同用户，查询对应会员等级后，返回不同报价；
*/
type TaobaoXhotelPriceGetForHelloAPIResponse struct {
    model.CommonResponse
    TaobaoXhotelPriceGetForHelloResponse
}

// 哈罗合作方获取酒店库存报价 成功返回结果
type TaobaoXhotelPriceGetForHelloResponse struct {
    XMLName xml.Name `xml:"xhotel_price_get_for_hello_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 库价结果封装
    Result   *HotelPriceResultSet `json:"result,omitempty" xml:"result,omitempty"`
}
