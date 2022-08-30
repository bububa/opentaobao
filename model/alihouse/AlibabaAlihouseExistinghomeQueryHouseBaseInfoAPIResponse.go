package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeQueryHouseBaseInfoAPIResponse 查询房源基本信息 API返回值
// alibaba.alihouse.existinghome.query.house.base.info
//
// 查询房源基本信息
type AlibabaAlihouseExistinghomeQueryHouseBaseInfoAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseExistinghomeQueryHouseBaseInfoAPIResponseModel
}

// AlibabaAlihouseExistinghomeQueryHouseBaseInfoAPIResponseModel is 查询房源基本信息 成功返回结果
type AlibabaAlihouseExistinghomeQueryHouseBaseInfoAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_existinghome_query_house_base_info_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回实体类
	Result *AlibabaAlihouseExistinghomeQueryHouseBaseInfoResult `json:"result,omitempty" xml:"result,omitempty"`
}
