package btrip

// SegmentCabinPriceRs 
type SegmentCabinPriceRs struct {
    // 仓位信息
    Cabin   *CabinRs `json:"cabin,omitempty" xml:"cabin,omitempty"`
    // 搜索提供的价格信息
    SearchPrice   *SearchPriceRs `json:"search_price,omitempty" xml:"search_price,omitempty"`
    // 航段
    Segment   *SegmentPositionRs `json:"segment,omitempty" xml:"segment,omitempty"`
}
