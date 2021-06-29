package eleenterpriseordernew

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/eleenterpriseordernew"
)

/* 
订单配送信息跟踪 
alibaba.ele.enterprise.ordernew.gettrackinginfo

订单配送信息跟踪
*/
func AlibabaEleEnterpriseOrdernewGettrackinginfo(clt *core.SDKClient, req *eleenterpriseordernew.AlibabaEleEnterpriseOrdernewGettrackinginfoRequest, session string) (*eleenterpriseordernew.AlibabaEleEnterpriseOrdernewGettrackinginfoAPIResponse, error) {
    var resp eleenterpriseordernew.AlibabaEleEnterpriseOrdernewGettrackinginfoAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
