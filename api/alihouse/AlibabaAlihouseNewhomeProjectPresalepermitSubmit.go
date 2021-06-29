package alihouse

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alihouse"
)

/* 
提交预售证 
alibaba.alihouse.newhome.project.presalepermit.submit

提交楼盘预售证信息
*/
func AlibabaAlihouseNewhomeProjectPresalepermitSubmit(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeProjectPresalepermitSubmitRequest, session string) (*alihouse.AlibabaAlihouseNewhomeProjectPresalepermitSubmitAPIResponse, error) {
    var resp alihouse.AlibabaAlihouseNewhomeProjectPresalepermitSubmitAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
