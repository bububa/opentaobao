package wdk

// BarcodeBo 
type BarcodeBo struct {

    // 条码
    Barcode   string `json:"barcode,omitempty"`

    // 条码商品规格，6：比如一个条码对应6瓶啤酒
    SpuSpec   int64 `json:"spu_spec,omitempty"`

    // 是否主条码:1是主条码，0非主条码
    MainFlag   string `json:"main_flag,omitempty"`

}
