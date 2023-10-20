package dt

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAdsDataQueryAPIResponse 导入数据查询 API返回值
// taobao.ads.data.query
//
// 导入数据查询
type TaobaoAdsDataQueryAPIResponse struct {
	model.CommonResponse
	TaobaoAdsDataQueryAPIResponseModel
}

// TaobaoAdsDataQueryAPIResponseModel is 导入数据查询 成功返回结果
type TaobaoAdsDataQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"ads_data_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 导入数据信息
	RetValue []ExternalTaskDataImportDto `json:"ret_value,omitempty" xml:"ret_value>external_task_data_import_dto,omitempty"`
	// 返回信息
	RetMsg string `json:"ret_msg,omitempty" xml:"ret_msg,omitempty"`
	// 0:成功/-1:失败
	RetCode int64 `json:"ret_code,omitempty" xml:"ret_code,omitempty"`
}
