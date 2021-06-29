package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
对推广组进行单独移动溢价 APIResponse
taobao.simba.adgroup.mobilediscount.update

对推广组进行单独移动溢价
*/
type TaobaoSimbaAdgroupMobilediscountUpdateAPIResponse struct {
    model.CommonResponse
    TaobaoSimbaAdgroupMobilediscountUpdateResponse
}

type TaobaoSimbaAdgroupMobilediscountUpdateResponse struct {
    XMLName xml.Name `xml:"simba_adgroup_mobilediscount_update_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 更新成功的个数
    
    Result   int64 `json:"result,omitempty" xml:"result,omitempty"`

    
    // 返回信息
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`

    
    // 错误码
    
    Key   string `json:"key,omitempty" xml:"key,omitempty"`

    
}
