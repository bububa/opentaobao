package dt

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoadsdataimportAPIResponse 数据导入 API返回值
// taobao.ads.data.import
//
// 数据导入
type TaobaoadsdataimportAPIResponse struct {
	model.CommonResponse
	TaobaoadsdataimportAPIResponseModel
}

// TaobaoadsdataimportAPIResponseModel is 数据导入 成功返回结果
type TaobaoadsdataimportAPIResponseModel struct {
	XMLName xml.Name `xml:"ads_data_import_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回信息
	RetMsg string `json:"ret_msg,omitempty" xml:"ret_msg,omitempty"`
	// 0:成功/-1:失败
	RetCode int64 `json:"ret_code,omitempty" xml:"ret_code,omitempty"`
}
