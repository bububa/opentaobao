package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoDeliveryTemplatesGetAPIRequest
获取用户下所有模板 API请求
taobao.delivery.templates.get

根据用户ID获取用户下所有模板 */
type TaobaoDeliveryTemplatesGetAPIRequest struct {
	model.Params
	// 需返回的字段列表。 <br/> <br/><B><br/>可选值:示例中的值;各字段之间用","隔开<br/></B><br/><br/><br/> <br/><font color=blue><br/>template_id：查询模板ID <br/> <br/>template_name:查询模板名称 <br/> <br/>assumer：查询服务承担放 <br/> <br/>valuation:查询计价规则 <br/> <br/>supports:查询增值服务列表 <br/> <br/>created:查询模板创建时间 <br/> <br/>modified:查询修改时间<br/> <br/>consign_area_id:运费模板上设置的卖家发货地址最小级别ID<br/> <br/>address:卖家地址信息<br/></font><br/><br/><br/> <br/><font color=bule><br/>query_cod:查询货到付款运费信息<br/> <br/>query_post:查询平邮运费信息 <br/> <br/>query_express:查询快递公司运费信息 <br/> <br/>query_ems:查询EMS运费信息<br/> <br/>query_bzsd:查询普通保障速递运费信息<br/> <br/>query_wlb:查询物流宝保障速递运费信息<br/> <br/>query_furniture:查询家装物流运费信息<br/><font color=blue><br/><br/><br/><br/><font color=red>不能有空格</font>
	_fields string
}

// New
