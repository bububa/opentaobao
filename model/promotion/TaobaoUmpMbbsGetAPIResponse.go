package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取营销积木块列表 API返回值 
taobao.ump.mbbs.get

获取营销积木块列表，可以根据类型获取，也可以将该字段设为空，获取所有的
*/
type TaobaoUmpMbbsGetAPIResponse struct {
    model.CommonResponse
    TaobaoUmpMbbsGetAPIResponseModel
}

// 获取营销积木块列表 成功返回结果
type TaobaoUmpMbbsGetAPIResponseModel struct {
    XMLName xml.Name `xml:"ump_mbbs_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 营销积木块内容列表，内容为json格式的，可以通过ump sdk里面的MBB.fromJson来处理
    Mbbs   []string `json:"mbbs,omitempty" xml:"mbbs>string,omitempty"`
}
