package jst

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/jst"
)

/* 
查询服务商下的商家是否开通了改地址 
alibaba.modifyaddress.isv.bindseller.check

鉴权服务商和商家的绑定关系，看商家是否开通了改地址
1. 没有授权
2. 已与当前appkey签约
3. 没有签约
4. 与其他服务商软件签约，如果是同一个isv name，返回appkey，否则不返回。
*/
func AlibabaModifyaddressIsvBindsellerCheck(clt *core.SDKClient, req *jst.AlibabaModifyaddressIsvBindsellerCheckRequest, session string) (*jst.AlibabaModifyaddressIsvBindsellerCheckResponse, error) {
    var resp jst.AlibabaModifyaddressIsvBindsellerCheckAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
