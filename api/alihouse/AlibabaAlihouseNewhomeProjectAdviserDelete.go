package alihouse

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alihouse"
)

/* 
删除楼盘顾问 
alibaba.alihouse.newhome.project.adviser.delete

删除楼盘顾问
*/
func AlibabaAlihouseNewhomeProjectAdviserDelete(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeProjectAdviserDeleteRequest, session string) (*alihouse.AlibabaAlihouseNewhomeProjectAdviserDeleteAPIResponse, error) {
    var resp alihouse.AlibabaAlihouseNewhomeProjectAdviserDeleteAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
