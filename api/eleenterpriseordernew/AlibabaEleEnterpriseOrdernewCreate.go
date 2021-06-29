package eleenterpriseordernew

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/eleenterpriseordernew"
)

/* 
创建订单 
alibaba.ele.enterprise.ordernew.create

创建订单
*/
func AlibabaEleEnterpriseOrdernewCreate(clt *core.SDKClient, req *eleenterpriseordernew.AlibabaEleEnterpriseOrdernewCreateRequest, session string) (*eleenterpriseordernew.AlibabaEleEnterpriseOrdernewCreateAPIResponse, error) {
    var resp eleenterpriseordernew.AlibabaEleEnterpriseOrdernewCreateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
