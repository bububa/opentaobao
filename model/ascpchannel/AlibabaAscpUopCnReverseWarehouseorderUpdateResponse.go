package ascpchannel

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
供应链中台逆向入库单修改服务 API返回值 
alibaba.ascp.uop.cn.reverse.warehouseorder.update

供应链中台逆向入库单修改服务
*/
type AlibabaAscpUopCnReverseWarehouseorderUpdateAPIResponse struct {
    model.CommonResponse
    AlibabaAscpUopCnReverseWarehouseorderUpdateResponse
}

// 供应链中台逆向入库单修改服务 成功返回结果
type AlibabaAscpUopCnReverseWarehouseorderUpdateResponse struct {
    XMLName xml.Name `xml:"alibaba_ascp_uop_cn_reverse_warehouseorder_update_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回值包装,result为返回具体消息内容
    UpdateReverseWarehouseOrderResponse   *ResultWrapper `json:"update_reverse_warehouse_order_response,omitempty" xml:"update_reverse_warehouse_order_response,omitempty"`
}
