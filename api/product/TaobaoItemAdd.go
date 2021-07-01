package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

/* TaobaoItemAdd
添加一个商品
taobao.item.add

此接口用于新增一个淘宝商品
商品的属性和sku的属性有包含的关系，商品的价格要位于sku的价格区间之中（例如，sku价格有5元、10元两种，那么商品的价格就需要大于等于5元，小于等于10元，否则新增商品会失败）
商品的类目和商品的价格、sku的价格都有一定的相关性（具体的关系要通过类目属性查询接口获得）
商品的运费承担方式和邮费设置有相关性，卖家承担运费不用设置邮费，买家承担运费需要设置邮费
当关键属性值选择了“其他”的时候，需要输入input_pids和input_str商品才能添加成功。
<br/><strong><a href="https://console.open.taobao.com/dingWeb.htm?from=itemapi" target="_blank">点击查看更多商品API说明</a></strong> */
func TaobaoItemAdd(clt *core.SDKClient, req *product.TaobaoItemAddAPIRequest, session string) (*product.TaobaoItemAddAPIResponse, error) {
	var resp product.TaobaoItemAddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
