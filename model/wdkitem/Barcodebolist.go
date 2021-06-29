package wdkitem

// Barcodebolist 
type Barcodebolist struct {
    // barcode
    Barcode   string `json:"barcode,omitempty" xml:"barcode,omitempty"`
    // 规格
    SpuSpec   string `json:"spu_spec,omitempty" xml:"spu_spec,omitempty"`
}
