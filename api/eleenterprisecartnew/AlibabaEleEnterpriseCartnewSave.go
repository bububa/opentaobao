package eleenterprisecartnew

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/eleenterprisecartnew"
)

/* 
新版创建购物车 
alibaba.ele.enterprise.cartnew.save

新版创建购物车
*/
func AlibabaEleEnterpriseCartnewSave(clt *core.SDKClient, req *eleenterprisecartnew.AlibabaEleEnterpriseCartnewSaveRequest, session string) (*eleenterprisecartnew.AlibabaEleEnterpriseCartnewSaveAPIResponse, error) {
    var resp eleenterprisecartnew.AlibabaEleEnterpriseCartnewSaveAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
