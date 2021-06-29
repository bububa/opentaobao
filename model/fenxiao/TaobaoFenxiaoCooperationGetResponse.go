package fenxiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商或分销商获取合作关系信息 APIResponse
taobao.fenxiao.cooperation.get

获取供应商的合作关系信息
*/
type TaobaoFenxiaoCooperationGetAPIResponse struct {
    model.CommonResponse
    TaobaoFenxiaoCooperationGetResponse
}

type TaobaoFenxiaoCooperationGetResponse struct {
    XMLName xml.Name `xml:"fenxiao_cooperation_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 结果记录数
    
    TotalResults   int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`

    
    // 合作分销关系
    
    Cooperations   []Cooperation `json:"cooperations,omitempty" xml:"cooperations>cooperation,omitempty"`
    
    
}
