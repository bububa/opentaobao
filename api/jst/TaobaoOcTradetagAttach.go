package jst

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// Taobaooctradetagattach 订单打标或者订单标签更新
// taobao.oc.tradetag.attach
//
// 对订单添加标签和更新标签。标签分为官方标签和自定义标签。&lt;br/&gt;官方标签目前有:赠品,电子发票,收货地址变更,预售。具体格式说明请看http://open.taobao.com/doc/detail.htm?id=102731&lt;br/&gt;自定义标签有2个通用属性:&lt;br/&gt;    `show_str:给消费者显示的字符串（如果可以显示的话）&lt;br/&gt;    `pic_urls:图片url,地址必须是图片空间的url,最多5张
func Taobaooctradetagattach(clt *core.SDKClient, req *jst.TaobaooctradetagattachAPIRequest, session string) (*jst.TaobaooctradetagattachAPIResponse, error) {
	var resp jst.TaobaooctradetagattachAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
