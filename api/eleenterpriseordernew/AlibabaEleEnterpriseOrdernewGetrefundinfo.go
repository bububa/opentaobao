package eleenterpriseordernew

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/eleenterpriseordernew"
)

/* 
退单和申诉 
alibaba.ele.enterprise.ordernew.getrefundinfo

退单和申诉
*/
func AlibabaEleEnterpriseOrdernewGetrefundinfo(clt *core.SDKClient, req *eleenterpriseordernew.AlibabaEleEnterpriseOrdernewGetrefundinfoAPIRequest, session string) (*eleenterpriseordernew.AlibabaEleEnterpriseOrdernewGetrefundinfoAPIResponse, error) {
    var resp eleenterpriseordernew.AlibabaEleEnterpriseOrdernewGetrefundinfoAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
