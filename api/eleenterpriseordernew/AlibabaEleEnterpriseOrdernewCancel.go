package eleenterpriseordernew

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/eleenterpriseordernew"
)

/* 
订单取消 
alibaba.ele.enterprise.ordernew.cancel

订单取消
*/
func AlibabaEleEnterpriseOrdernewCancel(clt *core.SDKClient, req *eleenterpriseordernew.AlibabaEleEnterpriseOrdernewCancelRequest, session string) (*eleenterpriseordernew.AlibabaEleEnterpriseOrdernewCancelAPIResponse, error) {
    var resp eleenterpriseordernew.AlibabaEleEnterpriseOrdernewCancelAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
