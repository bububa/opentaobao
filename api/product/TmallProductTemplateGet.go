package product

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/product"
)

/* 
产品接口 
tmall.product.template.get

产品模板获取接口，对于非关键属性的类目，发布达尔文(监管)产品时，必须先根据类目获取产品模板。<br/><br/>产品模板定义产品发布需要的类目属性，包括：<br/><br/>   关键属性:关键属性可以在类目上不存在。不存在的PID，默认为输入，没有子属性。属性名称在prop_name_str中取<br/>   绑定属性:内容为属性ID(PID)的列表,绑定属性肯定在类目上有，对应属性的类目特征，子属性请根据PID到类目上去取<br/><br/>   过滤属性:内容有属性ID(PID)列表，很重要的属性，filter_properties包含的属性，必须填写<br/><br/>   如果获取不到模板，非关键属性类目是不能发布产品的<br/>
*/
func TmallProductTemplateGet(clt *core.SDKClient, req *product.TmallProductTemplateGetAPIRequest, session string) (*product.TmallProductTemplateGetAPIResponse, error) {
    var resp product.TmallProductTemplateGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
