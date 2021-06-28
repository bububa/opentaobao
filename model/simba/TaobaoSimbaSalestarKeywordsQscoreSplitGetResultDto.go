package simba

// TaobaoSimbaSalestarKeywordsQscoreSplitGetResultDto 
/* model for simplify = false
type TaobaoSimbaSalestarKeywordsQscoreSplitGetResultDto struct {

    // 返回新质量分实体信息
    
    Result  *struct {
        QScoreSplitDto  *QScoreSplitDto `json:"q_score_split_dto,omitempty"`
    } `json:"result,omitempty"`
    

    // 返回成功/错误码
    
    Key   string `json:"key,omitempty"`
    

    // 返回信息
    
    Message   string `json:"message,omitempty"`
    

}
*/

// TaobaoSimbaSalestarKeywordsQscoreSplitGetResultDto 
type TaobaoSimbaSalestarKeywordsQscoreSplitGetResultDto struct {

    // 返回新质量分实体信息
    Result   *QScoreSplitDto `json:"result,omitempty"`

    // 返回成功/错误码
    Key   string `json:"key,omitempty"`

    // 返回信息
    Message   string `json:"message,omitempty"`

}
