package wms

import (
    "github.com/bububa/opentaobao/model"
)

/* 
商品信息查询 APIResponse
taobao.wlb.wms.sku.get

商品信息查询
*/
type TaobaoWlbWmsSkuGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoWlbWmsSkuGetResponse `json:"wlb_wms_sku_get_response,omitempty"` 
    TaobaoWlbWmsSkuGetResponse
}

/* model for simplify = false
type TaobaoWlbWmsSkuGetResponse struct {

    // 拓展属性, key-value结构，格式要求： 以英文分号“;”分隔每组key-value，以英文冒号“:”分隔key与value。如果value中带有分号，需要转成下划线“_”，如果带有冒号，需要转成中划线“-”
    
    ExtendFields   string `json:"extend_fields,omitempty"`
    

    // 是否启用批次管理
    
    IsBatchMgt   bool `json:"is_batch_mgt,omitempty"`
    

    // 启用标识
    
    UseYn   bool `json:"use_yn,omitempty"`
    

    // 成本价，单位分
    
    CostPrice   int64 `json:"cost_price,omitempty"`
    

    // 零售价，单位分
    
    ItemPrice   int64 `json:"item_price,omitempty"`
    

    // 吊牌价，单位分
    
    TagPrice   int64 `json:"tag_price,omitempty"`
    

    // 是否危险品
    
    IsDanger   bool `json:"is_danger,omitempty"`
    

    // 是否易碎品
    
    IsHygroscopic   bool `json:"is_hygroscopic,omitempty"`
    

    // 是否启用序列号管理
    
    IsSnMgt   bool `json:"is_sn_mgt,omitempty"`
    

    // 保质期预警天数
    
    AdventLifecycle   int64 `json:"advent_lifecycle,omitempty"`
    

    // 保质期禁售天数
    
    LockupLifecycle   int64 `json:"lockup_lifecycle,omitempty"`
    

    // 保质期禁收天数
    
    RejectLifecycle   int64 `json:"reject_lifecycle,omitempty"`
    

    // 保质期天数
    
    Lifecycle   int64 `json:"lifecycle,omitempty"`
    

    // 是否启用保质期管理
    
    IsShelflife   bool `json:"is_shelflife,omitempty"`
    

    // 批准文号
    
    ApprovalNumber   string `json:"approval_number,omitempty"`
    

    // 场地
    
    OriginAddress   int64 `json:"origin_address,omitempty"`
    

    // 箱规
    
    Pcs   int64 `json:"pcs,omitempty"`
    

    // 体积，单位立方厘米
    
    Volume   int64 `json:"volume,omitempty"`
    

    // 高度，单位毫米
    
    Height   int64 `json:"height,omitempty"`
    

    // 宽度，单位毫米
    
    Width   int64 `json:"width,omitempty"`
    

    // 长度，单位毫米
    
    Length   int64 `json:"length,omitempty"`
    

    // 净重，单位克
    
    NetWeight   int64 `json:"net_weight,omitempty"`
    

    // 毛重，单位克
    
    GrossWeight   int64 `json:"gross_weight,omitempty"`
    

    // 尺寸
    
    Size   string `json:"size,omitempty"`
    

    // 颜色
    
    Color   string `json:"color,omitempty"`
    

    // 规格
    
    Specification   string `json:"specification,omitempty"`
    

    // 品牌名称
    
    BrandName   string `json:"brand_name,omitempty"`
    

    // 品牌编码
    
    Brand   string `json:"brand,omitempty"`
    

    // 商品类别名称
    
    CategoryName   string `json:"category_name,omitempty"`
    

    // 商品类别编码（外部系统类别）
    
    Category   string `json:"category,omitempty"`
    

    // 商品类别 NORMAL：普通商品、 COMBINE：组合商品、 DISTRIBUTION：分销商品
    
    Type   string `json:"type,omitempty"`
    

    // 商品标题
    
    Title   string `json:"title,omitempty"`
    

    // 商品名称
    
    Name   string `json:"name,omitempty"`
    

    // 条形码，多条码请用”;”分隔；
    
    BarCode   string `json:"bar_code,omitempty"`
    

    // 商家商品编码,与itemid必须有一个值不为空
    
    IitemCode   string `json:"iitem_code,omitempty"`
    

    // 菜鸟商品ID,与itemcode必须有一个值不为空
    
    ItemId   string `json:"item_id,omitempty"`
    

    // 错误信息
    
    WlErrorMsg   string `json:"wl_error_msg,omitempty"`
    

    // 错误编码
    
    WlErrorCode   string `json:"wl_error_code,omitempty"`
    

    // 是否成功
    
    WlSuccess   bool `json:"wl_success,omitempty"`
    

    // 是否区域销售
    
    IsAreaSale   bool `json:"is_area_sale,omitempty"`
    

}
*/

type TaobaoWlbWmsSkuGetResponse struct {

    // 拓展属性, key-value结构，格式要求： 以英文分号“;”分隔每组key-value，以英文冒号“:”分隔key与value。如果value中带有分号，需要转成下划线“_”，如果带有冒号，需要转成中划线“-”
    ExtendFields   string `json:"extend_fields,omitempty"`

    // 是否启用批次管理
    IsBatchMgt   bool `json:"is_batch_mgt,omitempty"`

    // 启用标识
    UseYn   bool `json:"use_yn,omitempty"`

    // 成本价，单位分
    CostPrice   int64 `json:"cost_price,omitempty"`

    // 零售价，单位分
    ItemPrice   int64 `json:"item_price,omitempty"`

    // 吊牌价，单位分
    TagPrice   int64 `json:"tag_price,omitempty"`

    // 是否危险品
    IsDanger   bool `json:"is_danger,omitempty"`

    // 是否易碎品
    IsHygroscopic   bool `json:"is_hygroscopic,omitempty"`

    // 是否启用序列号管理
    IsSnMgt   bool `json:"is_sn_mgt,omitempty"`

    // 保质期预警天数
    AdventLifecycle   int64 `json:"advent_lifecycle,omitempty"`

    // 保质期禁售天数
    LockupLifecycle   int64 `json:"lockup_lifecycle,omitempty"`

    // 保质期禁收天数
    RejectLifecycle   int64 `json:"reject_lifecycle,omitempty"`

    // 保质期天数
    Lifecycle   int64 `json:"lifecycle,omitempty"`

    // 是否启用保质期管理
    IsShelflife   bool `json:"is_shelflife,omitempty"`

    // 批准文号
    ApprovalNumber   string `json:"approval_number,omitempty"`

    // 场地
    OriginAddress   int64 `json:"origin_address,omitempty"`

    // 箱规
    Pcs   int64 `json:"pcs,omitempty"`

    // 体积，单位立方厘米
    Volume   int64 `json:"volume,omitempty"`

    // 高度，单位毫米
    Height   int64 `json:"height,omitempty"`

    // 宽度，单位毫米
    Width   int64 `json:"width,omitempty"`

    // 长度，单位毫米
    Length   int64 `json:"length,omitempty"`

    // 净重，单位克
    NetWeight   int64 `json:"net_weight,omitempty"`

    // 毛重，单位克
    GrossWeight   int64 `json:"gross_weight,omitempty"`

    // 尺寸
    Size   string `json:"size,omitempty"`

    // 颜色
    Color   string `json:"color,omitempty"`

    // 规格
    Specification   string `json:"specification,omitempty"`

    // 品牌名称
    BrandName   string `json:"brand_name,omitempty"`

    // 品牌编码
    Brand   string `json:"brand,omitempty"`

    // 商品类别名称
    CategoryName   string `json:"category_name,omitempty"`

    // 商品类别编码（外部系统类别）
    Category   string `json:"category,omitempty"`

    // 商品类别 NORMAL：普通商品、 COMBINE：组合商品、 DISTRIBUTION：分销商品
    Type   string `json:"type,omitempty"`

    // 商品标题
    Title   string `json:"title,omitempty"`

    // 商品名称
    Name   string `json:"name,omitempty"`

    // 条形码，多条码请用”;”分隔；
    BarCode   string `json:"bar_code,omitempty"`

    // 商家商品编码,与itemid必须有一个值不为空
    IitemCode   string `json:"iitem_code,omitempty"`

    // 菜鸟商品ID,与itemcode必须有一个值不为空
    ItemId   string `json:"item_id,omitempty"`

    // 错误信息
    WlErrorMsg   string `json:"wl_error_msg,omitempty"`

    // 错误编码
    WlErrorCode   string `json:"wl_error_code,omitempty"`

    // 是否成功
    WlSuccess   bool `json:"wl_success,omitempty"`

    // 是否区域销售
    IsAreaSale   bool `json:"is_area_sale,omitempty"`

}
