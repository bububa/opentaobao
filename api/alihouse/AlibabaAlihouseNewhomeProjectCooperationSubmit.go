package alihouse

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alihouse"
)

/* 
提交KA合作楼盘 
alibaba.alihouse.newhome.project.cooperation.submit

提交KA合作楼盘
*/
func AlibabaAlihouseNewhomeProjectCooperationSubmit(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeProjectCooperationSubmitAPIRequest, session string) (*alihouse.AlibabaAlihouseNewhomeProjectCooperationSubmitAPIResponse, error) {
    var resp alihouse.AlibabaAlihouseNewhomeProjectCooperationSubmitAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
