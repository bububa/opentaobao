package alsc

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alsc"
)

/* 
搜索卡实例列表(支持号段查询) 
alibaba.alsc.crm.card.searchcard

搜索卡实例列表(支持号段查询)
*/
func AlibabaAlscCrmCardSearchcard(clt *core.SDKClient, req *alsc.AlibabaAlscCrmCardSearchcardAPIRequest, session string) (*alsc.AlibabaAlscCrmCardSearchcardAPIResponse, error) {
    var resp alsc.AlibabaAlscCrmCardSearchcardAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
