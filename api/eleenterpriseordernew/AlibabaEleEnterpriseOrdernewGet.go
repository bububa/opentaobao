package eleenterpriseordernew

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/eleenterpriseordernew"
)

/* 
查询订单详情 
alibaba.ele.enterprise.ordernew.get

查询订单详情
*/
func AlibabaEleEnterpriseOrdernewGet(clt *core.SDKClient, req *eleenterpriseordernew.AlibabaEleEnterpriseOrdernewGetRequest, session string) (*eleenterpriseordernew.AlibabaEleEnterpriseOrdernewGetAPIResponse, error) {
    var resp eleenterpriseordernew.AlibabaEleEnterpriseOrdernewGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
