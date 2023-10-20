package qimen

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimenwarehouseinfosynchronizeAPIResponse 仓库同步接口 API返回值
// taobao.qimen.warehouseinfo.synchronize
//
// 仓库同步接口
type TaobaoqimenwarehouseinfosynchronizeAPIResponse struct {
	model.CommonResponse
	TaobaoqimenwarehouseinfosynchronizeAPIResponseModel
}

// TaobaoqimenwarehouseinfosynchronizeAPIResponseModel is 仓库同步接口 成功返回结果
type TaobaoqimenwarehouseinfosynchronizeAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_warehouseinfo_synchronize_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应报文
	Response *TaobaoqimenwarehouseinfosynchronizeResponse `json:"response,omitempty" xml:"response,omitempty"`
}
