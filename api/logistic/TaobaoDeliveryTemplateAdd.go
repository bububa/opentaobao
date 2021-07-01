package logistic

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/logistic"
)

/* 
新增运费模板 
taobao.delivery.template.add

增加运费模板的外部接口
*/
func TaobaoDeliveryTemplateAdd(clt *core.SDKClient, req *logistic.TaobaoDeliveryTemplateAddAPIRequest, session string) (*logistic.TaobaoDeliveryTemplateAddAPIResponse, error) {
    var resp logistic.TaobaoDeliveryTemplateAddAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
