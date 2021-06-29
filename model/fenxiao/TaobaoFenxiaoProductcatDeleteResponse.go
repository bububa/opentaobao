package fenxiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
删除产品线 API返回值 
taobao.fenxiao.productcat.delete

删除产品线
*/
type TaobaoFenxiaoProductcatDeleteAPIResponse struct {
    model.CommonResponse
    TaobaoFenxiaoProductcatDeleteResponse
}

// 删除产品线 成功返回结果
type TaobaoFenxiaoProductcatDeleteResponse struct {
    XMLName xml.Name `xml:"fenxiao_productcat_delete_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 操作是否成功
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
