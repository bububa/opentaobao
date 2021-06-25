package qimen

import (
    "github.com/bububa/opentaobao/model"
)

/* 
商品关联门店查询接口 APIResponse
taobao.qimen.itemstore.query

商家在ERP等系统中调用该接口，查询线上商品所关联的门店列表
*/
type TaobaoQimenItemstoreQueryAPIResponse struct {
    model.CommonResponse
    Response *TaobaoQimenItemstoreQueryResponse `json:"taobao_qimen_itemstore_query_response,omitempty"`
}

type TaobaoQimenItemstoreQueryResponse struct {

    // 响应信息
    Message   string `json:"message,omitempty"`

    // 门店列表
    StoreIds   []Number `json:"store_ids,omitempty"`

    // 响应的标签
    Flag   string `json:"flag,omitempty"`

    // 总的门店数
    TotalLines   int64 `json:"total_lines,omitempty"`

    // 响应的code
    QimenCode   string `json:"qimen_code,omitempty"`

}
