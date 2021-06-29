package alihouse

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alihouse"
)

/* 
删除楼盘预售证 
alibaba.alihouse.newhome.project.presalepermit.delete

删除楼盘预售证信息
*/
func AlibabaAlihouseNewhomeProjectPresalepermitDelete(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeProjectPresalepermitDeleteRequest, session string) (*alihouse.AlibabaAlihouseNewhomeProjectPresalepermitDeleteAPIResponse, error) {
    var resp alihouse.AlibabaAlihouseNewhomeProjectPresalepermitDeleteAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
