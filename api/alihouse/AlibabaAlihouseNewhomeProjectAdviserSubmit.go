package alihouse

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alihouse"
)

/* 
提交楼盘顾问 
alibaba.alihouse.newhome.project.adviser.submit

提交楼盘顾问
*/
func AlibabaAlihouseNewhomeProjectAdviserSubmit(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeProjectAdviserSubmitAPIRequest, session string) (*alihouse.AlibabaAlihouseNewhomeProjectAdviserSubmitAPIResponse, error) {
    var resp alihouse.AlibabaAlihouseNewhomeProjectAdviserSubmitAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
