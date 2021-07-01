package fundplatform

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/fundplatform"
)

/* 
通知制卡商制卡 
alibaba.fundplatform.cardorder.make

该接口由内部定义，外部制卡商实现。当需要制卡商进行制卡操作时，将会调用该接口。
*/
func AlibabaFundplatformCardorderMake(clt *core.SDKClient, req *fundplatform.AlibabaFundplatformCardorderMakeAPIRequest, session string) (*fundplatform.AlibabaFundplatformCardorderMakeAPIResponse, error) {
    var resp fundplatform.AlibabaFundplatformCardorderMakeAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
