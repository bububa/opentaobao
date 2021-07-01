package alicom

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alicom"
)

/* 
查询指定商家的可用的话费宝贝接口 
alibaba.chongzhi.queryecards

查询指定商家的可用的话费宝贝
*/
func AlibabaChongzhiQueryecards(clt *core.SDKClient, req *alicom.AlibabaChongzhiQueryecardsAPIRequest, session string) (*alicom.AlibabaChongzhiQueryecardsAPIResponse, error) {
    var resp alicom.AlibabaChongzhiQueryecardsAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
