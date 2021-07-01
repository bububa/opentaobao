package eleenterprisecartnew

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/eleenterprisecartnew"
)

/* 
新版购物车查询 
alibaba.ele.enterprise.cartnew.query

新版购物车查询
*/
func AlibabaEleEnterpriseCartnewQuery(clt *core.SDKClient, req *eleenterprisecartnew.AlibabaEleEnterpriseCartnewQueryAPIRequest, session string) (*eleenterprisecartnew.AlibabaEleEnterpriseCartnewQueryAPIResponse, error) {
    var resp eleenterprisecartnew.AlibabaEleEnterpriseCartnewQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
