package category

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取标准类目属性值 APIResponse
taobao.itempropvalues.get

获取标准类目属性值
*/
type TaobaoItempropvaluesGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"itempropvalues_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 最近修改时间。格式:yyyy-MM-dd HH:mm:ss
    
    LastModified   string `json:"last_modified,omitempty" xml:"