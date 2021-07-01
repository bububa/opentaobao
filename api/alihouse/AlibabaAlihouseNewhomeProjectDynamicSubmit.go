package alihouse

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alihouse"
)

/* 
提交楼盘快讯 
alibaba.alihouse.newhome.project.dynamic.submit

提交楼盘快讯
*/
func AlibabaAlihouseNewhomeProjectDynamicSubmit(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeProjectDynamicSubmitAPIRequest, session string) (*alihouse.AlibabaAlihouseNewhomeProjectDynamicSubmitAPIResponse, error) {
    var resp alihouse.AlibabaAlihouseNewhomeProjectDynamicSubmitAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
