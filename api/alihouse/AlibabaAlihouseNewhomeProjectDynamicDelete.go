package alihouse

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alihouse"
)

/* 
楼盘快讯删除 
alibaba.alihouse.newhome.project.dynamic.delete

楼盘快讯删除
*/
func AlibabaAlihouseNewhomeProjectDynamicDelete(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeProjectDynamicDeleteAPIRequest, session string) (*alihouse.AlibabaAlihouseNewhomeProjectDynamicDeleteAPIResponse, error) {
    var resp alihouse.AlibabaAlihouseNewhomeProjectDynamicDeleteAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
