package trade

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaowtttradeservicegetAPIResponse 获取网厅号卡垂直标信息 API返回值
// taobao.wtt.trade.service.get
//
// 查询网厅订单信息
type TaobaowtttradeservicegetAPIResponse struct {
	model.CommonResponse
	TaobaowtttradeservicegetAPIResponseModel
}

// TaobaowtttradeservicegetAPIResponseModel is 获取网厅号卡垂直标信息 成功返回结果
type TaobaowtttradeservicegetAPIResponseModel struct {
	XMLName xml.Name `xml:"wtt_trade_service_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回实例
	Modules *WtverticalDto `json:"modules,omitempty" xml:"modules,omitempty"`
}
