package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// Tmallitemsizemappingtemplatecreate 新增天猫商品尺码表模板
// tmall.item.sizemapping.template.create
//
// 新增天猫商品尺码表模板&lt;br/&gt;&lt;br/&gt;男鞋、女鞋、运动鞋、户外鞋类目，尺码表维度为：&lt;br/&gt;脚长（cm） 必选&lt;br/&gt;&lt;br/&gt;内衣-文胸类目，尺码表维度为：&lt;br/&gt;上胸围（cm） 必选&lt;br/&gt;下胸围（cm） 必选&lt;br/&gt;上下胸围差（cm） 必选&lt;br/&gt;身高（cm）&lt;br/&gt;体重（公斤）&lt;br/&gt;&lt;br/&gt;内衣-内裤类目，尺码表维度为：&lt;br/&gt;腰围（cm） 必选&lt;br/&gt;臀围（cm） 必选&lt;br/&gt;身高（cm）&lt;br/&gt;体重（公斤）&lt;br/&gt;裤长（cm）&lt;br/&gt;裆部（cm）&lt;br/&gt;脚口（cm）&lt;br/&gt;腿围（cm）&lt;br/&gt;&lt;br/&gt;内衣-睡衣上衣/保暖上衣/睡袍类目，尺码表维度为：&lt;br/&gt;身高（cm） 必选&lt;br/&gt;胸围（cm） 必选&lt;br/&gt;体重（公斤）&lt;br/&gt;腰围（cm）&lt;br/&gt;肩宽（cm）&lt;br/&gt;袖长（cm）&lt;br/&gt;衣长（cm）&lt;br/&gt;背宽（cm）&lt;br/&gt;前长（cm）&lt;br/&gt;下摆围（cm）&lt;br/&gt;袖口（cm）&lt;br/&gt;袖肥（cm）&lt;br/&gt;领口（cm）&lt;br/&gt;&lt;br/&gt;内衣-睡裤/保暖裤类目，尺码维度为：&lt;br/&gt;身高（cm） 必选&lt;br/&gt;腰围（cm） 必选&lt;br/&gt;体重（公斤）&lt;br/&gt;臀围（cm）&lt;br/&gt;裆部（cm）&lt;br/&gt;裤长（cm）&lt;br/&gt;脚口（cm）&lt;br/&gt;腿围（cm）&lt;br/&gt;裤侧长（cm）&lt;br/&gt;&lt;br/&gt;内衣-睡裙类目，尺码维度为：&lt;br/&gt;身高（cm） 必选&lt;br/&gt;胸围（cm） 必选&lt;br/&gt;体重（公斤）&lt;br/&gt;裙长（cm）&lt;br/&gt;腰围（cm）&lt;br/&gt;袖长（cm）&lt;br/&gt;肩宽（cm）&lt;br/&gt;背宽（cm）&lt;br/&gt;腿围（cm）&lt;br/&gt;臀围（cm）&lt;br/&gt;底摆（cm）&lt;br/&gt;&lt;br/&gt;男装、女装、运动服、户外服等上装类目，尺码维度为（至少两项必选）：&lt;br/&gt;身高（cm）&lt;br/&gt;体重（公斤）&lt;br/&gt;肩宽（cm）&lt;br/&gt;胸围（cm）&lt;br/&gt;腰围（cm）&lt;br/&gt;袖长（cm）&lt;br/&gt;衣长（cm）&lt;br/&gt;背宽（cm）&lt;br/&gt;前长（cm）&lt;br/&gt;摆围（cm）&lt;br/&gt;下摆围（cm）&lt;br/&gt;袖口（cm）&lt;br/&gt;袖肥（cm）&lt;br/&gt;中腰（cm）&lt;br/&gt;领深（cm）&lt;br/&gt;领高（cm）&lt;br/&gt;领宽（cm）&lt;br/&gt;领围（cm）&lt;br/&gt;圆摆后中长（cm）&lt;br/&gt;平摆衣长（cm）&lt;br/&gt;圆摆衣长（cm）&lt;br/&gt;&lt;br/&gt;男装、女装、运动服、户外服等下装类目，尺码维度为（至少两项必选）：&lt;br/&gt;身高（cm）&lt;br/&gt;体重（公斤）&lt;br/&gt;腰围（cm）&lt;br/&gt;臀围（cm）&lt;br/&gt;裤长（cm）&lt;br/&gt;裙长（cm）&lt;br/&gt;裙摆长（cm）&lt;br/&gt;腿围（cm）&lt;br/&gt;膝围（cm）&lt;br/&gt;小脚围（cm）&lt;br/&gt;拉伸腰围（cm）&lt;br/&gt;坐围（cm）&lt;br/&gt;拉伸坐围（cm）&lt;br/&gt;脚口（cm）&lt;br/&gt;前浪（cm）&lt;br/&gt;后浪（cm）&lt;br/&gt;横档（cm）&lt;br/&gt;&lt;br/&gt;如果上述维度满足，可以自定义最多5个维度。&lt;br/&gt;&lt;br/&gt;模板格式为：&lt;br/&gt;尺码值:维度名称:数值&lt;br/&gt;如：M:身高（cm）:160,L:身高（cm）:170
func Tmallitemsizemappingtemplatecreate(clt *core.SDKClient, req *product.TmallitemsizemappingtemplatecreateAPIRequest, session string) (*product.TmallitemsizemappingtemplatecreateAPIResponse, error) {
	var resp product.TmallitemsizemappingtemplatecreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
