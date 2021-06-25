package logistic

import (
    "github.com/bububa/opentaobao/model"
)

/* 
预采购服务包查询接口 APIResponse
alibaba.ele.fengniao.service.package.query

查询门店所在经纬度可用服务包的接口
*/
type AlibabaEleFengniaoServicePackageQueryAPIResponse struct {
    model.CommonResponse
    Response *AlibabaEleFengniaoServicePackageQueryResponse `json:"alibaba_ele_fengniao_service_package_query_response,omitempty"`
}

type AlibabaEleFengniaoServicePackageQueryResponse struct {

    // servicePackages
    ServicePackages   []AlibabaEleFengniaoServicePackageQueryResult `json:"service_packages,omitempty"`

}
