package fenxiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
修改产品线 APIResponse
taobao.fenxiao.productcat.update

修改产品线
*/
type TaobaoFenxiaoProductcatUpdateAPIResponse struct {
    model.CommonResponse
    TaobaoFenxiaoProductcatUpdateResponse
}

type TaobaoFenxiaoProductcatUpdateResponse struct {
    XMLName xml.Name `xml:"fenxiao_productcat_update_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 操作是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
}
