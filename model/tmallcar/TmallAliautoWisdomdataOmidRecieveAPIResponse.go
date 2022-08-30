package tmallcar

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallAliautoWisdomdataOmidRecieveAPIResponse 大搜车车型参配数据接入 API返回值
// tmall.aliauto.wisdomdata.omid.recieve
//
// 大搜车车型参配数据接入
type TmallAliautoWisdomdataOmidRecieveAPIResponse struct {
	model.CommonResponse
	TmallAliautoWisdomdataOmidRecieveAPIResponseModel
}

// TmallAliautoWisdomdataOmidRecieveAPIResponseModel is 大搜车车型参配数据接入 成功返回结果
type TmallAliautoWisdomdataOmidRecieveAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_aliauto_wisdomdata_omid_recieve_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口调用状态信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 记录总数
	TotalNum int64 `json:"total_num,omitempty" xml:"total_num,omitempty"`
}
