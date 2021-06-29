package normalvisa

// TaobaoAlitripTravelNormalvisaGetResultSet 
type TaobaoAlitripTravelNormalvisaGetResultSet struct {
    // 结果
    Results   []NormalVisaInfo `json:"results,omitempty" xml:"results>normal_visa_info,omitempty"`
    // 结果数目
    TotalResults   int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
}
