package user

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
Open Account索引查询 API返回值 
taobao.open.account.index.find

Open Account索引查询
*/
type TaobaoOpenAccountIndexFindAPIResponse struct {
    model.CommonResponse
    TaobaoOpenAccountIndexFindAPIResponseModel
}

// Open Account索引查询 成功返回结果
type TaobaoOpenAccountIndexFindAPIResponseModel struct {
    XMLName xml.Name `xml:"open_account_index_find_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回结果
    Result   *OpenAccountResult `json:"result,omitempty" xml:"result,omitempty"`
}
