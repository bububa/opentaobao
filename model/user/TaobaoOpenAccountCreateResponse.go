package user

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
Open Account导入数据 APIResponse
taobao.open.account.create

Open Account导入数据
*/
type TaobaoOpenAccountCreateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"open_account_create_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 插入数据的Open Account Id的列表
    
    Datas   []OpenaccountLong `json:"datas,omitempty" xml:"