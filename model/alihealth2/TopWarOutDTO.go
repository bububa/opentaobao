package alihealth2

// TopWarOutDto 
type TopWarOutDto struct {
    // 出库单号
    BillNo   string `json:"bill_no,omitempty" xml:"bill_no,omitempty"`
    // 单据日期
    BizDate   string `json:"biz_date,omitempty" xml:"biz_date,omitempty"`
    // 孔雀翎O2O订单ID
    CepOrderId   int64 `json:"cep_order_id,omitempty" xml:"cep_order_id,omitempty"`
    // 商品
    GoodsList   []TopGoodsDto `json:"goods_list,omitempty" xml:"goods_list>top_goods_dto,omitempty"`
    // 淘宝sellerId
    TbSellerId   int64 `json:"tb_seller_id,omitempty" xml:"tb_seller_id,omitempty"`
    // 魔方ID
    CubeId   int64 `json:"cube_id,omitempty" xml:"cube_id,omitempty"`
}
