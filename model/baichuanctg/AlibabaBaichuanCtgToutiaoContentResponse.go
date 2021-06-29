package baichuanctg

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
微博输出头条数据 APIResponse
alibaba.baichuan.ctg.toutiao.content

百川头条内容获取
*/
type AlibabaBaichuanCtgToutiaoContentAPIResponse struct {
    model.CommonResponse
    AlibabaBaichuanCtgToutiaoContentResponse
}

type AlibabaBaichuanCtgToutiaoContentResponse struct {
    XMLName xml.Name `xml:"alibaba_baichuan_ctg_toutiao_content_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 内容总体结构
    
    Result   *CtgResponse `json:"result,omitempty" xml:"result,omitempty"`

    
}
