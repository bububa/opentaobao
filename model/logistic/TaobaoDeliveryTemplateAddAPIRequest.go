package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoDeliveryTemplateAddAPIRequest
新增运费模板 API请求
taobao.delivery.template.add

增加运费模板的外部接口 */
type TaobaoDeliveryTemplateAddAPIRequest struct {
	model.Params
	// 运费模板的名称，长度不能超过50个字节
	_name string
	// 可选值：0、1 ，说明如下<br>0:表示买家承担服务费;<br>1:表示卖家承担服务费
	_assumer int64
	// 可选值：0、1、3，说明如下。<br>0:表示按宝贝件数计算运费 <br>1:表示按宝贝重量计算运费 <br>3:表示按宝贝体积计算运费
	_valuation int64
	// 运费方式:平邮 (post),快递公司(express),EMS (ems),货到付款(cod),物流宝保障速递(wlb),家装物流(furniture)结构:value1;value2;value3;value4 如: post;express;ems;cod <br/><font color=red>注意:在添加多个运费方式时,字符串中使用 ";" 分号区分。template_dests(指定地区)template_start_standards(首费标准)、template_start_fees(首费)、template_add_standards(增费标准)、template_add_fees(增费)必须与template_types的分号数量相同. </font><br><font color=blue>注意：<br/>1、post,ems,express三种运费方式必须填写一个<br/>2、只有订购了货到付款用户可以填cod;只有家装物流用户可以填写furniture只有订购了保障速递的用户可以填写bzsd,只有物流宝用户可以填写wlb<br/>3、如果是货到付款用户当没有填写cod运送方式时，运费模板会默认继承express的费用为cod的费用<br/>4、如果是保障速递户当没有填写bzsd运送方式时，运费模板会默认继承express的费用为bzsd的费用<br/>5、由于3和4的原因所以当是货到付款用户或保障速递用户时如果没填写对应的运送方式express是必须填写的</font>
	_templateTypes string
	// 邮费子项涉及的地区.结构: value1;value2;value3,value4<br>如:1,110000;1,110000;1,310000;1,320000,330000。 aredId解释(1=全国,110000=北京,310000=上海,320000=江苏,330000=浙江)如果template_types设置为post;ems;exmpress;cod则表示为平邮(post)指定默认地区(全国)和北京地区的运费;其他的类似以分号区分一一对应<br/>可以用taobao.areas.get接口获取.或者参考：http://www.stats.gov.cn/tjbz/xzqhdm/t20080215_402462675.htm<br/><br/><font color=red>每个运费方式设置涉及的地区中必须包含全国地区（areaId=1）表示默认运费,可以只设置默认运费</font><br><font color=blue>注意:为多个地区指定指定不同（首费标准、首费、增费标准、增费一项不一样就算不同）的运费以逗号","区分，template_start_standards(首费标准)、template_start_fees(首费)、template_add_standards(增费标准)、template_add_fees(增费)必须与template_types分号数量相同。如果为需要为多个地区指定相同运费则地区之间用“|”隔开即可。</font><font color=red>如果卖家没有传入发货地址则默认地址库的发货地址</font>
	_templateDests string
	// 首费标准：当valuation(记价方式)为0时输入1-9999范围内的整数<br/><font color=red>当valuation(记价方式)为1时输入0.1-9999.9范围内的小数只能包含以为小数（单位为千克）<br/><font color=blue>当valuation(记价方式)为3时输入0.1-999.9范围内的数值，数值只能包含一位小数（单位为 立方米）<br/><font color=red>输入的格式分号个数和template_types数量一致，逗号个数必须与template_dests以分号隔开之后一一对应的数量一致</font>
	_templateStartStandards string
	// 首费：输入0.00-999.99（最多包含两位小数）<br/><font color=red>输入的格式分号个数和template_types数量一致，逗号个数必须与template_dests以分号隔开之后一一对应的数量一致</font>
	_templateStartFees string
	// 增费标准：当valuation(记价方式)为0时输入1-9999范围内的整数<br/><font color=red>当valuation(记价方式)为1时输入0.1-9999.9范围内的小数只能包含以为小数（单位为千克）<br/><font color=blue>当valuation(记价方式)为3时输入0.1-999.9范围内的数值，数值只能包含一位小数（单位为 立方米）<br/><br><font color=red>输入的格式分号个数和template_types数量一致，逗号个数必须与template_dests以分号隔开之后一一对应的数量一致</font>
	_templateAddStandards string
	// 增费：输入0.00-999.99（最多包含两位小数）<br/><br/><br/><br/><font color=blue>增费必须小于等于首费，但是当首费为0时增费可以大于首费</font><br/><br/><br/><br/><br/><font color=red>输入的格式分号个数和template_types数量一致，逗号个数必须与template_dests以分号隔开之后一一对应的数量一致</font>
	_templateAddFees string
	// 卖家发货地址区域ID<br/><br/><br/><font color=blue>可以不填，如果没有填写取卖家默认发货地址的区域ID，如果需要输入必须用taobao.areas.get接口获取.或者参考：http://www.stats.gov.cn/tjbz/xzqhdm/t20080215_402462675.htm <br/></font><br/><br/><br/><br/><font color=red>注意：填入该值时必须取您的发货地址最小区域级别ID，比如您的发货地址是：某省XX市xx区（县）时需要填入区(县)的ID，当然有些地方没有区或县可以直接填市级别的ID</font>
	_consignAreaId int64
}

// NewTaobaoDeliveryTemplateAddRequest 初始化TaobaoDeliveryTemplateAddAPIRequest对象
func NewTaobaoDeliveryTemplateAddRequest() *TaobaoDeliveryTemplateAddAPIRequest {
	return &TaobaoDeliveryTemplateAddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoDeliveryTemplateAddAPIRequest) GetApiMethodName() string {
	return "taobao.delivery.template.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoDeliveryTemplateAddAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Name Setter
// 运费模板的名称，长度不能超过50个字节
func (r *TaobaoDeliveryTemplateAddAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// Get Name Getter
func (r TaobaoDeliveryTemplateAddAPIRequest) GetName() string {
	return r._name
}

// Set is Assumer Setter
// 可选值：0、1 ，说明如下<br>0:表示买家承担服务费;<br>1:表示卖家承担服务费
func (r *TaobaoDeliveryTemplateAddAPIRequest) SetAssumer(_assumer int64) error {
	r._assumer = _assumer
	r.Set("assumer", _assumer)
	return nil
}

// Get Assumer Getter
func (r TaobaoDeliveryTemplateAddAPIRequest) GetAssumer() int64 {
	return r._assumer
}

// Set is Valuation Setter
// 可选值：0、1、3，说明如下。<br>0:表示按宝贝件数计算运费 <br>1:表示按宝贝重量计算运费 <br>3:表示按宝贝体积计算运费
func (r *TaobaoDeliveryTemplateAddAPIRequest) SetValuation(_valuation int64) error {
	r._valuation = _valuation
	r.Set("valuation", _valuation)
	return nil
}

// Get Valuation Getter
func (r TaobaoDeliveryTemplateAddAPIRequest) GetValuation() int64 {
	return r._valuation
}

// Set is TemplateTypes Setter
// 运费方式:平邮 (post),快递公司(express),EMS (ems),货到付款(cod),物流宝保障速递(wlb),家装物流(furniture)结构:value1;value2;value3;value4 如: post;express;ems;cod <br/><font color=red>注意:在添加多个运费方式时,字符串中使用 ";" 分号区分。template_dests(指定地区)template_start_standards(首费标准)、template_start_fees(首费)、template_add_standards(增费标准)、template_add_fees(增费)必须与template_types的分号数量相同. </font><br><font color=blue>注意：<br/>1、post,ems,express三种运费方式必须填写一个<br/>2、只有订购了货到付款用户可以填cod;只有家装物流用户可以填写furniture只有订购了保障速递的用户可以填写bzsd,只有物流宝用户可以填写wlb<br/>3、如果是货到付款用户当没有填写cod运送方式时，运费模板会默认继承express的费用为cod的费用<br/>4、如果是保障速递户当没有填写bzsd运送方式时，运费模板会默认继承express的费用为bzsd的费用<br/>5、由于3和4的原因所以当是货到付款用户或保障速递用户时如果没填写对应的运送方式express是必须填写的</font>
func (r *TaobaoDeliveryTemplateAddAPIRequest) SetTemplateTypes(_templateTypes string) error {
	r._templateTypes = _templateTypes
	r.Set("template_types", _templateTypes)
	return nil
}

// Get TemplateTypes Getter
func (r TaobaoDeliveryTemplateAddAPIRequest) GetTemplateTypes() string {
	return r._templateTypes
}

// Set is TemplateDests Setter
// 邮费子项涉及的地区.结构: value1;value2;value3,value4<br>如:1,110000;1,110000;1,310000;1,320000,330000。 aredId解释(1=全国,110000=北京,310000=上海,320000=江苏,330000=浙江)如果template_types设置为post;ems;exmpress;cod则表示为平邮(post)指定默认地区(全国)和北京地区的运费;其他的类似以分号区分一一对应<br/>可以用taobao.areas.get接口获取.或者参考：http://www.stats.gov.cn/tjbz/xzqhdm/t20080215_402462675.htm<br/><br/><font color=red>每个运费方式设置涉及的地区中必须包含全国地区（areaId=1）表示默认运费,可以只设置默认运费</font><br><font color=blue>注意:为多个地区指定指定不同（首费标准、首费、增费标准、增费一项不一样就算不同）的运费以逗号","区分，template_start_standards(首费标准)、template_start_fees(首费)、template_add_standards(增费标准)、template_add_fees(增费)必须与template_types分号数量相同。如果为需要为多个地区指定相同运费则地区之间用“|”隔开即可。</font><font color=red>如果卖家没有传入发货地址则默认地址库的发货地址</font>
func (r *TaobaoDeliveryTemplateAddAPIRequest) SetTemplateDests(_templateDests string) error {
	r._templateDests = _templateDests
	r.Set("template_dests", _templateDests)
	return nil
}

// Get TemplateDests Getter
func (r TaobaoDeliveryTemplateAddAPIRequest) GetTemplateDests() string {
	return r._templateDests
}

// Set is TemplateStartStandards Setter
// 首费标准：当valuation(记价方式)为0时输入1-9999范围内的整数<br/><font color=red>当valuation(记价方式)为1时输入0.1-9999.9范围内的小数只能包含以为小数（单位为千克）<br/><font color=blue>当valuation(记价方式)为3时输入0.1-999.9范围内的数值，数值只能包含一位小数（单位为 立方米）<br/><font color=red>输入的格式分号个数和template_types数量一致，逗号个数必须与template_dests以分号隔开之后一一对应的数量一致</font>
func (r *TaobaoDeliveryTemplateAddAPIRequest) SetTemplateStartStandards(_templateStartStandards string) error {
	r._templateStartStandards = _templateStartStandards
	r.Set("template_start_standards", _templateStartStandards)
	return nil
}

// Get TemplateStartStandards Getter
func (r TaobaoDeliveryTemplateAddAPIRequest) GetTemplateStartStandards() string {
	return r._templateStartStandards
}

// Set is TemplateStartFees Setter
// 首费：输入0.00-999.99（最多包含两位小数）<br/><font color=red>输入的格式分号个数和template_types数量一致，逗号个数必须与template_dests以分号隔开之后一一对应的数量一致</font>
func (r *TaobaoDeliveryTemplateAddAPIRequest) SetTemplateStartFees(_templateStartFees string) error {
	r._templateStartFees = _templateStartFees
	r.Set("template_start_fees", _templateStartFees)
	return nil
}

// Get TemplateStartFees Getter
func (r TaobaoDeliveryTemplateAddAPIRequest) GetTemplateStartFees() string {
	return r._templateStartFees
}

// Set is TemplateAddStandards Setter
// 增费标准：当valuation(记价方式)为0时输入1-9999范围内的整数<br/><font color=red>当valuation(记价方式)为1时输入0.1-9999.9范围内的小数只能包含以为小数（单位为千克）<br/><font color=blue>当valuation(记价方式)为3时输入0.1-999.9范围内的数值，数值只能包含一位小数（单位为 立方米）<br/><br><font color=red>输入的格式分号个数和template_types数量一致，逗号个数必须与template_dests以分号隔开之后一一对应的数量一致</font>
func (r *TaobaoDeliveryTemplateAddAPIRequest) SetTemplateAddStandards(_templateAddStandards string) error {
	r._templateAddStandards = _templateAddStandards
	r.Set("template_add_standards", _templateAddStandards)
	return nil
}

// Get TemplateAddStandards Getter
func (r TaobaoDeliveryTemplateAddAPIRequest) GetTemplateAddStandards() string {
	return r._templateAddStandards
}

// Set is TemplateAddFees Setter
// 增费：输入0.00-999.99（最多包含两位小数）<br/><br/><br/><br/><font color=blue>增费必须小于等于首费，但是当首费为0时增费可以大于首费</font><br/><br/><br/><br/><br/><font color=red>输入的格式分号个数和template_types数量一致，逗号个数必须与template_dests以分号隔开之后一一对应的数量一致</font>
func (r *TaobaoDeliveryTemplateAddAPIRequest) SetTemplateAddFees(_templateAddFees string) error {
	r._templateAddFees = _templateAddFees
	r.Set("template_add_fees", _templateAddFees)
	return nil
}

// Get TemplateAddFees Getter
func (r TaobaoDeliveryTemplateAddAPIRequest) GetTemplateAddFees() string {
	return r._templateAddFees
}

// Set is ConsignAreaId Setter
// 卖家发货地址区域ID<br/><br/><br/><font color=blue>可以不填，如果没有填写取卖家默认发货地址的区域ID，如果需要输入必须用taobao.areas.get接口获取.或者参考：http://www.stats.gov.cn/tjbz/xzqhdm/t20080215_402462675.htm <br/></font><br/><br/><br/><br/><font color=red>注意：填入该值时必须取您的发货地址最小区域级别ID，比如您的发货地址是：某省XX市xx区（县）时需要填入区(县)的ID，当然有些地方没有区或县可以直接填市级别的ID</font>
func (r *TaobaoDeliveryTemplateAddAPIRequest) SetConsignAreaId(_consignAreaId int64) error {
	r._consignAreaId = _consignAreaId
	r.Set("consign_area_id", _consignAreaId)
	return nil
}

// Get ConsignAreaId Getter
func (r TaobaoDeliveryTemplateAddAPIRequest) GetConsignAreaId() int64 {
	return r._consignAreaId
}
