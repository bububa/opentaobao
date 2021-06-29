package simba

// TaobaoSimbaSalestarKeywordsQscoreSplitGetResultDTO 
type TaobaoSimbaSalestarKeywordsQscoreSplitGetResultDTO struct {
    // 返回新质量分实体信息
    Result   *QScoreSplitDTO `json:"result,omitempty" xml:"result,omitempty"`
    // 返回成功/错误码
    Key   string `json:"key,omitempty" xml:"key,omitempty"`
    // 返回信息
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
}
