package logistic

import (
    "github.com/bububa/opentaobao/model"
)

/* 
商家按照仓的类型获取仓库接口 APIResponse
taobao.wlb.stores.baseinfo.get

通过USERID和仓库类型，获取商家自有仓库或菜鸟仓库或全部仓库
*/
type TaobaoWlbStoresBaseinfoGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoWlbStoresBaseinfoGetResponse `json:"taobao_wlb_stores_baseinfo_get_response,omitempty"`
}

type TaobaoWlbStoresBaseinfoGetResponse struct {

    // 仓库列表
    StoreInfoList   []StoreInfo `json:"store_info_list,omitempty"`

    // 返回的总数
    TotalCount   int64 `json:"total_count,omitempty"`

}
