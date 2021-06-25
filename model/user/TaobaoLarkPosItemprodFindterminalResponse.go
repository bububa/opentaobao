package user

import (
    "github.com/bububa/opentaobao/model"
)

/* 
终端配置支持 APIResponse
taobao.lark.pos.itemprod.findterminal

终端配置支持,读取如果不存在则创建和远程的连接配置并返回
*/
type TaobaoLarkPosItemprodFindterminalAPIResponse struct {
    model.CommonResponse
    Response *TaobaoLarkPosItemprodFindterminalResponse `json:"taobao_lark_pos_itemprod_findterminal_response,omitempty"`
}

type TaobaoLarkPosItemprodFindterminalResponse struct {

    // 终端配置信息响应
    Data   string `json:"data,omitempty"`

}
