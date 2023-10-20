package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkbmpaiyangstatdataqueryAPIResponse 派样统计数据查询 API返回值
// alibaba.wdk.bm.paiyang.stat.data.query
//
// 派样统计数据查询
type AlibabawdkbmpaiyangstatdataqueryAPIResponse struct {
	model.CommonResponse
	AlibabawdkbmpaiyangstatdataqueryAPIResponseModel
}

// AlibabawdkbmpaiyangstatdataqueryAPIResponseModel is 派样统计数据查询 成功返回结果
type AlibabawdkbmpaiyangstatdataqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_bm_paiyang_stat_data_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参对象
	Result *BmPageResult `json:"result,omitempty" xml:"result,omitempty"`
}
