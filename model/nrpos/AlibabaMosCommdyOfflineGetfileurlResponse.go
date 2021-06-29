package nrpos

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
去前置机pos商品离线文件下载地址查询接口 APIResponse
alibaba.mos.commdy.offline.getfileurl

去前置机-pos查询离线文件下载地址接口
*/
type AlibabaMosCommdyOfflineGetfileurlAPIResponse struct {
    model.CommonResponse
    AlibabaMosCommdyOfflineGetfileurlResponse
}

type AlibabaMosCommdyOfflineGetfileurlResponse struct {
    XMLName xml.Name `xml:"alibaba_mos_commdy_offline_getfileurl_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *AlibabaMosCommdyOfflineGetfileurlResultDo `json:"result,omitempty" xml:"result,omitempty"`

    
}
