package qimen

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimenwarehouseinfoqueryAPIResponse 货主仓库资源查询接口 API返回值
// taobao.qimen.warehouseinfo.query
//
// 货主仓库资源查询
type TaobaoqimenwarehouseinfoqueryAPIResponse struct {
	model.CommonResponse
	TaobaoqimenwarehouseinfoqueryAPIResponseModel
}

// TaobaoqimenwarehouseinfoqueryAPIResponseModel is 货主仓库资源查询接口 成功返回结果
type TaobaoqimenwarehouseinfoqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_warehouseinfo_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//
	Response *TaobaoqimenwarehouseinfoqueryResponse `json:"response,omitempty" xml:"response,omitempty"`
}
