package wlbimports

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取所有服务列表 APIResponse
taobao.wlb.imports.resource.get

一般进口TOP接口，获取所有服务列表。
*/
type TaobaoWlbImportsResourceGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoWlbImportsResourceGetResponse `json:"taobao_wlb_imports_resource_get_response,omitempty"`
}

type TaobaoWlbImportsResourceGetResponse struct {

    // 一般进口所有服务商列表
    Resources   []ResourceResult `json:"resources,omitempty"`

}
