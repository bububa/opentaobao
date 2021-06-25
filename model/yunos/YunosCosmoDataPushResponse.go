package yunos

import (
    "github.com/bububa/opentaobao/model"
)

/* 
COSMO-PUSH模式数据接入 APIResponse
yunos.cosmo.data.push

YunOS提供外部数据源接入，并输出到多端设备上，该接口提供了PUSH模式的数据接入
*/
type YunosCosmoDataPushAPIResponse struct {
    model.CommonResponse
    Response *YunosCosmoDataPushResponse `json:"yunos_cosmo_data_push_response,omitempty"`
}

type YunosCosmoDataPushResponse struct {

    // result
    Result   *DpResult `json:"result,omitempty"`

}
