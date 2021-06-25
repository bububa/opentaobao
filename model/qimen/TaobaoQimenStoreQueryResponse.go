package qimen

import (
    "github.com/bububa/opentaobao/model"
)

/* 
门店信息查询接口 APIResponse
taobao.qimen.store.query

商家在ERP等系统中调用该接口，查询门店相关信息
*/
type TaobaoQimenStoreQueryAPIResponse struct {
    model.CommonResponse
    Response *TaobaoQimenStoreQueryResponse `json:"taobao_qimen_store_query_response,omitempty"`
}

type TaobaoQimenStoreQueryResponse struct {

    // 门店名称
    StoreName   string `json:"store_name,omitempty"`

    // 备注
    Remark   string `json:"remark,omitempty"`

    // 门店主营类目
    MainCategory   int64 `json:"main_category,omitempty"`

    // 响应code
    QimenCode   string `json:"qimen_code,omitempty"`

    // 关闭营业时间
    EndTime   string `json:"end_time,omitempty"`

    // 商户名称
    CompanyName   string `json:"company_name,omitempty"`

    // 开始营业时间
    StartTime   string `json:"start_time,omitempty"`

    // 响应消息
    Message   string `json:"message,omitempty"`

    // 门店状态
    StoreStatus   string `json:"store_status,omitempty"`

    // 响应标示
    Flag   string `json:"flag,omitempty"`

    // 商户介绍
    StoreDescription   string `json:"store_description,omitempty"`

    // 地址信息
    Address   *Address `json:"address,omitempty"`

    // 需要关联的线上店铺ID
    ShopId   string `json:"shop_id,omitempty"`

    // 所有者信息
    StoreKeeper   *StoreKeeper `json:"store_keeper,omitempty"`

    // 类型
    StoreType   string `json:"store_type,omitempty"`

    // ERP系统中 门店编码
    StoreCode   string `json:"store_code,omitempty"`

}
