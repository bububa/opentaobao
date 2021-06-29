package yunos

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
COSMO-PUSH模式数据接入 APIResponse
yunos.cosmo.data.push

YunOS提供外部数据源接入，并输出到多端设备上，该接口提供了PUSH模式的数据接入
*/
type YunosCosmoDataPushAPIResponse struct {
    model.CommonResponse
    YunosCosmoDataPushResponse
}

type YunosCosmoDataPushResponse struct {
    XMLName xml.Name `xml:"yunos_cosmo_data_push_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *DpResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
