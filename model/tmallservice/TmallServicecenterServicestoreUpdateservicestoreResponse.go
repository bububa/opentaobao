package tmallservice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
修改网点信息 APIResponse
tmall.servicecenter.servicestore.updateservicestore

修改网点信息。对于同一个服务商，通过 service_store_code 保证网点唯一性。需要保证网点存在才能修改。
错误码
1, 服务商昵称无效
2, 缺少省份
3, 缺少城市
4, 缺少区域
5, 缺少详细地址
6, 传入地址不在标准地址库中，无法解析
7, 缺少网点编码
8, 缺少网点名称
9, 缺少网点电话
10, 网点已存在
11, 网点不存在
12, 系统错误
*/
type TmallServicecenterServicestoreUpdateservicestoreAPIResponse struct {
    model.CommonResponse
    TmallServicecenterServicestoreUpdateservicestoreResponse
}

type TmallServicecenterServicestoreUpdateservicestoreResponse struct {
    XMLName xml.Name `xml:"tmall_servicecenter_servicestore_updateservicestore_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *ResultBase `json:"result,omitempty" xml:"result,omitempty"`

    
}
