package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

/* TmallItemSizemappingTemplateCreate
新增天猫商品尺码表模板
tmall.item.sizemapping.template.create

新增天猫商品尺码表模板<br/><br/>男鞋、女鞋、运动鞋、户外鞋类目，尺码表维度为：<br/>脚长（cm） 必选<br/><br/>内衣-文胸类目，尺码表维度为：<br/>上胸围（cm） 必选<br/>下胸围（cm） 必选<br/>上下胸围差（cm） 必选<br/>身高（cm）<br/>体重（公斤）<br/><br/>内衣-内裤类目，尺码表维度为：<br/>腰围（cm） 必选<br/>臀围（cm） 必选<br/>身高（cm）<br/>体重（公斤）<br/>裤长（cm）<br/>裆部（cm）<br/>脚口（cm）<br/>腿围（cm）<br/><br/>内衣-睡衣上衣/保暖上衣/睡袍类目，尺码表维度为：<br/>身高（cm） 必选<br/>胸围（cm） 必选<br/>体重（公斤）<br/>腰围（cm）<br/>肩宽（cm）<br/>袖长（cm）<br/>衣长（cm）<br/>背宽（cm）<br/>前长（cm）<br/>下摆围（cm）<br/>袖口（cm）<br/>袖肥（cm）<br/>领口（cm）<br/><br/>内衣-睡裤/保暖裤类目，尺码维度为：<br/>身高（cm） 必选<br/>腰围（cm） 必选<br/>体重（公斤）<br/>臀围（cm）<br/>裆部（cm）<br/>裤长（cm）<br/>脚口（cm）<br/>腿围（cm）<br/>裤侧长（cm）<br/><br/>内衣-睡裙类目，尺码维度为：<br/>身高（cm） 必选<br/>胸围（cm） 必选<br/>体重（公斤）<br/>裙长（cm）<br/>腰围（cm）<br/>袖长（cm）<br/>肩宽（cm）<br/>背宽（cm）<br/>腿围（cm）<br/>臀围（cm）<br/>底摆（cm）<br/><br/>男装、女装、运动服、户外服等上装类目，尺码维度为（至少两项必选）：<br/>身高（cm）<br/>体重（公斤）<br/>肩宽（cm）<br/>胸围（cm）<br/>腰围（cm）<br/>袖长（cm）<br/>衣长（cm）<br/>背宽（cm）<br/>前长（cm）<br/>摆围（cm）<br/>下摆围（cm）<br/>袖口（cm）<br/>袖肥（cm）<br/>中腰（cm）<br/>领深（cm）<br/>领高（cm）<br/>领宽（cm）<br/>领围（cm）<br/>圆摆后中长（cm）<br/>平摆衣长（cm）<br/>圆摆衣长（cm）<br/><br/>男装、女装、运动服、户外服等下装类目，尺码维度为（至少两项必选）：<br/>身高（cm）<br/>体重（公斤）<br/>腰围（cm）<br/>臀围（cm）<br/>裤长（cm）<br/>裙长（cm）<br/>裙摆长（cm）<br/>腿围（cm）<br/>膝围（cm）<br/>小脚围（cm）<br/>拉伸腰围（cm）<br/>坐围（cm）<br/>拉伸坐围（cm）<br/>脚口（cm）<br/>前浪（cm）<br/>后浪（cm）<br/>横档（cm）<br/><br/>如果上述维度满足，可以自定义最多5个维度。<br/><br/>模板格式为：<br/>尺码值:维度名称:数值<br/>如：M:身高（cm）:160,L:身高（cm）:170 */
func TmallItemSizemappingTemplateCreate(clt *core.SDKClient, req *product.TmallItemSizemappingTemplateCreateAPIRequest, session string) (*product.TmallItemSizemappingTemplateCreateAPIResponse, error) {
	var resp product.TmallItemSizemappingTemplateCreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
