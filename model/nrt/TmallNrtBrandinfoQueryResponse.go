package nrt

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
品牌数据查询 APIResponse
tmall.nrt.brandinfo.query

商家获取自己旗舰店授权的品牌id列表
*/
type TmallNrtBrandinfoQueryAPIResponse struct {
    model.CommonResponse
    TmallNrtBrandinfoQueryResponse
}

type TmallNrtBrandinfoQueryResponse struct {
    XMLName xml.Name `xml:"tmall_nrt_brandinfo_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 品牌id
    
    Datas   []string `json:"datas,omitempty" xml:"datas>string,omitempty"`
    
    
}
