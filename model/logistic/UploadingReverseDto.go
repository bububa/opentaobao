package logistic

// UploadingReverseDTO 
type UploadingReverseDTO struct {
    // 扩展字段，JSONObject格式
    Extra   string `json:"extra,omitempty" xml:"extra,omitempty"`
    // 商品行列表
    GoodsItemDTOList   []WarehouseReverseGoodsItemDTO `json:"goods_item_d_t_o_list,omitempty" xml:"goods_item_d_t_o_list>warehouse_reverse_goods_item_dto,omitempty"`
    // 仓库名称
    WarehouseName   string `json:"warehouse_name,omitempty" xml:"warehouse_name,omitempty"`
    // 详细地址
    Adr   string `json:"adr,omitempty" xml:"adr,omitempty"`
    // 县区
    DistrictName   string `json:"district_name,omitempty" xml:"district_name,omitempty"`
    // 城市
    CityName   string `json:"city_name,omitempty" xml:"city_name,omitempty"`
    // 省份
    ProvinceName   string `json:"province_name,omitempty" xml:"province_name,omitempty"`
    // 国家
    CountryName   string `json:"country_name,omitempty" xml:"country_name,omitempty"`
    // 创建时间
    CreateTime   string `json:"create_time,omitempty" xml:"create_time,omitempty"`
    // 销退单状态(1=已创建；2=待入库；3=已入库，5=已取消)
    Status   int64 `json:"status,omitempty" xml:"status,omitempty"`
    // 物流公司code
    CpCode   string `json:"cp_code,omitempty" xml:"cp_code,omitempty"`
    // 物流公司名称
    CpName   string `json:"cp_name,omitempty" xml:"cp_name,omitempty"`
    // 运单编号
    MailNo   string `json:"mail_no,omitempty" xml:"mail_no,omitempty"`
    // 主订单编号
    Tid   int64 `json:"tid,omitempty" xml:"tid,omitempty"`
    // 销退单ID
    Id   string `json:"id,omitempty" xml:"id,omitempty"`
}
