package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoDeliveryTemplateUpdateAPIRequest 修改运费模板 API请求
// taobao.delivery.template.update
//
// 修改运费模板
type TaobaoDeliveryTemplateUpdateAPIRequest struct {
	model.Params
	// 模板名称，长度不能大于50个字节
	_name string
	// 需要修改的模板对应的模板ID
	_templateId int64
	// 可选值：0,1 <br>  说明<br>0:表示买家承担服务费;<br>1:表示卖家承担服务费
	_assumer int64
	// 运费方式:平邮 (post),快递公司(express),EMS (ems),货到付款(cod)结构:value1;value2;value3;value4 <br/>如: post;express;ems;cod <br/><br/><font color=red><br/>注意:在添加多个运费方式时,字符串中使用 ";" 分号区分。template_dests(指定地区) template_start_standards(首费标准)、template_start_fees(首费)、template_add_standards(增费标准)、template_add_fees(增费)必须与template_types的分号数量相同. <br/> </font><br/><br/><br/><font color=blue><br/>普通用户：post,ems,express三种运费方式必须填写一个，不能填写cod。<br/>货到付款用户：如果填写了cod运费方式，则post,ems,express三种运费方式也必须填写一个，如果没有填写cod则填写的运费方式中必须存在express</font>
	_templateTypes string
	// 邮费子项涉及的地区.结构: value1;value2;value3,value4<br>如:1,110000;1,110000;1,310000;1,320000,330000。 aredId解释(1=全国,110000=北京,310000=上海,320000=江苏,330000=浙江)如果template_types设置为post;ems;exmpress;cod则表示为平邮(post)指定默认地区(全国)和北京地区的运费;其他的类似以分号区分一一对应<br/>可以用taobao.areas.get接口获取.或者参考：http://www.stats.gov.cn/tjbz/xzqhdm/t20080215_402462675.htm<br/><font color=red>每个运费方式设置的设涉及地区中必须包含全国地区（areaId=1）表示默认运费,可以只设置默认运费</font><br><font color=blue>注意:为多个地区指定指定不同（首费标准、首费、增费标准、增费一项不一样就算不同）的运费以逗号","区分，template_start_standards(首费标准)、template_start_fees(首费)、template_add_standards(增费标准)、template_add_fees(增费)必须与template_types分号数量相同。如果为需要为多个地区指定相同运费则地区之间用“|”隔开即可。</font>
	_templateDests string
	// 首费标准：当valuation(记价方式)为0时输入1-9999范围内的整数<br><font color=blue>首费标准目前只能为1</br><br><font color=red>输入的格式分号个数和template_types数量一致，逗号个数必须与template_dests以分号隔开之后一一对应的数量一致</font>
	_templateStartStandards string
	// 首费：输入0.01-999.99（最多包含两位小数）<br/><br/><font color=blue> 首费不能为0</font><br><font color=red>输入的格式分号个数和template_types数量一致，逗号个数必须与template_dests以分号隔开之后一一对应的数量一致</font>
	_templateStartFees string
	// 增费标准：当valuation(记价方式)为0时输入1-9999范围内的整数<br><font color=blue>增费标准目前只能为1</font><br><font color=red>输入的格式分号个数和template_types数量一致，逗号个数必须与template_dests以分号隔开之后一一对应的数量一致</font>
	_templateAddStandards string
	// 增费：输入0.00-999.99（最多包含两位小数）<br/><font color=blue>增费可以为0</font><br/><font color=red>输入的格式分号个数和template_types数量一致，逗号个数必须与template_dests以分号隔开之后一一对应的数量一致</font>
	_templateAddFees string
}

// NewTaobaoDeliveryTemplateUpdateRequest 初始化TaobaoDeliveryTemplateUpdateAPIRequest对象
func NewTaobaoDeliveryTemplateUpdateRequest() *TaobaoDeliveryTemplateUpdateAPIRequest {
	return &TaobaoDeliveryTemplateUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoDeliveryTemplateUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.delivery.template.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoDeliveryTemplateUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Name Setter
// 模板名称，长度不能大于50个字节
func (r *TaobaoDeliveryTemplateUpdateAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// Get Name Getter
func (r TaobaoDeliveryTemplateUpdateAPIRequest) GetName() string {
	return r._name
}

// Set is TemplateId Setter
// 需要修改的模板对应的模板ID
func (r *TaobaoDeliveryTemplateUpdateAPIRequest) SetTemplateId(_templateId int64) error {
	r._templateId = _templateId
	r.Set("template_id", _templateId)
	return nil
}

// Get TemplateId Getter
func (r TaobaoDeliveryTemplateUpdateAPIRequest) GetTemplateId() int64 {
	return r._templateId
}

// Set is Assumer Setter
// 可选值：0,1 <br>  说明<br>0:表示买家承担服务费;<br>1:表示卖家承担服务费
func (r *TaobaoDeliveryTemplateUpdateAPIRequest) SetAssumer(_assumer int64) error {
	r._assumer = _assumer
	r.Set("assumer", _assumer)
	return nil
}

// Get Assumer Getter
func (r TaobaoDeliveryTemplateUpdateAPIRequest) GetAssumer() int64 {
	return r._assumer
}

// Set is TemplateTypes Setter
// 运费方式:平邮 (post),快递公司(express),EMS (ems),货到付款(cod)结构:value1;value2;value3;value4 <br/>如: post;express;ems;cod <br/><br/><font color=red><br/>注意:在添加多个运费方式时,字符串中使用 ";" 分号区分。template_dests(指定地区) template_start_standards(首费标准)、template_start_fees(首费)、template_add_standards(增费标准)、template_add_fees(增费)必须与template_types的分号数量相同. <br/> </font><br/><br/><br/><font color=blue><br/>普通用户：post,ems,express三种运费方式必须填写一个，不能填写cod。<br/>货到付款用户：如果填写了cod运费方式，则post,ems,express三种运费方式也必须填写一个，如果没有填写cod则填写的运费方式中必须存在express</font>
func (r *TaobaoDeliveryTemplateUpdateAPIRequest) SetTemplateTypes(_templateTypes string) error {
	r._templateTypes = _templateTypes
	r.Set("template_types", _templateTypes)
	return nil
}

// Get TemplateTypes Getter
func (r TaobaoDeliveryTemplateUpdateAPIRequest) GetTemplateTypes() string {
	return r._templateTypes
}

// Set is TemplateDests Setter
// 邮费子项涉及的地区.结构: value1;value2;value3,value4<br>如:1,110000;1,110000;1,310000;1,320000,330000。 aredId解释(1=全国,110000=北京,310000=上海,320000=江苏,330000=浙江)如果template_types设置为post;ems;exmpress;cod则表示为平邮(post)指定默认地区(全国)和北京地区的运费;其他的类似以分号区分一一对应<br/>可以用taobao.areas.get接口获取.或者参考：http://www.stats.gov.cn/tjbz/xzqhdm/t20080215_402462675.htm<br/><font color=red>每个运费方式设置的设涉及地区中必须包含全国地区（areaId=1）表示默认运费,可以只设置默认运费</font><br><font color=blue>注意:为多个地区指定指定不同（首费标准、首费、增费标准、增费一项不一样就算不同）的运费以逗号","区分，template_start_standards(首费标准)、template_start_fees(首费)、template_add_standards(增费标准)、template_add_fees(增费)必须与template_types分号数量相同。如果为需要为多个地区指定相同运费则地区之间用“|”隔开即可。</font>
func (r *TaobaoDeliveryTemplateUpdateAPIRequest) SetTemplateDests(_templateDests string) error {
	r._templateDests = _templateDests
	r.Set("template_dests", _templateDests)
	return nil
}

// Get TemplateDests Getter
func (r TaobaoDeliveryTemplateUpdateAPIRequest) GetTemplateDests() string {
	return r._templateDests
}

// Set is TemplateStartStandards Setter
// 首费标准：当valuation(记价方式)为0时输入1-9999范围内的整数<br><font color=blue>首费标准目前只能为1</br><br><font color=red>输入的格式分号个数和template_types数量一致，逗号个数必须与template_dests以分号隔开之后一一对应的数量一致</font>
func (r *TaobaoDeliveryTemplateUpdateAPIRequest) SetTemplateStartStandards(_templateStartStandards string) error {
	r._templateStartStandards = _templateStartStandards
	r.Set("template_start_standards", _templateStartStandards)
	return nil
}

// Get TemplateStartStandards Getter
func (r TaobaoDeliveryTemplateUpdateAPIRequest) GetTemplateStartStandards() string {
	return r._templateStartStandards
}

// Set is TemplateStartFees Setter
// 首费：输入0.01-999.99（最多包含两位小数）<br/><br/><font color=blue> 首费不能为0</font><br><font color=red>输入的格式分号个数和template_types数量一致，逗号个数必须与template_dests以分号隔开之后一一对应的数量一致</font>
func (r *TaobaoDeliveryTemplateUpdateAPIRequest) SetTemplateStartFees(_templateStartFees string) error {
	r._templateStartFees = _templateStartFees
	r.Set("template_start_fees", _templateStartFees)
	return nil
}

// Get TemplateStartFees Getter
func (r TaobaoDeliveryTemplateUpdateAPIRequest) GetTemplateStartFees() string {
	return r._templateStartFees
}

// Set is TemplateAddStandards Setter
// 增费标准：当valuation(记价方式)为0时输入1-9999范围内的整数<br><font color=blue>增费标准目前只能为1</font><br><font color=red>输入的格式分号个数和template_types数量一致，逗号个数必须与template_dests以分号隔开之后一一对应的数量一致</font>
func (r *TaobaoDeliveryTemplateUpdateAPIRequest) SetTemplateAddStandards(_templateAddStandards string) error {
	r._templateAddStandards = _templateAddStandards
	r.Set("template_add_standards", _templateAddStandards)
	return nil
}

// Get TemplateAddStandards Getter
func (r TaobaoDeliveryTemplateUpdateAPIRequest) GetTemplateAddStandards() string {
	return r._templateAddStandards
}

// Set is TemplateAddFees Setter
// 增费：输入0.00-999.99（最多包含两位小数）<br/><font color=blue>增费可以为0</font><br/><font color=red>输入的格式分号个数和template_types数量一致，逗号个数必须与template_dests以分号隔开之后一一对应的数量一致</font>
func (r *TaobaoDeliveryTemplateUpdateAPIRequest) SetTemplateAddFees(_templateAddFees string) error {
	r._templateAddFees = _templateAddFees
	r.Set("template_add_fees", _templateAddFees)
	return nil
}

// Get TemplateAddFees Getter
func (r TaobaoDeliveryTemplateUpdateAPIRequest) GetTemplateAddFees() string {
	return r._templateAddFees
}
