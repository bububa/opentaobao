package simba

// TaobaoSimbaSalestarKeywordsQscoreSplitGetResultDto 
type TaobaoSimbaSalestarKeywordsQscoreSplitGetResultDto struct {

    // 返回新质量分实体信息
    Result   *QScoreSplitDto `json:"result,omitempty"`

    // 返回成功/错误码
    Key   string `json:"key,omitempty"`

    // 返回信息
    Message   string `json:"message,omitempty"`

}
