package qimen

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimenexpressinfoqueryAPIResponse 配送公司信息查询接口 API返回值
// taobao.qimen.expressinfo.query
//
// 配送公司信息查询
type TaobaoqimenexpressinfoqueryAPIResponse struct {
	model.CommonResponse
	TaobaoqimenexpressinfoqueryAPIResponseModel
}

// TaobaoqimenexpressinfoqueryAPIResponseModel is 配送公司信息查询接口 成功返回结果
type TaobaoqimenexpressinfoqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_expressinfo_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//
	Response *TaobaoqimenexpressinfoqueryResponse `json:"response,omitempty" xml:"response,omitempty"`
}
