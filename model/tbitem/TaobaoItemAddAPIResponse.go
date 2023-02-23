package tbitem

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoItemAddAPIResponse 添加一个商品 API返回值
// taobao.item.add
//
// 此接口用于新增一个淘宝商品
// 商品的属性和sku的属性有包含的关系，商品的价格要位于sku的价格区间之中（例如，sku价格有5元、10元两种，那么商品的价格就需要大于等于5元，小于等于10元，否则新增商品会失败）
// 商品的类目和商品的价格、sku的价格都有一定的相关性（具体的关系要通过类目属性查询接口获得）
// 商品的运费承担方式和邮费设置有相关性，卖家承担运费不用设置邮费，买家承担运费需要设置邮费
// 当关键属性值选择了“其他”的时候，需要输入input_pids和input_str商品才能添加成功。
// &lt;br/&gt;&lt;strong&gt;&lt;a href=&#34;https://console.open.taobao.com/dingWeb.htm?from=itemapi&#34; target=&#34;_blank&#34;&gt;点击查看更多商品API说明&lt;/a&gt;&lt;/strong&gt;
type TaobaoItemAddAPIResponse struct {
	model.CommonResponse
	TaobaoItemAddAPIResponseModel
}

// TaobaoItemAddAPIResponseModel is 添加一个商品 成功返回结果
type TaobaoItemAddAPIResponseModel struct {
	XMLName xml.Name `xml:"item_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 商品结构,仅有numIid和created返回
	Item *Item `json:"item,omitempty" xml:"item,omitempty"`
}
