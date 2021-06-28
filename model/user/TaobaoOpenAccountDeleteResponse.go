package user

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
OpenAccount删除数据 APIResponse
taobao.open.account.delete

OpenAccount删除数据
*/
type TaobaoOpenAccountDeleteAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"open_account_delete_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 删除结果
    
    Datas   []OpenaccountVoid `json:"datas,omitempty" xml:"