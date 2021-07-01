package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取一个产品的信息 API返回值 
taobao.product.get

天猫商家发布商品时，查询关联产品信息时使用，非商品查询接口。商品查询接口：taobao.item.seller.get<br>
两种方式查看一个产品详细信息: 
传入product_id来查询；传入cid和props来查询
<br/><strong><a href="https://console.open.taobao.com/dingWeb.htm?from=itemapi" target="_blank">点击查看更多商品API说明</a></strong>
*/
type TaobaoProductGetAPIResponse struct {
    model.CommonResponse
    TaobaoProductGetAPIResponseModel
}

// 获取一个产品的信息 成功返回结果
type TaobaoProductGetAPIResponseModel struct {
    XMLName xml.Name `xml:"product_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回具体信息为入参fields请求的字段信息
    Product   *Product `json:"product,omitempty" xml:"product,omitempty"`
}
