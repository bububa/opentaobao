package user

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoopenaccountsearchAPIResponse open account数据搜索 API返回值
// taobao.open.account.search
//
// open account数据搜索
type TaobaoopenaccountsearchAPIResponse struct {
	model.CommonResponse
	TaobaoopenaccountsearchAPIResponseModel
}

// TaobaoopenaccountsearchAPIResponseModel is open account数据搜索 成功返回结果
type TaobaoopenaccountsearchAPIResponseModel struct {
	XMLName xml.Name `xml:"open_account_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Data *OpenAccountSearchResult `json:"data,omitempty" xml:"data,omitempty"`
}
