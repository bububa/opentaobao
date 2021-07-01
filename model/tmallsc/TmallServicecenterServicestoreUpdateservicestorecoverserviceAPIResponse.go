package tmallsc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
更新网点覆盖的服务 API返回值 
tmall.servicecenter.servicestore.updateservicestorecoverservice

更新网点覆盖的服务，唯一性校验：服务商淘宝账号+网点编码+biz_type
前提是网点要存在，
如果需要新增的网点覆盖的服务不存在，会更新失败。
网点覆盖的服务包含了业务类型(比如电器预约安装)、天猫服务的servicecode列表、授权的类目和品牌
*/
type TmallServicecenterServicestoreUpdateservicestorecoverserviceAPIResponse struct {
    model.CommonResponse
    TmallServicecenterServicestoreUpdateservicestorecoverserviceAPIResponseModel
}

// 更新网点覆盖的服务 成功返回结果
type TmallServicecenterServicestoreUpdateservicestorecoverserviceAPIResponseModel struct {
    XMLName xml.Name `xml:"tmall_servicecenter_servicestore_updateservicestorecoverservice_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *ResultBase `json:"result,omitempty" xml:"result,omitempty"`
}
