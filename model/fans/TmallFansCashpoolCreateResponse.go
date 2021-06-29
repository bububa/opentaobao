package fans

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
创建资金池 API返回值 
tmall.fans.cashpool.create

商家创建资金池接口
*/
type TmallFansCashpoolCreateAPIResponse struct {
    model.CommonResponse
    TmallFansCashpoolCreateResponse
}

// 创建资金池 成功返回结果
type TmallFansCashpoolCreateResponse struct {
    XMLName xml.Name `xml:"tmall_fans_cashpool_create_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回对象
    FansResult   *FansResult `json:"fans_result,omitempty" xml:"fans_result,omitempty"`
}
