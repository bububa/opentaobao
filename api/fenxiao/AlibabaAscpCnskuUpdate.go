package fenxiao

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/fenxiao"
)

/* 
供应链中台货品修改接口 
alibaba.ascp.cnsku.update

供应链中台货品修改接口
*/
func AlibabaAscpCnskuUpdate(clt *core.SDKClient, req *fenxiao.AlibabaAscpCnskuUpdateAPIRequest, session string) (*fenxiao.AlibabaAscpCnskuUpdateAPIResponse, error) {
    var resp fenxiao.AlibabaAscpCnskuUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
