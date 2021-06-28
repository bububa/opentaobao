package bus

// TopStandardInsRecommendResponse 
/* model for simplify = false
type TopStandardInsRecommendResponse struct {

    // 推荐结果
    
    ResultList  struct {
        TopInsProduct  []TopInsProduct `json:"top_ins_product,omitempty"`
    } `json:"result_list,omitempty"`
    

}
*/

// TopStandardInsRecommendResponse 
type TopStandardInsRecommendResponse struct {

    // 推荐结果
    ResultList   []TopInsProduct `json:"result_list,omitempty"`

}
