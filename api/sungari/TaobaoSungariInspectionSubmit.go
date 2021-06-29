package sungari

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/sungari"
)

/* 
抽检指令录入 
taobao.sungari.inspection.submit

抽检指令录入
*/
func TaobaoSungariInspectionSubmit(clt *core.SDKClient, req *sungari.TaobaoSungariInspectionSubmitRequest, session string) (*sungari.TaobaoSungariInspectionSubmitAPIResponse, error) {
    var resp sungari.TaobaoSungariInspectionSubmitAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
