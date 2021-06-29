package nlife

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
b2c下载订单 APIResponse
alibaba.nlife.b2c.trade.download

下载零售商在零售+平台创建的订单
*/
type AlibabaNlifeB2cTradeDownloadAPIResponse struct {
    model.CommonResponse
    AlibabaNlifeB2cTradeDownloadResponse
}

type AlibabaNlifeB2cTradeDownloadResponse struct {
    XMLName xml.Name `xml:"alibaba_nlife_b2c_trade_download_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 查询命中数量
    
    Total   int64 `json:"total,omitempty" xml:"total,omitempty"`

    
    // 订单列表
    
    OrderList   []Order `json:"order_list,omitempty" xml:"order_list>order,omitempty"`
    
    
}
