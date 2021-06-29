package logistic

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
AG退货入仓状态写接口 API返回值 
taobao.nextone.logistics.warehouse.update

商家上传退货入仓状态给ag
*/
type TaobaoNextoneLogisticsWarehouseUpdateAPIResponse struct {
    model.CommonResponse
    TaobaoNextoneLogisticsWarehouseUpdateResponse
}

// AG退货入仓状态写接口 成功返回结果
type TaobaoNextoneLogisticsWarehouseUpdateResponse struct {
    XMLName xml.Name `xml:"nextone_logistics_warehouse_update_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // resultData
    ResultData   string `json:"result_data,omitempty" xml:"result_data,omitempty"`
    // errorInfo
    ErrInfo   string `json:"err_info,omitempty" xml:"err_info,omitempty"`
    // errorCode
    ErrCode   string `json:"err_code,omitempty" xml:"err_code,omitempty"`
    // success
    Succeed   bool `json:"succeed,omitempty" xml:"succeed,omitempty"`
}
