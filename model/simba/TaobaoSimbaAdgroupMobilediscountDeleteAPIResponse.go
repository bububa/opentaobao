package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
批量删除adgroup的移动溢价 API返回值 
taobao.simba.adgroup.mobilediscount.delete

批量删除adgroup的移动溢价
*/
type TaobaoSimbaAdgroupMobilediscountDeleteAPIResponse struct {
    model.CommonResponse
    TaobaoSimbaAdgroupMobilediscountDeleteAPIResponseModel
}

// 批量删除adgroup的移动溢价 成功返回结果
type TaobaoSimbaAdgroupMobilediscountDeleteAPIResponseModel struct {
    XMLName xml.Name `xml:"simba_adgroup_mobilediscount_delete_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回成功个数
    Result   int64 `json:"result,omitempty" xml:"result,omitempty"`
    // 返回信息
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    // 返回码
    Key   string `json:"key,omitempty" xml:"key,omitempty"`
}
