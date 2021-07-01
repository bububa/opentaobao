package btrip

// SegmentCabinPriceRS 
type SegmentCabinPriceRS struct {
    // 仓位信息
    Cabin   *CabinRS `json:"cabin,omitempty" xml:"cabin,omitempty"`
    // 搜索提供的价格信息
    SearchPrice   *SearchPriceRS `json:"search_price,omitempty" xml:"search_price,omitempty"`
    // 航段
    Segment   *SegmentPositionRS `json:"segment,omitempty" xml:"segment,omitempty"`
}
