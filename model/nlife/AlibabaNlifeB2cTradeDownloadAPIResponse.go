package nlife

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaNlifeB2cTradeDownloadAPIResponse
b2c下载订单 API返回值
alibaba.nlife.b2c.trade.download

下载零售商在零售+平台创建的订单 */
type AlibabaNlifeB2cTradeDownloadAPIResponse struct {
	model.CommonResponse
	AlibabaNlifeB2cTradeDownloadAPIResponseModel
}

// AlibabaNlifeB2cTradeDownloadAPIResponseModel is b2c下载订单 成功返回结果
type AlibabaNlifeB2cTradeDownloadAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_nlife_b2c_trade_download_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询命中数量
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
	// 订单列表
	OrderList []Order `json:"order_list,omitempty" xml:"order_list>order,omitempty"`
}
