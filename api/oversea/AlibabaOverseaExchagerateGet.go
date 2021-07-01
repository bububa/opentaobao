package oversea

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/oversea"
)

/* 
汇率信息获取 
alibaba.oversea.exchagerate.get

提供外部汇率查询接口
*/
func AlibabaOverseaExchagerateGet(clt *core.SDKClient, req *oversea.AlibabaOverseaExchagerateGetAPIRequest, session string) (*oversea.AlibabaOverseaExchagerateGetAPIResponse, error) {
    var resp oversea.AlibabaOverseaExchagerateGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
