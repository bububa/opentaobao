package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新增天猫商品尺码表模板 API请求
tmall.item.sizemapping.template.create

新增天猫商品尺码表模板<br/><br/>男鞋、女鞋、运动鞋、户外鞋类目，尺码表维度为：<br/>脚长（cm） 必选<br/><br/>内衣-文胸类目，尺码表维度为：<br/>上胸围（cm） 必选<br/>下胸围（cm） 必选<br/>上下胸围差（cm） 必选<br/>身高（cm）<br/>体重（公斤）<br/><br/>内衣-内裤类目，尺码表维度为：<br/>腰围（cm） 必选<br/>臀围（cm） 必选<br/>身高（cm）<br/>体重（公斤）<br/>裤长（cm）<br/>裆部（cm）<br/>脚口（cm）<br/>腿围（cm）<br/><br/>内衣-睡衣上衣/保暖上衣/睡袍类目，尺码表维度为：<br/>身高（cm） 必选<br/>胸围（cm） 必选<br/>体重（公斤）<br/>腰围（cm）<br/>肩宽（cm）<br/>袖长（cm）<br/>衣长（cm）<br/>背宽（cm）<br/>前长（cm）<br/>下摆围（cm）<br/>袖口（cm）<br/>袖肥（cm）<br/>领口（cm）<br/><br/>内衣-睡裤/保暖裤类目，尺码维度为：<br/>身高（cm） 必选<br/>腰围（cm） 必选<br/>体重（公斤）<br/>臀围（cm）<br/>裆部（cm）<br/>裤长（cm）<br/>脚口（cm）<br/>腿围（cm）<br/>裤侧长（cm）<br/><br/>内衣-睡裙类目，尺码维度为：<br/>身高（cm） 必选<br/>胸围（cm） 必选<br/>体重（公斤）<br/>裙长（cm）<br/>腰围（cm）<br/>袖长（cm）<br/>肩宽（cm）<br/>背宽（cm）<br/>腿围（cm）<br/>臀围（cm）<br/>底摆（cm）<br/><br/>男装、女装、运动服、户外服等上装类目，尺码维度为（至少两项必选）：<br/>身高（cm）<br/>体重（公斤）<br/>肩宽（cm）<br/>胸围（cm）<br/>腰围（cm）<br/>袖长（cm）<br/>衣长（cm）<br/>背宽（cm）<br/>前长（cm）<br/>摆围（cm）<br/>下摆围（cm）<br/>袖口（cm）<br/>袖肥（cm）<br/>中腰（cm）<br/>领深（cm）<br/>领高（cm）<br/>领宽（cm）<br/>领围（cm）<br/>圆摆后中长（cm）<br/>平摆衣长（cm）<br/>圆摆衣长（cm）<br/><br/>男装、女装、运动服、户外服等下装类目，尺码维度为（至少两项必选）：<br/>身高（cm）<br/>体重（公斤）<br/>腰围（cm）<br/>臀围（cm）<br/>裤长（cm）<br/>裙长（cm）<br/>裙摆长（cm）<br/>腿围（cm）<br/>膝围（cm）<br/>小脚围（cm）<br/>拉伸腰围（cm）<br/>坐围（cm）<br/>拉伸坐围（cm）<br/>脚口（cm）<br/>前浪（cm）<br/>后浪（cm）<br/>横档（cm）<br/><br/>如果上述维度满足，可以自定义最多5个维度。<br/><br/>模板格式为：<br/>尺码值:维度名称:数值<br/>如：M:身高（cm）:160,L:身高（cm）:170
*/
type TmallItemSizemappingTemplateCreateRequest struct {
    model.Params
    // 尺码表模板名称
    templateName   string
    // 尺码表模板内容，格式为"尺码值:维度名称:数值,尺码值:维度名称:数值"。其中，数值的单位，长度单位为厘米（cm），体重单位为公斤（kg）。尺码值，维度数据不能包含数字，特殊字符。数值为0-999.9的数字，且最多一位小数。
    templateContent   string
}

// 初始化TmallItemSizemappingTemplateCreateRequest对象
func NewTmallItemSizemappingTemplateCreateRequest() *TmallItemSizemappingTemplateCreateRequest{
    return &TmallItemSizemappingTemplateCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallItemSizemappingTemplateCreateRequest) GetApiMethodName() string {
    return "tmall.item.sizemapping.template.create"
}

// IRequest interface 方法, 获取API参数
func (r TmallItemSizemappingTemplateCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TemplateName Setter
// 尺码表模板名称
func (r *TmallItemSizemappingTemplateCreateRequest) SetTemplateName(templateName string) error {
    r.templateName = templateName
    r.Set("template_name", templateName)
    return nil
}

// TemplateName Getter
func (r TmallItemSizemappingTemplateCreateRequest) GetTemplateName() string {
    return r.templateName
}
// TemplateContent Setter
// 尺码表模板内容，格式为"尺码值:维度名称:数值,尺码值:维度名称:数值"。其中，数值的单位，长度单位为厘米（cm），体重单位为公斤（kg）。尺码值，维度数据不能包含数字，特殊字符。数值为0-999.9的数字，且最多一位小数。
func (r *TmallItemSizemappingTemplateCreateRequest) SetTemplateContent(templateContent string) error {
    r.templateContent = templateContent
    r.Set("template_content", templateContent)
    return nil
}

// TemplateContent Getter
func (r TmallItemSizemappingTemplateCreateRequest) GetTemplateContent() string {
    return r.templateContent
}
