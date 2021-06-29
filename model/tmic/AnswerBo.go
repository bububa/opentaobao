package tmic

// AnswerBo 
type AnswerBo struct {

    // 问题编码，问卷中的问题的唯一编码，从问卷信息接口的应答中获取
    
    QuestionCode   string `json:"question_code,omitempty" xml:"question_code,omitempty"`
    

    // 子问卷编码，问卷中某些问卷包含子问题，唯一区分子问题,从问卷信息接口的应答中获取
    
    SubQuestionCode   string `json:"sub_question_code,omitempty" xml:"sub_question_code,omitempty"`
    

    // 选项唯一编码，每道问题的每个选项的唯一区别标识，从问卷信息接口的应答中获取
    
    OptionCode   string `json:"option_code,omitempty" xml:"option_code,omitempty"`
    

    // 该选项是否被选中，1选中，2未选中
    
    OptionChecked   string `json:"option_checked,omitempty" xml:"option_checked,omitempty"`
    

    // 文本题的答案，如果是选择题则不填
    
    AnswerValue   string `json:"answer_value,omitempty" xml:"answer_value,omitempty"`
    

}
