package fenxiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据商品ID或商家编码修改后端商品 APIResponse
taobao.scitem.update

根据商品ID或商家编码修改后端商品
*/
type TaobaoScitemUpdateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"scitem_update_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 更新商品数量,1表示成功更新了一条数据，0：表示未找到匹配的数据
    
    UpdateRows   int64 `json:"update_rows,omitempty" xml:"