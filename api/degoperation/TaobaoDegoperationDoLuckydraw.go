package degoperation

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/degoperation"
)

/* 
激励抽奖 
taobao.degoperation.do.luckydraw

激励平台抽奖接口。用户可以通过接口完成抽奖功能
*/
func TaobaoDegoperationDoLuckydraw(clt *core.SDKClient, req *degoperation.TaobaoDegoperationDoLuckydrawAPIRequest, session string) (*degoperation.TaobaoDegoperationDoLuckydrawAPIResponse, error) {
    var resp degoperation.TaobaoDegoperationDoLuckydrawAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
