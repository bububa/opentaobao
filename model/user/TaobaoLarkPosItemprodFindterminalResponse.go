package user

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
终端配置支持 APIResponse
taobao.lark.pos.itemprod.findterminal

终端配置支持,读取如果不存在则创建和远程的连接配置并返回
*/
type TaobaoLarkPosItemprodFindterminalAPIResponse struct {
    model.CommonResponse
    TaobaoLarkPosItemprodFindterminalResponse
}

type TaobaoLarkPosItemprodFindterminalResponse struct {
    XMLName xml.Name `xml:"lark_pos_itemprod_findterminal_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 终端配置信息响应
    
    Data   string `json:"data,omitempty" xml:"data,omitempty"`

    
}
