package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoLogisticsCompaniesGetAPIRequest
查询物流公司信息 API请求
taobao.logistics.companies.get

查询淘宝网合作的物流公司信息，用于发货接口。 */
type TaobaoLogisticsCompaniesGetAPIRequest struct {
	model.Params
	// 需返回的字段列表。可选值:LogisticCompany 结构中的所有字段;多个字段间用","逗号隔开.<br/>如:id,code,name,reg_mail_no<br/><br><font color='red'>说明：</font><br/><br>id：物流公司ID<br/><br>code：物流公司code<br/><br>name：物流公司名称<br/><br>reg_mail_no：物流公司对应的运单规则
	_fields []string
	// 是否查询推荐物流公司.可选值:true,false.如果不提供此参数,将会返回所有支持电话联系的物流公司.
	_isRecommended bool
	// 推荐物流公司的下单方式.可选值:offline(电话联系/自己联系),online(在线下单),all(即电话联系又在线下单). 此参数仅仅用于is_recommended 为ture时。就是说对于推荐物流公司才可用.如果不选择此参数将会返回推荐物流中支持电话联系的物流公司.
	_orderMode string
}

// New
